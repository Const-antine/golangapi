FROM golang:latest AS gobuilder

WORKDIR /go/src/goapi

COPY ./ .

RUN go get -d -v ./...

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/src/goapi/main/mainexec ./


FROM scratch AS prod

ENV APP_BUILD_PATH="/var/app" \
    APP_BUILD_NAME="mainexec"

WORKDIR ${APP_BUILD_PATH}

COPY --from=gobuilder /go/src/goapi/main/mainexec /var/app/mainexec

ENTRYPOINT ["/var/app/mainexec"]
CMD ""
