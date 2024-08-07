FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o publisher

CMD ["./publisher"]
