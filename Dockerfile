FROM golang:1.17-alpine AS builder

COPY . /app

WORKDIR /app

# Add gcc
RUN apk add build-base

RUN go mod download && \
    go env -w GOFLAGS=-mod=mod && \
    go get . && \
    go build -v .

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/HarmonoidProxy ./HarmonoidProxy

EXPOSE 80
CMD ./HarmonoidProxy