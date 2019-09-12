package main

import (
	_ "github.com/Shows/init"
	_ "github.com/Shows/models"
	_ "github.com/Shows/routers"
	"github.com/Shows/temp_func"
	"github.com/astaxie/beego"
)

func main() {

	beego.SetStaticPath("/episode/episode_files", "static/episode_file")
	beego.AddFuncMap("add", temp_func.Add)
	beego.AddFuncMap("randomImg", temp_func.RandomImage)
	beego.AddFuncMap("episodeReplace", temp_func.EpisodeReplace)
	beego.Run()

}
