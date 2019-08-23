package routers

import (
	"github.com/Shows/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/home", &controllers.HomeController{}, "*:Home")
	beego.Router("/movies/page/:pageIndex", &controllers.MovieController{}, "*:Movies")
	beego.Router("/manga", &controllers.MangaController{}, "*:Manga")
	beego.Router("/episode/:id", &controllers.MovieController{}, "*:Episode")
}
