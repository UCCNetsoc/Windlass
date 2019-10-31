FROM golang:1.13 AS dev

WORKDIR /windlass

RUN go get github.com/go-task/task/cmd/task \
    github.com/derekparker/delve/cmd/dlv

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . . 

RUN go install github.com/UCCNetworkingSociety/Windlass/cmd/windlass

RUN go mod vendor

CMD [ "go", "run", "cmd/windlass/main.go" ]

FROM alpine

WORKDIR /bin

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

COPY --from=dev /go/bin/windlass ./windlass

CMD [ "windlass" ]