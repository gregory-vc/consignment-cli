FROM golang:1.11.5-alpine as builder

WORKDIR /go/src/github.com/gregory-vc/consignment-cli

COPY . .

RUN go build

FROM alpine:latest
RUN apk --no-cache add ca-certificates iputils bash

RUN mkdir -p /app
WORKDIR /app

COPY --from=builder /go/src/github.com/gregory-vc/consignment-cli/consignment.json .
COPY --from=builder /go/src/github.com/gregory-vc/consignment-cli/consignment-cli .

CMD ["./consignment-cli"]