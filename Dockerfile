FROM golang:1.24-alpine AS build
WORKDIR /go/src/github.com/utilitywarehouse/csi-nodeinfo-proxy
COPY . /go/src/github.com/utilitywarehouse/csi-nodeinfo-proxy
ENV CGO_ENABLED=0
RUN \
  apk --no-cache add git upx \
    && go get -t ./... \
    && go build -ldflags='-s -w' -o /csi-nodeinfo-proxy . \
    && upx /csi-nodeinfo-proxy

FROM alpine:3.18
COPY --from=build /csi-nodeinfo-proxy /csi-nodeinfo-proxy
ENTRYPOINT [ "/csi-nodeinfo-proxy" ]
