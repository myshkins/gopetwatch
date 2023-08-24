FROM golang:1.20.4-bullseye as builder

WORKDIR /go/src/app

COPY . .

RUN go mod download && go mod verify && go mod tidy

RUN go build -o /go/bin/gopetwatch .

FROM ubuntu:jammy

COPY --from=builder /go/bin/gopetwatch /usr/local/bin/gopetwatch

CMD ["gopetwatch"]

