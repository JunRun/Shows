package routers

import (
	"github.com/Shows/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/home", &controllers.HomeController{}, "*:Home")
	beego.Router("/movies", &controllers.MovieController{}, "*:Movies")
	beego.Router("/movies/page/:pageIndex", &controllers.MovieController{}, "*:Movies")
	beego.Router("/manga", &controllers.MangaController{}, "*:Manga")
	beego.Router("/episode/:id", &controllers.MovieController{}, "*:Episode")
	beego.Router("/episode/player/:id", &controllers.MovieController{}, "*:Player")
	beego.Router("/home/movies/name", &controllers.MovieController{}, "*:GetMovieByName")
	beego.Router("/movies/tags/:tag", &controllers.MovieController{}, "*:GetMovieByTags")
	beego.Router("/movies/random", &controllers.MovieController{}, "*:GetMovieByRandom")
	beego.Router("/movies/lasted", &controllers.MovieController{}, "*:GetMovieLasted")
}
