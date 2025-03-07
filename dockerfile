FROM golang:1.24.1-bullseye AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o my-api cmd/api/main.go

ARG DB_USER
ARG DB_PASSWORD
ARG DB_NAME
ARG DB_HOST
ARG DB_PORT

ENV DB_USER=$DB_USER \
    DB_PASSWORD=$DB_PASSWORD \
    DB_NAME=$DB_NAME \
    DB_HOST=$DB_HOST \
    DB_PORT=$DB_PORT

ENV DB_USER=${DB_USER}
ENV DB_PASSWORD=${DB_PASSWORD}
ENV DB_NAME=${DB_NAME}
ENV DB_HOST=${DB_HOST}
ENV DB_PORT=${DB_PORT}

FROM golang:1.20-buster

WORKDIR /root/

COPY --from=builder /app/my-api .

EXPOSE 8080

CMD ["./my-api"]
