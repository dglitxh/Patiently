# syntax=docker/dockerfile:1
FROM golang:1.16-buster as build

WORKDIR /app

COPY go.mod ./
COPY go.sum  ./

RUN go mod download

COPY ./ ./

RUN go build -o /patiently .


FROM alpine

RUN apk --no-cache add ca-certificates

WORKDIR /

COPY --from=build /patiently /patiently 

EXPOSE 8080

ENTRYPOINT ["/patiently"]
