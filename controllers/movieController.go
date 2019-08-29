package controllers

import (
	"github.com/Shows/models"
	"github.com/astaxie/beego"
)

type MovieController struct {
	beego.Controller
}

//动漫页面
func (m *MovieController) Movies() {
	pageIndex, _ := m.GetInt(":pageIndex")

	m.Data["PageIndex"] = pageIndex
	m.Data["moviesList"] = models.QueryMovies(pageIndex)
	m.Data["moviesCount"], m.Data["PageCount"] = models.CountMovies()
	m.TplName = "movies.html"
}

//剧集页面
func (m *MovieController) Episode() {
	movieId := m.GetString(":id")
	var video models.Movie
	video = models.GetMovie(movieId)
	m.Data["Video"] = video
	m.Data["List"] = models.GetEpisodeList(movieId)
	m.TplName = "episode.html"

}
func (m *MovieController) Player() {
	episodeId := m.GetString(":id")
	var movieId string
	m.Data["M3u8"], movieId = models.GetEpisodeUrl(episodeId)
	m.Data["List"] = models.GetEpisodeList(movieId)

	m.TplName = "player.html"

}
func (m *MovieController) GetMovieByName() {
	movieName := m.GetString("name")
	var movies []models.Movie
	movies = models.GetMovieName(movieName)
	m.Data["moviesList"] = movies
	m.Data["PageIndex"] = 1
	m.Data["moviesCount"] = 23
	m.Data["PageCount"] = 1
	m.TplName = "movies.html"
}
