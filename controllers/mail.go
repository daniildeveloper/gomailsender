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

// Get is a function
func (c *MailController) Get() {
	c.Data["Website"] = "beego.daniiltserin.ru"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MailController) subscribe() {
	c.Data["id"] = c.Ctx.Input.Param(":id")
}
