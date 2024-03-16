generate-api-go:
    go generate ./backend/...

install:
    go get ./backend/...
    npm i querystring-es3
    npm install -D @sap-cloud-sdk/openapi-generator

generate-api-js:
    docker run --rm -v .:/local openapitools/openapi-generator-cli /bin/bash \
    -c "java -jar /opt/openapi-generator/modules/openapi-generator-cli/target/openapi-generator-cli.jar generate -i /local/api.yaml -g javascript -o /local/frontend/out; mv /local/frontend/out/src/* /local/frontend/src/api/; rm -rf /local/frontend/out"
