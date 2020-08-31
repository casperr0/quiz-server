FROM golang:alpine

ENV GO111MODULE=on
WORKDIR /service/
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build -o ./build/ ./...
CMD ["./build/cmd"]