package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type MailController struct {
	beego.Controller
}

func (c *MailController) send(to string, subject string, mail string) {
	fmt.Printf("Mail to %s. Subject: %s is sended.", to, subject)
}

func init() {

}
