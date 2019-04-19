FROM golang:1.12 AS dev

WORKDIR /windlass

ENV GO111MODULES=on

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . . 

RUN go install

CMD [ "go" "run" "main.go" ]

FROM alpine

WORKDIR /bin

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

COPY --from=dev /go/bin/Windlass ./Windlass

CMD [ "Windlass" ]