package main

import (
	_ "github.com/Shows/init"
	_ "github.com/Shows/models"
	_ "github.com/Shows/routers"
	"github.com/Shows/temp_func"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	temp_func2 "github.com/es_anime/temp_func"
)

func main() {

	tk := toolbox.NewTask("updateM3u8", "0 19 */1 * * ", temp_func2.UpdateM3u8)
	toolbox.AddTask("updateM3u8", tk)
	beego.SetStaticPath("/episode/episode_files", "static/episode_file")
	beego.AddFuncMap("add", temp_func.Add)
	beego.AddFuncMap("randomImg", temp_func.RandomImage)
	beego.AddFuncMap("episodeReplace", temp_func.EpisodeReplace)
	beego.Run()

}
