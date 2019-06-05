FROM golang:1.12

WORKDIR /go/src/app

COPY . .

RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/gin-gonic/gin
# RUN go get -u github.com/go-xorm/xorm

