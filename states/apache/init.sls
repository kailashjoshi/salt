install_apache:
  pkg.installed:
    - name: httpd

start_apache:
  service.running:
    - name: httpd
    - enable: True

welcome_page:
  file.managed:
    - name: /var/www/html/index.html
    - source: salt://apache/files/index.html
