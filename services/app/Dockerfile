FROM golang:1.20.4-bullseye as builder

WORKDIR /go/src/app

COPY . .

RUN go mod download && go mod verify && go mod tidy

RUN go build -o /go/bin/gopetwatch .

FROM ubuntu:jammy

RUN useradd -rm -d /home/api -s /bin/bash -u 1000 iceking

WORKDIR /home/app

ARG GIN_USER_ID
ARG GIN_PW

ENV GIN_USER_ID=$GIN_USER_ID
ENV GIN_PW=$GIN_PW

COPY templates templates

COPY --from=builder /go/bin/gopetwatch .

CMD ["./gopetwatch"]

