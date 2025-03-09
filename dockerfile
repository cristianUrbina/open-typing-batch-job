FROM golang:1.24.1-bullseye AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o my-api cmd/api/main.go

FROM golang:1.24.1-bullseye

WORKDIR /root/

COPY --from=builder /app/my-api .

EXPOSE 8080

CMD ["./my-api"]
