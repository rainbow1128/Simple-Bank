package utils

import (
	"flag"
	"fmt"

	"gitlab.com/Simple-Bank/types"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var OrmInstance orm.Ormer

func InitDB() {
	dbURI := beego.AppConfig.String("dbURI")
	idledbconnection, _ := beego.AppConfig.Int("idledbconnection")
	var host string
	flag.StringVar(&host, "host", "0.0.0.0", "host of mysql")
	flag.Parse()
	dbURI = fmt.Sprintf(dbURI, host)
	maxdbconnection, _ := beego.AppConfig.Int("maxdbconnection")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dbURI, idledbconnection, maxdbconnection)
	orm.RegisterModel(
		new(types.Customer),
		new(types.Account),
		new(types.Transaction),
	)
	orm.RunSyncdb("default", false, true)
	OrmInstance = orm.NewOrm()
}
