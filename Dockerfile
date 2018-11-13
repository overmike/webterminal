FROM golang:1.10 as builder

WORKDIR /go/src/github.com/overmike/webterminal/
ADD . /go/src/github.com/overmike/webterminal/

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o webterminal .

FROM alpine

RUN apk add --no-cache bash

COPY --from=builder /go/src/github.com/overmike/webterminal/webterminal /usr/local/bin/

ENTRYPOINT ["/usr/local/bin/webterminal"]
CMD ["serve"]
