package main

import (
	"errors"
	"fmt"
	"moco/config"
	"moco/db/model"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func seeds(db *gorm.DB) error {
	var u model.User
	err := db.First(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			for i := 0; i < 10; i++ {
				user := model.User{Name: "hoge"+strconv.Itoa(i), ProfileURL: "https://tpc.googlesyndication.com/simgad/8498538992993752138"}
				if err := db.Create(&user).Error; err != nil {
					fmt.Printf("%+v", err)
				}
			}
		}
	}

	var m model.Message
	err = db.First(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			for i := 0; i < 50; i++ {
				message := model.Message{Message: "hogehogehogehogehogehoge", UserID: 1}
				if err := db.Create(&message).Error; err != nil {
					fmt.Printf("%+v", err)
				}
			}
		}
	}

    return nil
}
 
func main() {
    db := config.SqlConnect()
    defer db.Close()
    if err := seeds(db); err != nil {
        fmt.Printf("%+v", err)
        return
    }
}
