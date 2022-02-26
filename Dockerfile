FROM golang:1.17-alpine

WORKDIR /usr/src/app

RUN apk add curl

COPY ./src/go.mod ./src/go.sum ./
RUN go mod download && go mod verify

COPY ./src .
RUN go build -v -o /usr/local/bin/app ./...

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

CMD ["app"]
