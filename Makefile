.PHONY: lint
lint:
	go fmt ./...

.PHONY: build
# build: lint docs
build: lint
	GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo -o ./main/mainexec ./

.PHONY: build-docker
build-docker:
	@docker build --build-arg VERSION=$${VERSION:-development} -f Dockerfile -t my-go-app-golang:make .

.PHONY: docker-image
docker-image: build build-docker

.PHONY: docs
docs:
	./swag_cli init .

# Executes integration test using a local Docker Compose stack
test: 
	@docker-compose up -d
	go test ./goapi_test.go
	@docker-compose down


