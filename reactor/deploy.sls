deploy:
  cmd.state.sls:
    - tgt: {{ data.id }}
    - kwarg:
        mods: {{ data.data.module }}