FROM golang:1.18.0-alpine

RUN apk update && apk add --no-cache git

WORKDIR /go/src/app

COPY . .

RUN go mod download && go mod verify

RUN go mod tidy && go mod vendor

RUN chmod +x /go/src/app

ENV WAIT_VERSION 2.7.3
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait

RUN go build -o /server cmd/main.go

EXPOSE 3000

CMD ["/server"]