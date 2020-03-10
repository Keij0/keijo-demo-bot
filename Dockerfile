FROM golang:alpine

RUN apk add git

WORKDIR /go/src/app

COPY . .

RUN go get gopkg.in/tucnak/telebot.v2

RUN go get -d -v ./...

RUN go install -v ./...

CMD ["app"]

