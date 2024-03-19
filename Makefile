generate:
    docker run --rm -v .:/local openapitools/openapi-generator-cli /bin/bash \
    -c "rm -rf /local/frontend/src/api/*; java -jar /opt/openapi-generator/modules/openapi-generator-cli/target/openapi-generator-cli.jar generate -i /local/api.yaml -g javascript -o /local/frontend/out; mv /local/frontend/out/src/* /local/frontend/src/api/; rm -rf /local/frontend/out"
    cd backend
    go generate ./...

build:
    docker-compose up --build

run:
    docker-compose up
