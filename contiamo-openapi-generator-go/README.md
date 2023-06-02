# openapi-generator-go

https://github.com/contiamo/openapi-generator-go

### Getting started

Install openapi-generator-go
```shell
go install github.com/contiamo/openapi-generator-go@latest
```

Generate
```shell
openapi-generator-go generate --spec ./openapi3-modified.yaml --output ./gen/
```

NOTE: This generator requires a modified openapi specification to generate the correct router. Each endpoint contains a 'x-handler-group' field.

Run server
```shell
go run server.go
```

Visit Swagger
```shell
open http://localhost:8000/swagger/
```
