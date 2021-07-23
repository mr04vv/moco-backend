package main

import (
	"moco/config"
	"moco/db/model"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
  db := config.SqlConnect()
  defer db.Close()

  router := gin.Default()

  router.GET("/users", func(c *gin.Context) {
    c.JSONP(http.StatusOK, gin.H{
      "message": "ok",
      "data": getUsers((db)),
    })
  })

  router.GET("/messages", func(c *gin.Context) {
    c.JSONP(http.StatusOK, gin.H{
      "message": "ok",
      "data": getMessages((db)),
    })
  })

  router.Run()
}


func getUsers(db *gorm.DB) ([]model.User){
  var users []model.User
  db.Find(&users)
  return users
}

func getMessages(db *gorm.DB) ([]model.Message){
  var messages []model.Message
  db.Find(&messages)
  return messages
}
