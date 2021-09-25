FROM golang:1.17-alpine AS builder
WORKDIR /src
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app

FROM scratch
ENV GIN_MODE=release
COPY --from=builder /app /app
EXPOSE 8080
ENTRYPOINT [ "/app" ]
