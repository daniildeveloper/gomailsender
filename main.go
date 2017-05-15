package main

import (
	_ "mailsender/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

