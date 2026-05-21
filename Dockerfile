FROM harbor.bevz.net/dockerhub-proxy/library/golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o helloworld

FROM harbor.bevz.net/dockerhub-proxy/library/alpine:latest
WORKDIR /app
COPY --from=builder /app/helloworld .
CMD ["./helloworld"]
