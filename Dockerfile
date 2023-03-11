# First stage - build the Go binary
FROM golang:1.17-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download && go build -o app

# Second stage - final image with only the binary
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/app /app

EXPOSE 8080
ENTRYPOINT ["/app/app"]
