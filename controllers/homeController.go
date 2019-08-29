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

	var list []models.Manga
	var moviesList []models.Movie
	list = models.GetMangaList()
	randPage := rand.Intn(5)
	moviesList = models.QueryMovies(randPage)
	home.Data["mangaCount"] = models.CountManga()
	home.Data["mangaList"] = list
	home.Data["moviesList"] = moviesList
	home.Data["movieCount"], _ = models.CountMovies()
	home.TplName = "home.html"
}
