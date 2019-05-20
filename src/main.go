package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"test/apitest"
	"test/src/app"
)

func main() {
	fmt.Print("hello,world")
	apitest.Test()
	app.CeshiGo()
	beego.Run()
}
