package main

import (
	_ "chatAppServer/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

