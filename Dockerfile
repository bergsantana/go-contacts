# Use official Golang image
FROM golang:1.24.3-alpine

RUN apk add --no-cache git sqlite sqlite-dev
RUN apk add build-base

ENV CGO_ENABLED=1

WORKDIR /app

 
COPY . .

RUN go mod tidy

# Expose API port
EXPOSE 3000

# #  
# RUN go run cmd/main.go seed

# CMD ["go", "run", "cmd/main.go"]


