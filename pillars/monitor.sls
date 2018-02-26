beacons:
  inotify:
    disable_during_state_run: True
    /srv/conflictrepo:
      mask:
        - create
        - close_write
        - modify
        - delete
        - delete_self
      recurse: True
      exclude:
        - /srv/conflictrepo/.git
    coalesce: True
