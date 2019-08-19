package main

import (
	"github.com/astaxie/beego"
	_ "myFirstProject/init"
	_ "myFirstProject/models"
	_ "myFirstProject/routers"
)

func main() {

	beego.Run()

}
