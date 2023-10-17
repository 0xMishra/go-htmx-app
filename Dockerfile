FROM golang

WORKDIR /usr/src/app

COPY . .

RUN go build

CMD ["./go-todo-app"]
