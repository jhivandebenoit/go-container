# syntax=docker/dockerfile:1

FROM golang:latest AS builder
WORKDIR /app
COPY . ./
RUN CGO_ENABLED=0 go build -a -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/app ./
EXPOSE 8080
CMD ["./app"]
