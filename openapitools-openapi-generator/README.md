# openapi-generator

https://openapi-generator.tech/

## Getting started

Install openapi-generator-cli
```shell
git clone https://github.com/openapitools/openapi-generator
cd openapi-generator
mvn clean package
```

Note: the go-server generators have been updated recently, and are therefore not part of the latest pre-built binaries. Therefore, it's best to build them yourself

Generate
```shell
java -jar /path/to/openapi-generator-cli.jar generate -i ../openapi3.yaml -g go-server -o gen --enable-post-process-file --git-user-id=ldej --git-repo-id=go-openapi-example/openapitools-openapi-generator/gen --additional-properties=router=chi,outputAsLibrary=true,onlyInterfaces=true
```
