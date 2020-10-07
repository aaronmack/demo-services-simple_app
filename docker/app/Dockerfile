# Development environment
# -----------------------

FROM golang:1.15-alpine
WORKDIR /demo-services-simple_app

# ENV HTTP_PROXY="http://127.0.0.1:10809"
# ENV HTTPS_PROXY="http://127.0.0.1:10809"

RUN apk update && apk add --no-cache gcc musl-dev git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/app ./cmd/app \
    && go build -ldflags '-w -s' -a -o ./bin/migrate ./cmd/migrate \
    && chmod +x /demo-services-simple_app/docker/app/bin/*

EXPOSE 8080