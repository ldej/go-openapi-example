# oapi-codegen

https://github.com/deepmap/oapi-codegen

## Getting started

Install oapi-codegen
```shell
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
```

Generate
```shell
oapi-codegen -config server.cfg.yaml ../openapi3.yaml
```

Run server
```shell
go run server.go
```

Visit Swagger
```shell
open http://localhost:8000/swagger/
```
