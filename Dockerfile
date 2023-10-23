FROM golang:alpine

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest

COPY . .

EXPOSE 8080

CMD ["air","server","--port","8080"]
