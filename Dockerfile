FROM golang:1.17-alpine

WORKDIR /go_xstore

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY api ./api

WORKDIR /go_xstore/api

RUN go build -o ./xstore

EXPOSE 8080

CMD [ "./xstore" ]
