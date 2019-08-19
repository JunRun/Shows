package controllers

import (
	"github.com/astaxie/beego"
	"myFirstProject/models"
)

type MovieController struct {
	beego.Controller
}

func (m *MovieController) Movies()  {
//   off,_:= m.GetInt("off")
//   endIndex,_:= m.GetInt("endIndex")
   m.Data["moviesList"] = models.QueryMovies()
   m.Data["moviesCount"] = models.CountMovies()
   m.TplName = "movies.html"
}