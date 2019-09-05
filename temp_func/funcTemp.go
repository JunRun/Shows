package temp_func

import "github.com/astaxie/beego/orm"

func Add(args1 int, args2 int) int {
	return args1 + args2
}

func RandomImage() string {
	o := orm.NewOrm()
	var imageUrl string
	o.Raw("SELECT img_url FROM movie ORDER BY RAND() LIMIT 1").QueryRow(&imageUrl)
	return imageUrl

}
