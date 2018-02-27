extract_podium:
  archive.extracted:
    - name: /tmp/
    - source: http://192.168.209.130:8081/repository/raw-kj/podium/44/podium.tar.gz
    - source_hash: sha1=48d3e7ac44d951d1e80d8b82a02836ff02820953
    - user: tmwadm
    - group: s3mw
    - overwrite: True

podiumwebapps:
  file.copy:
      - name: /opt/podium
      - source: /tmp/podium
      - user: tmwadm
      - group: s3mw
      - dir_mode: 755
      - exclude_pat: core-env.properties
      - recurse:
          - user
          - group
          - mode
      - file_mode: 774
      - keey_symlinks: true
      - include_empty: true
      - require:
          - archive: extract_podium
