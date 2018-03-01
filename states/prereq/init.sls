python_selinux_pkg:
  pkg.installed:
    - name: policycoreutils-python

disable_selinux:
  selinux.mode:
    - name: permissive