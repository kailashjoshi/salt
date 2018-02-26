execute_git_sscript:
  cmd.script:
    - name: salt://gitreset/files/gitreset.sh
    - shell: /bin/bash
    - cwd: /srv/conflictrepo