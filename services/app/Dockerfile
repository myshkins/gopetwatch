FROM golang:1.20.4-bullseye

WORKDIR /usr/local/go/src/gopetwatch

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./

RUN go mod download && go mod verify && go mod tidy

COPY . .

RUN go build -o /usr/local/bin/gopetwatch

CMD go install && go run .
