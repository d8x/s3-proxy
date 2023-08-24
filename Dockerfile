FROM golang:1.21-alpine3.18 AS builder
WORKDIR /opt/
RUN apk update
RUN apk add git make

ADD . .
RUN make build

FROM alpine:3.18
WORKDIR /opt/
COPY --from=builder /opt/sgw ./sgw

EXPOSE 7312
ENTRYPOINT ["./sgw"]


