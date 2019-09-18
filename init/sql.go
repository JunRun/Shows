package init

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.Debug = true
	fmt.Println(beego.AppConfig.String("mysql_username"))
	username := beego.AppConfig.String("mysql_username")

	password := beego.AppConfig.String("mysql_password")
	database := beego.AppConfig.String("mysql_database")
	showsDatabase := beego.AppConfig.String("mysql_shows")
	host := beego.AppConfig.String("mysql_host")
	port, _ := beego.AppConfig.Int("mysql_port")

	orm.RegisterDriver("mysql", orm.DRMySQL)

	_ = orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		username, password, host, port, database))
	_ = orm.RegisterDataBase("shows", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		username, password, host, port, showsDatabase))
	//o.Using("txin")
	//fmt.Println("ok")
}
