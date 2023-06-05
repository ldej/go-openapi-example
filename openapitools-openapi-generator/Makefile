all: generate run

generate:
	java -jar ./openapi-generator-cli-7.0.0-SNAPSHOT.jar generate -i ../openapi3.yaml -g go-server -o gen --enable-post-process-file --git-user-id=ldej --git-repo-id=go-openapi-example/openapitools-openapi-generator/gen -c config.yaml
run:
	go run server.go

visit:
	open http://localhost:8000/swagger/
