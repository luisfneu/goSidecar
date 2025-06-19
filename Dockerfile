FROM golang:1.24 as builder

WORKDIR /app
COPY . .
RUN go mod init sidecar && go build -o sidecar .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/sidecar .
ENTRYPOINT ["./sidecar"]
