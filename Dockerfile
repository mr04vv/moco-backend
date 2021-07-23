FROM golang:latest

RUN mkdir /app
WORKDIR /app
COPY ./src .

RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/jinzhu/gorm
RUN go get github.com/rubenv/sql-migrate/...
RUN go get -u github.com/cosmtrek/air
RUN go build -o /go/bin/air github.com/cosmtrek/air
CMD ["air", "-c", ".air.toml"]
