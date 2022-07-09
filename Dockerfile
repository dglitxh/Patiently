# syntax=docker/dockerfile:1
FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod  ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /patiently


FROM alpine:LATEST

WORKDIR /

COPY --from=build /patiently /patiently

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/patiently"]
