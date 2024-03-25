# syntax=docker/dockerfile:1
FROM golang:1.22.1-bookworm

LABEL org.opencontainers.image.authors="Alexandr Razzhivin"

RUN --mount=type=cache,target=/var/cache/apt \
    apt-get update && apt-get install -y build-essential

WORKDIR /usr/src/app

# copy and download dependencies
COPY . .

RUN go mod tidy
RUN go mod verify
RUN go mod download

# expose port
EXPOSE 3000

# run app: cmd/main.go
CMD ["go", "run", "cmd/main.go"]