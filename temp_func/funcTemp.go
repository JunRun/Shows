package temp_func

import (
	"github.com/astaxie/beego/orm"
	"github.com/es_anime/models"
	"log"
	"sync"
)

func Add(args1 int, args2 int) int {
	return args1 + args2
}

func RandomImage() string {
	o := orm.NewOrm()
	var imageUrl string
	o.Raw("SELECT img_url FROM movie ORDER BY RAND() LIMIT 1").QueryRow(&imageUrl)
	return imageUrl

}

func EpisodeReplace(name string) string {
	l := name[7:]
	return l
}
func UpdateM3u8() error {
	o := orm.NewOrm()
	o.Using("shows")
	var m3u8 []models.Episode
	sql := "select id,episode_url from episode"
	_, err := o.Raw(sql).QueryRows(m3u8)
	if err != nil {
		log.Fatalln("数据库查询失败")
		return err
	}
	err = o.Using("default")
	if err != nil {
		log.Fatalln("数据库切换失败")
		return err
	}
	sy1 := sync.WaitGroup{}
	pool := make(chan int, 50)
	updateSql := "update episode set episode_url =? where id=?"
	for _, m3u8url := range m3u8 {
		id := m3u8url.Id
		m3u8Path := m3u8url.EpisodeUrl
		sy1.Add(1)
		pool <- 1
		go func(id int, n string) {
			defer func() {
				<-pool
				sy1.Done()
			}()

		}(id, m3u8Path)
		_, e := o.Raw(updateSql, m3u8Path, id).Exec()
		if e != nil {
			log.Fatalf("id := %d ,更新失败", id)
		}
	}
	return nil
}
