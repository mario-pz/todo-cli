FROM golang:1.20.3-alpine3.17

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . . 

RUN go build -o todo-cli

EXPOSE 8080

ENTRYPOINT [ "./todo-cli" ]