package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// User is
type User struct {
	gorm.Model
	id       int
	name     string
	email    string
	password string
	role     string
}

func (u *User) say() {
	fmt.Print("say")
}
