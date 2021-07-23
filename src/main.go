package main

import (
	"moco/config"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
  db := config.SqlConnect()
  defer db.Close()

  router := gin.Default()

  router.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "hello world",
    })
})

  router.Run()
}


