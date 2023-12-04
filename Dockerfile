FROM golang:1.21 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 go build -o main

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main /app/main
RUN ls -la /app  # List the contents of /app in the final stage
CMD ["/app/main"]
