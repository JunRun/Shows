package controllers

import (
	"github.com/Shows/models"
	"github.com/astaxie/beego"
	"math/rand"
)

type HomeController struct {
	beego.Controller
}

func (home *HomeController) Home() {

	var moviesList []models.Movie
	randPage := rand.Intn(5)
	moviesList = models.QueryMovies(randPage)
	home.Data["moviesList"] = moviesList
	home.Data["movieCount"], _ = models.CountMovies()
	home.Data["headerMovie"] = models.GetMovieByRandom(20)
	home.TplName = "home.html"
}
