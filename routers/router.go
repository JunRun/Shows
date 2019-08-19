package routers

import (
	"myFirstProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/home",&controllers.HomeController{},"*:Home")
    beego.Router("/movies",&controllers.MovieController{},"*:Movies")
    beego.Router("/manga",&controllers.MangaController{},"*:Manga")
}
