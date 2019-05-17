package controllers

import (
	"github.com/astaxie/beego"
)

type ControllerController struct {
	beego.Controller
}

func (c *ControllerController) Get() {
	c.TplName = "index.tpl"
}
