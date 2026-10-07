# gopu

Simple http service written in go

### install

To set it up in kubernetes:
```sh
helm install gopu chart/ -n gopu --create-namespace
helm upgrade --install gopu chart/ -n gopu
```
