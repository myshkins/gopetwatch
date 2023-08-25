FROM golang:1.20.4-bullseye as builder

WORKDIR /go/src/app

COPY . .

RUN go mod download && go mod verify && go mod tidy

RUN go build -o /go/bin/gopetwatch .

FROM ubuntu:jammy

RUN useradd -rm -d /home/api -s /bin/bash -u 1000 iceking

WORKDIR /home/app

COPY templates .

COPY --from=builder /go/bin/gopetwatch .

CMD ["gopetwatch"]

