extract_podium:
  archive.extracted:
    - name: /opt
    - source: http://192.168.209.130:8081/repository/raw-kj/podium/44/podium.tar.gz
    - source_hash: sha1=48d3e7ac44d951d1e80d8b82a02836ff02820953
    - overwrite: True

  file.directory:
    - name: /opt/podium
    - user: tmwadm
    - group: s3mw
    - dir_mode: 755
    - file_mode: 774
    - recurse:
        - user
        - group
        - mode
  cmd.run:
    - name: find /opt/podium -name "core-env.properties" -print0 | xargs -0 rm -rf