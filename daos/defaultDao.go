package daos

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var dao orm.Ormer

func init() {
	orm.RegisterDriver(beego.AppConfig.String("dataSource.driverName"), orm.DRMySQL)
	orm.RegisterDataBase(beego.AppConfig.String("dataSource.aliasName"),
		beego.AppConfig.String("dataSource.driverName"),
		beego.AppConfig.String("dataSource.dataSource"),
		beego.AppConfig.DefaultInt("dataSource.maxIdle", 5),
		beego.AppConfig.DefaultInt("dataSource.maxConn", 10))

	orm.Debug = beego.AppConfig.DefaultBool("orm.Debug", true)

	dao = orm.NewOrm()
	dao.Using(beego.AppConfig.String("dataSource.ormName"))

}
