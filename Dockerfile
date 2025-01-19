# Build stage
FROM golang:1.22.5 AS builder
WORKDIR /app
COPY . .
RUN go build -o messaging-app .

# Run stage
FROM ubuntu:latest
LABEL authors="can.zanat"
WORKDIR /app/
COPY --from=builder /app/messaging-app ./
COPY ./.config ./.config
RUN chmod +x messaging-app
CMD ["./messaging-app"]
