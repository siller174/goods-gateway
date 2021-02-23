FROM        golang:1.13-alpine3.10 as build

WORKDIR     /build

ENV         GOBIN=/build/bin

ADD         . /build

RUN         apk update && apk add --no-cache git

RUN         go install -mod=vendor /build/cmd/goodsGateway/...

FROM		golang:1.13-alpine3.10

WORKDIR     /opt/goodsGateway

COPY        --from=build /build/bin/goodsGateway /opt/goodsGateway
COPY        --from=build /build/configs/goodsGateway/goodsGateway_prod.properties /opt/goodsGateway/goodsGateway.properties

ENTRYPOINT ["/opt/goodsGateway/goodsGateway"]

EXPOSE      8080