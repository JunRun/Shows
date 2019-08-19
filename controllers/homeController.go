package controllers

import (
	"github.com/astaxie/beego"
	"myFirstProject/models"
)

type HomeController struct {
	beego.Controller
}

func (home *HomeController)Home()  {
	var list []models.Manga
	var moviesList []models.Movie
	list = models.GetMangaList()
	moviesList = models.QueryMovies()
	home.Data["mangaCount"] = models.CountManga()
	home.Data["mangaList"] = list
	home.Data["moviesList"] = moviesList
	home.Data["movieCount"] = models.CountMovies()
	home.TplName="home.html"
}