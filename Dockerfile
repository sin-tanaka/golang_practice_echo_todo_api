FROM golang:latest

WORKDIR /go/src/github.com/sin_tanaka/echo_todo_crud
COPY . .
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init
RUN dep ensure
ENV PORT 8888
CMD go run server.go

