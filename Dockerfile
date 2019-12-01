FROM golang:1.13

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . /app

CMD go run cmd/main.go
