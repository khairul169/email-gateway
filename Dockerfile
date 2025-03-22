# build stage
FROM golang:1.24-alpine AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o email-gateway main.go

# deployment stage
FROM alpine:latest
WORKDIR /app

COPY --from=build /app/email-gateway .
EXPOSE 5000

CMD ["./email-gateway"]
