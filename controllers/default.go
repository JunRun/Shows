package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "rrrrrrrr"
	c.Data["Email"] = "dasdasdadadadasdsada"
	c.TplName = "index.tpl"
}
