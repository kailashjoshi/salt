fix_git:
  cmd.state.sls:
    - tgt: {{ data.id }}
    - arg:
        - gitreset