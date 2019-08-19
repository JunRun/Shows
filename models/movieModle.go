package models

import "github.com/astaxie/beego/orm"

type Movie struct {
	Id int `orm:"pk"`
	MovieName string
	ImgUrl string
	MovieUrl string
	Publisher string
	MovieType string
	MovieMark string
	MovieYear string
	Info string
}

func QueryMovies() []Movie {
	o := orm.NewOrm()
	var list []Movie
	_,_=o.Raw("select * from movie order by movie_mark desc limit ?,?",1,20).QueryRows(&list)
	return list
}

func CountMovies() int{
	o:= orm.NewOrm()
	var count int
	_ = o.Raw("select count(*) from movie").QueryRow(&count)
	return count
}