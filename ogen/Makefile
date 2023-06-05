all: install generate run

install:
	go install -v github.com/ogen-go/ogen/cmd/ogen@v0.68.4

generate:
	ogen -generate-tests -target gen -clean ../openapi3.yaml

run:
	go run server.go

visit:
	open http://localhost:8000/swagger/
