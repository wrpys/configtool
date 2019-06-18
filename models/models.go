package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type ConfigProperties struct {
	ConfigId    int    `orm:"pk";form:"configId"`
	ConfigKey   string `form:"configKey"`
	ConfigValue string `form:"configValue"`
	Application string `form:"application"`
	Profile     string `form:"profile"`
	Label       string
	CreateDate  time.Time `orm:"column(CREATE_DATE);auto_now_add;type(datetime)"`
	ChangeDate  time.Time `orm:"column(CHANGE_DATE);auto_now;type(datetime)"`
	ConfigDesc  string    `form:"configDesc"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(ConfigProperties))
}
