FROM golang:1.15

ENV GO111MODULE=on
WORKDIR /service/
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build -o ./build/quiz-server ./...
CMD ["./build/quiz-server"]