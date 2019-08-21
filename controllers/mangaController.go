package controllers

import (
	"github.com/Shows/models"
	"github.com/astaxie/beego"
)

type MangaController struct {
	beego.Controller
}

func (m *MangaController) Manga() {
	var list []models.Manga
	list = models.GetMangaList()
	m.Data["mangaCount"] = models.CountManga()
	m.Data["mangaList"] = list
	m.TplName = "manga.html"
}
