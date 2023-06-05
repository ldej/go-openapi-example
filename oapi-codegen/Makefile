all: install generate run

install:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.13.0

generate:
	oapi-codegen -config server.cfg.yaml ../openapi3.yaml

run:
	go run server.go

visit:
	open http://localhost:8000/swagger/
