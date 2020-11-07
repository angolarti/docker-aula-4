FROM golang:1.15

WORKDIR /go/src/checkout

COPY . .


RUN go get github.com/streadway/amqp \
    && go get github.com/joho/godotenv \
    && go get github.com/satori/go.uuid \
    && go get github.com/wesleywillians/go-rabbitmq/queue \
    && GOOS=linux go build

EXPOSE 9090

ENTRYPOINT ["./checkout"]