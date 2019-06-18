package main

import (
	_ "configtool/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func init() {
	logs.SetLogger("console")
	logs.SetLogger(logs.AdapterFile, `{"filename":"configtool.log","level":7}`)
	logs.EnableFuncCallDepth(true)
}

func main() {
	beego.Run()
}
