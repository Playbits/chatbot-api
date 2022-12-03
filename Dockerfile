#Stage 1
FROM golang:1.18-alpine AS builder

WORKDIR /chatbot

ENV CGO_ENABLED 0
ENV GOPATH /go
ENV GOCACHE /go-build

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod/cache \
    go mod download

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod/cache \
    --mount=type=cache,target=/go-build \
    go build -o bin/backend main.go

CMD ["/chatbot/bin/backend"]

# Stage 2
FROM builder AS dev-envs

RUN apk update && apk add git

RUN addgroup -S docker && \
    adduser -S --shell /bin/bash --ingroup docker vscode

# install Docker tools (cli, buildx, compose)
COPY --from=gloursdocker/docker / /

CMD ["go", "run", "main.go"]

#Stage 3
FROM alpine:3.12
EXPOSE 8500
COPY --from=build /chatbot/bin/backend /usr/local/bin/backend
CMD ["/usr/local/bin/backend"]
