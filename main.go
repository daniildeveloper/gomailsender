package main

import (
	_ "mailsender/models"
	_ "mailsender/routers"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// In order to handle time.Time, you need to include parseTime as a parameter.
	db, err := gorm.Open("mysql", "root@/mail_sender?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)

	}
	u := db.HasTable("users")
	if !u {
		// db.AutoMigrate(&models.)
	}
	// beego.Run()
}

func init() {

}
