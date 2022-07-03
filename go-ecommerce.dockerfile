FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o go-ecommerce ./

RUN chmod +x /app/go-ecommerce


FROM alpine:latest

RUN mkdir /app

COPY --from=builder  /app/go-ecommerce /app

CMD [ "/app/go-ecommerce" ]