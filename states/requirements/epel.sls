epel-repo:
  pkgrepo.managed:
    - name: epel
    - humanname: Extra Packages for Enterprise Linux $releasever - $basearch
    - mirrorlist: https://mirrors.fedoraproject.org/metalink?repo=epel-{{ grains['osmajorrelease'] }}&arch=$basearch
    - failovermethod: priority
    - exclude: nginx*,php*
    - gpgcheck: 1
    - gpgkey: file:///etc/pki/rpm-gpg/RPM-GPG-KEY-EPEL-{{ grains['osmajorrelease'] }}
    - disabled: 0
