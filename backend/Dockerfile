FROM golang:1.24.4-alpine AS builder

WORKDIR /app

RUN go install github.com/google/wire/cmd/wire@latest && \
  go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/server .

FROM alpine:latest AS server

WORKDIR /app

COPY --from=0 /app/server /app/server

EXPOSE 8080

CMD ["/app/server"]
