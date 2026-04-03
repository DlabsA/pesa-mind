FROM golang:1.26-alpine as builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o pesa-mind ./cmd/api

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/pesa-mind .
COPY .env.example .env
EXPOSE 8080
CMD ["./pesa-mind"]

