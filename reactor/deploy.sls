deploy_mods:
  cmd.state.sls:
    - tgt: {{ data.id }}
    - kwarg:
        mods: {{ data.data.module }}