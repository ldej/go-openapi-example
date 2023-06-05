all: install generate run

install:
	go install github.com/contiamo/openapi-generator-go@latest

generate:
	openapi-generator-go generate --spec ./openapi3-modified.yaml --output ./gen/

run:
	go run server.go

visit:
	open http://localhost:8000/swagger/
