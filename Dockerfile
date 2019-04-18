FROM node:8.15.1-alpine as assetbuilder

RUN npm i -g yarn@1.12.3

WORKDIR /github.com/overmike/webterminal/js/
ADD js/ /github.com/overmike/webterminal/js/

RUN yarn
RUN yarn clean 
RUN yarn build


FROM golang:1.12 as builder

WORKDIR /github.com/overmike/webterminal/

ADD go.mod /github.com/overmike/webterminal/
ADD go.sum /github.com/overmike/webterminal/
RUN go mod download

ADD . /github.com/overmike/webterminal/

ARG LDFLAGS
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="${LDFLAGS}" -a -installsuffix cgo -o webterminal .




FROM alpine

RUN apk add --no-cache bash

COPY --from=builder /github.com/overmike/webterminal/webterminal /usr/local/bin/
COPY --from=assetbuilder /github.com/overmike/webterminal/js/dist /github.com/overmike/webterminal/js/dist

EXPOSE 8081 50051
ENTRYPOINT ["/usr/local/bin/webterminal"]
CMD ["serve", "-i"]
