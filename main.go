package main

import (
	_ "github.com/asiainfoLDP/datahub_custom/docs"
	_ "github.com/asiainfoLDP/datahub_custom/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {
	//if beego.BConfig.RunMode == "dev" {
	//	beego.BConfig.WebConfig.DirectoryIndex = true
	//	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	//}
	orm.Debug = true
	beego.Run()
}
