# STAGE 1
FROM golang:1.17-alpine AS builder

# Define build env
ENV GOOS linux
ENV CGO_ENABLED 0

# create a working directory inside the image
WORKDIR /usr/src/app

# copy Go modules and dependencies to image
# download Go modules and dependencies
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# copy directory files i.e all files ending with .go
COPY . .

# Build app
RUN go build -v -o /usr/local/bin/chatBotApp

# STAGE 2

FROM alpine:latest

WORKDIR /
# Add certificates
RUN apk add --no-cache ca-certificates

COPY --from=builder /usr/local/bin/chatBotApp ./

EXPOSE 8080

CMD ["./chatBotApp"]