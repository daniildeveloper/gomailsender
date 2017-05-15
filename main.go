package main

import (
	_ "mailsender/routers"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User
type User struct {
	Id       int
	name     string
	email    string
	password string
	role     string
}

func main() {
	// In order to handle time.Time, you need to include parseTime as a parameter.
	db, err := gorm.Open("mysql", "root@/mail_sender?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)

	}
	db.CreateTable("users")
	// beego.Run()
}

func init() {

}
