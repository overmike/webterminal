FROM golang:1.10 as builder

RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/overmike/webterminal/

ADD Gopkg.lock /go/src/github.com/overmike/webterminal/
ADD Gopkg.toml /go/src/github.com/overmike/webterminal/
RUN dep ensure -v --vendor-only

ADD . /go/src/github.com/overmike/webterminal/


ARG LDFLAGS
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="${LDFLAGS}" -a -installsuffix cgo -o webterminal .



FROM alpine

RUN apk add --no-cache bash

COPY --from=builder /go/src/github.com/overmike/webterminal/webterminal /usr/local/bin/

EXPOSE 8081 50051
ENTRYPOINT ["/usr/local/bin/webterminal"]
CMD ["serve"]
