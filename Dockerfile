FROM golang:alpine
WORKDIR /go/src/github.com/step/sauron/

ADD . .
RUN apk update && apk add --no-cache git ca-certificates make && go get ./... && make sauron

FROM alpine
WORKDIR /app
COPY --from=0 /go/src/github.com/step/sauron/bin/sauron ./sauron
COPY --from=0 /go/src/github.com/step/sauron/config.toml .
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs
ENTRYPOINT ["sh","-c", "/app/sauron -redis-address $REDIS_ADDRESS -redis-db $REDIS_DB -log-filename $LOG_FILE_NAME -log-path $LOG_FILE_PATH"]