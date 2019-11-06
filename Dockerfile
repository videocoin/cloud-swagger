FROM golang:1.12.4 as builder
WORKDIR /go/src/github.com/videocoin/cloud-swagger/
COPY . .
RUN make build

FROM bitnami/minideb:jessie
WORKDIR /opt/videocoin/bin/
COPY --from=builder /go/src/github.com/videocoin/cloud-swagger/bin/swagger /opt/videocoin/bin/swagger
COPY --from=builder /go/src/github.com/videocoin/cloud-swagger/proto /opt/videocoin/bin/proto
COPY --from=builder /go/src/github.com/videocoin/cloud-swagger/service /opt/videocoin/bin/service
CMD ["/opt/videocoin/bin/swagger"]

