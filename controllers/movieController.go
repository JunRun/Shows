package controllers

import (
	"github.com/Shows/models"
	"github.com/astaxie/beego"
)

type MovieController struct {
	beego.Controller
}

func (m *MovieController) Movies() {
	//   off,_:= m.GetInt("off")
	//   endIndex,_:= m.GetInt("endIndex")
	m.Data["moviesList"] = models.QueryMovies()
	m.Data["moviesCount"] = models.CountMovies()
	m.TplName = "movies.html"
}
func (m *MovieController) Episode() {

	movieId := m.GetString(":id")
	var video models.Movie
	video = models.GetMovie(movieId)
	m.Data["Video"] = video
	m.TplName = "episode.html"

}
