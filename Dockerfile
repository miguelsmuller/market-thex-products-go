FROM golang:1.21.3

WORKDIR /app

COPY ./src /app
COPY ./settings /app

RUN go build -o api_thex_products_bin

EXPOSE 3002

ENTRYPOINT ["./api_thex_products_bin"]
