FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o helloworld

FROM alpine
WORKDIR /app
COPY --from=builder /app/myapp .
CMD ["./helloworld"]
