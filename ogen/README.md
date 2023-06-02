# ogen

https://ogen.dev

## Getting started

Install ogen
```shell
go install -v github.com/ogen-go/ogen/cmd/ogen@latest
```

Generate
```shell
ogen -generate-tests -target gen -clean ../openapi3.yaml
```

Run server
```shell
go run server.go
```

Visit Swagger
```shell
open http://localhost:8000/swagger/
```
