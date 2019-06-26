package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"time"
)

const (
	privKeyPath = "/Users/kjoshi/go/src/github.com/auth/keys/app.rsa"     // `$ openssl genrsa -out app.rsa 2048`
	pubKeyPath  = "/Users/kjoshi/go/src/github.com/auth/keys/app.rsa.pub" // `$ openssl rsa -in app.rsa -pubout > app.rsa.pub`
)

const RefreshTokenValidTime = time.Hour * 72
const AuthTokenValidTime = time.Minute * 15

type AccessType string

const (
	Access  = "access"
	REFRESH = "refresh"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

type User struct {
	Username, PasswordHash, Role string
}

func init() {
	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatal(err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatal(err)
	}

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal(err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatal(err)
	}

}

// https://tools.ietf.org/html/rfc7519
type TokenClaims struct {
	jwt.StandardClaims
	Role string `json:"role"`
	Csrf string `json:"csrf"`
	Type string `json:"type"`
}

func GenerateCSRFSecret() (string, error) {
	return generateRandomString(32)
}

func generateRandomString(s int) (string, error) {
	b := make([]byte, s)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), err
}

func createAuthTokenString(uuid string, role string, csrfString string) (authTokenString string, err error) {

	authTokenExp := time.Now().Add(AuthTokenValidTime).Unix()
	authClaims := TokenClaims{
		jwt.StandardClaims{
			Subject:   uuid,
			ExpiresAt: authTokenExp,
		},
		role,
		csrfString,
		Access,
	}

	// create a signer for rsa 256
	authJwt := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), authClaims)

	// generate the auth token string
	authTokenString, err = authJwt.SignedString(signKey)
	return
}

func createRefreshTokenString(uuid string, role string, csrfString string) (refreshTokenString string, err error) {
	refreshTokenExp := time.Now().Add(RefreshTokenValidTime).Unix()
	refreshClaims := TokenClaims{
		jwt.StandardClaims{
			Subject:   uuid,
			ExpiresAt: refreshTokenExp,
		},
		role,
		csrfString,
		REFRESH,
	}

	// create a signer for rsa 256
	refreshJwt := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), refreshClaims)

	// generate the refresh token string
	refreshTokenString, err = refreshJwt.SignedString(signKey)
	return
}

func CreateNewTokens(uuid string, role string) (authTokenString, refreshTokenString, csrfSecret string, err error) {
	// generate the csrf secret
	csrfSecret, err = GenerateCSRFSecret()
	if err != nil {
		return
	}

	// generate the refresh token
	refreshTokenString, err = createRefreshTokenString(uuid, role, csrfSecret)

	// generate the auth token
	authTokenString, err = createAuthTokenString(uuid, role, csrfSecret)
	if err != nil {
		return
	}
	// don't need to check for err bc we're returning everything anyway
	return
}

func updateRefreshTokenExp(oldRefreshTokenString string) (newRefreshTokenString string, err error) {
	refreshToken, err := jwt.ParseWithClaims(oldRefreshTokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})

	oldRefreshTokenClaims, ok := refreshToken.Claims.(*TokenClaims)
	if !ok {
		return
	}

	refreshTokenExp := time.Now().Add(RefreshTokenValidTime).Unix()

	refreshClaims := TokenClaims{
		jwt.StandardClaims{
			Id:        oldRefreshTokenClaims.StandardClaims.Id, // jti
			Subject:   oldRefreshTokenClaims.StandardClaims.Subject,
			ExpiresAt: refreshTokenExp,
		},
		oldRefreshTokenClaims.Role,
		oldRefreshTokenClaims.Csrf,
		oldRefreshTokenClaims.Type,
	}

	// create a signer for rsa 256
	refreshJwt := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), refreshClaims)

	// generate the refresh token string
	newRefreshTokenString, err = refreshJwt.SignedString(signKey)
	return
}

func updateAuthTokenString(refreshTokenString string, oldAuthTokenString string) (newAuthTokenString, csrfSecret string, err error) {
	refreshToken, err := jwt.ParseWithClaims(refreshTokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})
	_, ok := refreshToken.Claims.(*TokenClaims)
	if !ok {
		err = errors.New("Error reading jwt claims")
		return
	}

	// check if the refresh token has been revoked

	// the refresh token has not been revoked
	// has it expired?

	// nope, the refresh token has not expired
	// issue a new auth token
	authToken, _ := jwt.ParseWithClaims(oldAuthTokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})

	oldAuthTokenClaims, ok := authToken.Claims.(*TokenClaims)
	if !ok {
		err = errors.New("Error reading jwt claims")
		return
	}

	// our policy is to regenerate the csrf secret for each new auth token
	csrfSecret, err = GenerateCSRFSecret()
	if err != nil {
		return
	}

	newAuthTokenString, err = createAuthTokenString(oldAuthTokenClaims.StandardClaims.Subject, oldAuthTokenClaims.Role, csrfSecret)

	return

}
