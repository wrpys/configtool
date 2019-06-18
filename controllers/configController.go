package controllers

import (
	"configtool/daos"
	"configtool/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type ConfigController struct {
	beego.Controller
}

type Options struct {
	Profiles     []string
	Applications []string
}

const TAG = "ConfigController"

func (c *ConfigController) GetSearchOptions() {
	logs.Info(TAG, "GetSearchOptions")
	profiles := daos.GetDistinctProfile()
	applications := daos.GetDistinctApplication()
	c.Data["json"] = Options{profiles, applications}
	logs.Info(TAG, "GetSearchOptions", c.Data["json"])
	c.ServeJSON()
}

func (c *ConfigController) List() {
	configProperties := models.ConfigProperties{}
	if err := c.ParseForm(&configProperties); err != nil {
		c.Data["json"] = fail(err.Error())
	} else {
		total, rows := daos.FindConfigProperties(BuildPageInfo(c), configProperties)
		c.Data["json"] = map[string]interface{}{"total": total, "rows": rows}
	}
	c.ServeJSON()
}

func (c *ConfigController) Add() {
	configProperties := new(models.ConfigProperties)
	if err := c.ParseForm(configProperties); err != nil {
		c.Data["json"] = fail(err.Error())
	} else {
		configProperties.Label = "master"
		daos.CreateConfigProperties(configProperties)
		c.Data["json"] = success()
	}
	c.ServeJSON()
}

func (c *ConfigController) Update() {
	configId, err := c.GetInt("configId")
	if err != nil {
		c.Data["json"] = fail(err.Error())
	}
	configProperties := new(models.ConfigProperties)
	if err := c.ParseForm(configProperties); err != nil {
		c.Data["json"] = fail(err.Error())
	} else {
		configProperties.ConfigId = configId
		daos.UpdateConfigProperties(configProperties)
		c.Data["json"] = success()
	}
	c.ServeJSON()
}

func (c *ConfigController) Delete() {
	configIds := make([]int, 0)
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &configIds); err != nil {
		c.Data["json"] = fail(err.Error())
	} else {
		daos.DeleteConfigProperties(configIds)
		c.Data["json"] = success()
	}
	c.ServeJSON()
}

func (c *ConfigController) BatchUpdateProfile() {
	currentProfile := c.Ctx.Input.Param(":currentProfile")
	deployProfile := c.Ctx.Input.Param(":deployProfile")
	if currentProfile == deployProfile {
		c.Data["json"] = fail("环境名相同")
	} else {
		daos.BatchUpdateProfile(currentProfile, deployProfile)
		c.Data["json"] = success()
	}
	c.ServeJSON()
}

func success() map[string]interface{} {
	return map[string]interface{}{"flag": 1}
}

func fail(msg string) map[string]interface{} {
	return map[string]interface{}{"flag": 0, "msg": msg}
}
