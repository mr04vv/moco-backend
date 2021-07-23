package config

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

func SqlConnect() (database *gorm.DB) {
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
  