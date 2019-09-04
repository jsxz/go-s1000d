package main

import (
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/jsxz/go-s1000d/web/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.InsertFilter("*", beego.BeforeRouter,cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		//AllowOrigins: []string{"https://*.foo.com"},
		AllowMethods: []string{"GET","PUT", "PATCH","POST"},
		AllowHeaders: []string{"Origin"},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials: true,
	}))
	beego.Run()
}

