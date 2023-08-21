FROM golang:1.20.4-bullseye as builder

WORKDIR /go/src/gopetwatch

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY . .

RUN go mod download && go mod verify && go mod tidy

RUN go build -o /go/bin/gopetwatch .

COPY --from=builder /go/bin/gopetwatch /usr/local/bin/gopetwatch

CMD ["gopetwatch"]

