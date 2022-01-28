OpenFaaS Restful Demo
===


- OpenFaaS
- Restful
- golang


pull template

```bash
faas-cli template store pull golang-middleware
```

set password

```bash
vim openfaas-db-password
```

write db password

```bash
faas-cli secret create openfaas-db-password --from-file openfaas-db-password
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

get pod status
```bash
kubectl get pods -n openfaas-fn
```

print pod logs
```bash
kubectl logs ${pod name} -n openfaas-fn
```
