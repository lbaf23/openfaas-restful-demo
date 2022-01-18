OpenFaaS Restful Demo
===


- OpenFaaS
- Restful
- golang


```bash
faas-cli template store pull golang-middleware
```

```bash
faas-cli up -f record.yml
```

or

```bash
faas-cli build -f record.yml
faas-cli push -f record.yml
faas-cli deploy -f record.yml
```