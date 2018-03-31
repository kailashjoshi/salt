git:
  pkg.removed:
    - name: git
    {% if 'minion.*' in grains['fqdn'] %}
    - verision: 1.8.3.1-11.el7
    {% endif %}
    - ignore_epoch: True