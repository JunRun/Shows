package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(new(Manga),new(Movie))
}
