package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Orm orm.Ormer
)

func Init() {

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/fileserver?charset=utf8")
	orm.RunSyncdb("default", false, true)
	Orm = orm.NewOrm()
}
