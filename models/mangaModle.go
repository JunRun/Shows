package models

import 	"github.com/astaxie/beego/orm"

type  Manga struct {
	Id string `json:"id" orm:"pk"`
	Img string `json:"img"`
	Url string
	Name string
	Alt string


}
func (manga *Manga) TableName() string{
	return "manga"
}
func GetMangaList() []Manga{
	o := orm.NewOrm()
	var list []Manga
	_,_=o.Raw("select * from manga").QueryRows(&list)
	return list
}
func CountManga() int {
	o:=orm.NewOrm()
	var count int
	_= o.Raw("select count(*) from manga").QueryRow(&count)
	return count
}