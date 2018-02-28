epel_rpm_install:
  pkg.installed:
    - source:
        - epel: http://dl.fedoraproject.org/pub/epel/7/x86_64/Packages/e/epel-release-7-11.noarch.rpm
    - allow_updates: True