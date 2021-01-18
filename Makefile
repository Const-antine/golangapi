.PHONY: lint
lint:
	go fmt ./...

.PHONY: build
# build: lint docs
build: lint
	GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo -o ./main/mainexec ./main

.PHONY: build-docker
build-docker:
	@docker build --build-arg VERSION=$${VERSION:-development} -f Dockerfile -t my-go-app-golang:make .

.PHONY: docker-image
docker-image: build build-docker

.PHONY: docs
docs:
	swag init --output ./main/docs --dir ./main  --generalInfo ./main/test.go

# Executes integration test using a local Docker Compose stack
test:
	@docker-compose up -d
	eval $$(cat .env) go test ./...
	@docker-compose down

#cd ./app/http/ && docker run -t -i -p 8246:8080 -e SWAGGER_JSON=/var/specs/swagger.yaml  -v $PWD/docs:/var/specs swaggerapi/swagger-ui