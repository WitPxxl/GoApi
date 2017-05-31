FROM golang:1.8.1

WORKDIR /go/src/github.com/Witpxxl/GoApi

RUN go get -u github.com/kardianos/govendor

CMD ["go", "run", "main.go"]