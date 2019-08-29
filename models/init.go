package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

func init() {
	fmt.Print("dadas")
	orm.RegisterModel(new(Manga), new(Movie))
}
