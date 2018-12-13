FROM golang:1.11 as builder

WORKDIR /github.com/overmike/webterminal/

ADD . /github.com/overmike/webterminal/

RUN go mod download

ARG LDFLAGS
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="${LDFLAGS}" -a -installsuffix cgo -o webterminal .




FROM alpine

RUN apk add --no-cache bash

COPY --from=builder /github.com/overmike/webterminal/webterminal /usr/local/bin/

EXPOSE 8081 50051
ENTRYPOINT ["/usr/local/bin/webterminal"]
CMD ["serve"]
