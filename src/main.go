package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
  db := sqlConnect()
  defer db.Close()

  router := gin.Default()

  router.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "hello world",
    })
})

  router.Run()
}

func sqlConnect() (database *gorm.DB) {
  DBMS := "mysql"
  USER := "root"
  PASS := "root"
  PROTOCOL := "tcp(db:3306)"
  DBNAME := "moco"

  CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

  count := 0
  db, err := gorm.Open(DBMS, CONNECT)
  if err != nil {
    for {
      if err == nil {
        fmt.Println("")
        break
      }
      fmt.Print(".")
      time.Sleep(time.Second)
      count++
      if count > 180 {
        fmt.Println("")
        fmt.Println("DB接続失敗")
        panic(err)
      }
      db, err = gorm.Open(DBMS, CONNECT)
    }
  }
  fmt.Println("DB接続成功")

  return db
}
