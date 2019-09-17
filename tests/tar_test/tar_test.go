package tar_test

import (
	"encoding/json"
	"fmt"
	"github.com/Shows/conf"
	_ "github.com/Shows/init"
	"github.com/Shows/models"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/orm"
	"github.com/dollarkillerx/easyutils"
	"github.com/dollarkillerx/easyutils/clog"
	"strconv"
	"sync"
	"testing"
)

type Res struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

func TestTar(t *testing.T) {
	configMap := conf.InitConfig("/Users/jun/go/src/github.com/Shows/tests/tar_test/app.conf")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.Debug = true
	//fmt.Println(beego.AppConfig.String("mysql_username"))
	username := configMap["mysql_username"]
	password := configMap["mysql_password"]
	database := configMap["mysql_database"]
	host, _ := strconv.Atoi(configMap["mysql_host"])
	port, _ := configMap["mysql_port"]

	_ = orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%d:%s)/%s?charset=utf8",
		username, password, host, port, database))
	o := orm.NewOrm()
	var episodeList []models.Episode
	sy1 := sync.WaitGroup{}
	_, _ = o.Raw("select * from episode").QueryRows(&episodeList)
	pool := make(chan int, 10)
	for _, episode := range episodeList {
		sy1.Add(1)
		id := episode.Id
		info := episode.EpisodeInfo
		pool <- 1
		go func(id int, info string) {
			defer func() {
				<-pool
				sy1.Done()
			}()
			//西班牙语  澳大利亚语
			tarUrl := fmt.Sprintf("https://translate.dollarkiller.com/translate?sl=de&tl=es&text=%s", info)
			url, er := easyutils.UrlEncoding(tarUrl)
			if er != nil {
				panic(episode.EpisodeInfo + " -----错误")
			}
		ki:
			bytes := httplib.Get(url)
			s, e := bytes.String()

			if e != nil {
				clog.Println("请求错误 进行尝试")
				clog.Println(e.Error())
				goto ki
			}
			var res Res
			if err := json.Unmarshal([]byte(s), &res); err == nil {

			}

			_, e = o.Raw("update episode set de_episode_info =? where id = ?", res.Msg, id).Exec()

		}(id, info)
	}
	sy1.Wait()
}

func TestEn(t *testing.T) {
	configMap := conf.InitConfig("/Users/jun/go/src/github.com/Shows/tests/tar_test/app.conf")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.Debug = true
	//fmt.Println(beego.AppConfig.String("mysql_username"))
	username := configMap["mysql_username"]
	password := configMap["mysql_password"]
	database := configMap["mysql_database"]
	host, _ := strconv.Atoi(configMap["mysql_host"])
	port, _ := configMap["mysql_port"]

	_ = orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%d:%s)/%s?charset=utf8",
		username, password, host, port, database))
	o := orm.NewOrm()
	var movieList []models.Movie
	sy1 := sync.WaitGroup{}
	_, _ = o.Raw("select * from movie").QueryRows(&movieList)
	pool := make(chan int, 50)
	for _, movie := range movieList {
		sy1.Add(1)
		id := movie.Id
		info := movie.Info
		tags := movie.MovieType
		pool <- 1
		go func(id string, info string, name string) {
			defer func() {
				<-pool
				sy1.Done()
			}()
			nameUrl := fmt.Sprintf("https://translate.dollarkiller.com/translate?sl=&tl=de&text=%s", tags)
			//tarUrl := fmt.Sprintf("https://translate.dollarkiller.com/translate?sl=&tl=de&text=%s", info)
			//url, er := easyutils.UrlEncoding(tarUrl)
			NUrl, err := easyutils.UrlEncoding(nameUrl)
			//if er != nil {
			//	panic(movie.Info + " -----描述错误")
			//}
			if err != nil {
				panic(name + " -----------------------名字信息错误")
			}
		ki:
			//bytes := httplib.Get(url)
			nameBytes := httplib.Get(NUrl)
			//s, e := bytes.String()
			l, e := nameBytes.String()

			if e != nil {
				clog.Println("请求错误 进行尝试")
				clog.Println(e.Error())
				goto ki
			}

			//var res Res
			var nameRes Res
			//if err := json.Unmarshal([]byte(s), &res); err == nil {
			//
			//}
			if err := json.Unmarshal([]byte(l), &nameRes); err == nil {
			}

			_, e = o.Raw("update movie set movie_type = ? where id = ?", nameRes.Msg, id).Exec()

		}(id, info, movie.MovieName)
	}
	sy1.Wait()
}
func TestUrlEnco(t *testing.T) {
	tarUrl := fmt.Sprintf("https://translate.dollarkiller.com/translate?sl=en&tl=de&text=%s", "Murasaki Wakako is a twenty-seven year old working lady. As per her daily routine, she rewards herself for a hard day at                         work by bar-hopping in search of the best alcohol and seafood the world has to offer. On tonights menu, she has some crispy fried salted salmon and cold sake. This ultimate combination makes her purr in satisfaction.")
	url, err := easyutils.UrlEncoding(tarUrl)
	clog.Println(url)
	if err != nil {
		panic(err.Error())
	}

	get := httplib.Get(url)
	s, err := get.String()
	if err == nil {
		t.Log(s)
	}
}

func TestString(t *testing.T) {
	name := "Episode20"
	l := name[7:]
	fmt.Println(l)

}
