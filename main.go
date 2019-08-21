package main

import (
	_ "github.com/Shows/init"
	_ "github.com/Shows/models"
	_ "github.com/Shows/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/episode_files", "static/episode_file")
	beego.Run()

}
