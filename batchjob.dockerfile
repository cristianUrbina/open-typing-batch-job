FROM golang:1.24.1-bullseye AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o batchjob cmd/batchjob/main.go

FROM golang:1.24.1-bullseye

WORKDIR /root/

COPY --from=builder /app/batchjob .

EXPOSE 8080

CMD ["./batchjob"]
