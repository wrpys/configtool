package controllers

import (
	"configtool/common"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.tpl"
}

/**
构造分页对象
*/
func BuildPageInfo(c *ConfigController) (common.PageInfo, error) {
	pageInfo := common.PageInfo{}
	err := c.ParseForm(&pageInfo)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		pageInfo.Limit = pageInfo.PageSize
		pageInfo.Offset = pageInfo.PageSize * (pageInfo.PageNumber - 1)
	}
	return pageInfo, err
}
