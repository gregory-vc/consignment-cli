FROM golang:1.11.5 as builder

WORKDIR /go/src/consignment-cli

COPY . .

RUN go build

FROM alpine:latest

RUN mkdir -p /app
WORKDIR /app

COPY --from=builder /go/src/consignment-cli/consignment.json .
COPY --from=builder /go/src/consignment-cli/consignment-cli .

CMD ["./consignment-cli"]