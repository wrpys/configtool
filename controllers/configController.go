package controllers

import (
	"github.com/astaxie/beego"
)

type ConfigController struct {
	beego.Controller
}

func (c *ConfigController) Get() {
	c.TplName = "index.tpl"
}
