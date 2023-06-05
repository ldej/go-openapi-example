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

Visit Swagger UI
```shell
open http://localhost:8000/swagger/
```

Use `test` for the `X-Api-Key` header or for `Authorize` button in Swagger UI
