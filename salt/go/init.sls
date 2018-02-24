extract_go_1.10:
  archive.extracted:
    - name: /usr/local/
    - source: http://192.168.209.130:8081/repository/raw-kj/golang/1.10-amd64/go1.10.linux-amd64.tar.gz
    - source_hash: sha1=693e2dfd99ef2d608b41c450efd9d277bf2a7f2f
    - user: root
    - group: root

go_lang_env_variables:
  file.managed:
    - name: /etc/profile.d/go.sh
    - source: salt://go/files/go.sh