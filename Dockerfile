FROM golang:1.18-alpine as builder

RUN mkdir /app

ADD . /app/

WORKDIR /app
RUN go build -o main .


FROM alpine:latest
COPY --from=builder /app/main /app/main
COPY --from=builder /app/config.yaml /app/config.json
WORKDIR /app

ENTRYPOINT ./main -c config.yaml
