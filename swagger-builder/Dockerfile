FROM alpine:edge


ARG SWAG_VER=1.7.0


WORKDIR /root/

RUN apk update && apk upgrade && apk add wget

RUN wget https://github.com/swaggo/swag/releases/download/v${SWAG_VER}/swag_${SWAG_VER}_Linux_aarch64.tar.gz && tar xfz swag_${SWAG_VER}_Linux_aarch64.tar.gz

CMD ["./swag"]
