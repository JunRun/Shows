package models

import (
	"github.com/astaxie/beego/orm"
	"math"
)

type Movie struct {
	Id        string `orm:"pk"`
	MovieName string
	ImgUrl    string
	MovieUrl  string
	Publisher string
	MovieType string
	MovieMark string
	MovieYear string
	Info      string
	DeName    string
	DeInfo    string
}

var pageRange int = 15

func QueryMovies(pageIndex int) []Movie {
	o := orm.NewOrm()
	var list []Movie

	_, _ = o.Raw("select * from movie order by movie_mark desc limit ?,?", (pageIndex-1)*pageRange, pageRange).QueryRows(&list)
	return list
}

func CountMovies() (int, int) {
	o := orm.NewOrm()
	var count int
	_ = o.Raw("select count(*) from movie").QueryRow(&count)
	pageNum := math.Ceil(float64(count) / float64(pageRange))
	return count, int(pageNum)
}

func GetMovie(id string) Movie {
	var movie Movie
	o := orm.NewOrm()
	_ = o.Raw("select * from movie where id = ?", id).QueryRow(&movie)
	return movie
}

func GetMovieName(name string) []Movie {
	var movie []Movie
	o := orm.NewOrm()
	_, _ = o.Raw("select * from movie where movie_name  like ?", "%"+name+"%").QueryRows(&movie)
	return movie
}

func GetMovieByType(tags string) []Movie {
	var movie []Movie
	o := orm.NewOrm()
	_, _ = o.Raw("select * from movie where movie_type like ?", "%"+tags+"%").QueryRows(&movie)
	return movie
}
