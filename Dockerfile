FROM golang:latest AS gobuilder

WORKDIR /go/src/goapi

COPY ./ .

RUN go get -d -v ./...

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/src/goapi/mainexec ./main/main.go

RUN chmod +x /go/src/goapi/mainexec



FROM scratch AS prod

ENV APP_BUILD_PATH="/var/app" \
    APP_BUILD_NAME="mainexec"
WORKDIR ${APP_BUILD_PATH}
COPY --from=gobuilder /go/src/goapi/mainexec /var/app/mainexec
EXPOSE 8080
CMD ["/var/app/mainexec"]
