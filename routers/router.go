package routers

import (
	"configtool/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/config/getSearchOptions", &controllers.ConfigController{}, "get:GetSearchOptions")
	beego.Router("/config/list", &controllers.ConfigController{}, "get:List")
	beego.Router("/config/add", &controllers.ConfigController{}, "post:Add")
	beego.Router("/config/update", &controllers.ConfigController{}, "post:Update")
	beego.Router("/config/delete", &controllers.ConfigController{}, "post:Delete")
	beego.Router("/config/batchUpdateProfile/:currentProfile/:deployProfile", &controllers.ConfigController{}, "post:BatchUpdateProfile")

	beego.InsertFilter("/*", beego.BeforeRouter, FilterRequestParamsPrint)
	beego.InsertFilter("/*", beego.FinishRouter, FilterResponseBodyPrint, false)

}

var FilterRequestParamsPrint = func(ctx *context.Context) {
	logs.Info("request uri:", ctx.Request.RequestURI)
	logs.Info("request Form:", ctx.Request.Form)
}

var FilterResponseBodyPrint = func(ctx *context.Context) {
	logs.Info("request uri:", ctx.Request.RequestURI, ",ok:", ctx.Output.IsOk())
}
