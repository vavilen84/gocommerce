package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/vavilen84/gocommerce/routers"
	"github.com/vavilen84/gocommerce/store"
)

func init() {
	store.InitORM()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
