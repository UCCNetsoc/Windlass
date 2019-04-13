FROM golang:1.12 AS dev

ENV GO111MODULES=on

COPY go.mod .
COPY go.sum .

RUN go mod download

FROM golang:1.12-alpine

