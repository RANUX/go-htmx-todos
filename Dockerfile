# syntax=docker/dockerfile:1
FROM golang:1.22.1 as builder

LABEL org.opencontainers.image.authors="Alexandr Razzhivin"

RUN --mount=type=cache,target=/var/cache/apt \
    apt-get update && apt-get install -y build-essential

WORKDIR /app

# copy and download dependencies
COPY . .

RUN go mod tidy
RUN go mod verify
RUN go mod download

# CGO_ENABLED=0 - needs to build statically (https://golang.org/src/cmd/cgo/doc.go)
# GOOS=linux - need to build for linux only
RUN CGO_ENABLED=0 GOOS=linux go build -o todo-app cmd/main.go

# Final stage
FROM scratch AS final
WORKDIR /app
COPY --from=builder /app/todo-app .
COPY .env .env
COPY --from=builder /app/public ./public
EXPOSE 3000
CMD ["./todo-app"]