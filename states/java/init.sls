download_java:
  cmd.run:
    - name: 'wget -c --header "Cookie: oraclelicense=accept-securebackup-cookie" http://download.oracle.com/otn-pub/java/jdk/8u131-b11/d54c1d3a095b4ff2b6607d096fa80163/jdk-8u131-linux-x64.rpm'
    - cwd: /tmp
    - unless: java -version

install_java8:
  cmd.run:
    - name: yum -y install /tmp/jdk-8u131-linux-x64.rpm
    - unless: java -version