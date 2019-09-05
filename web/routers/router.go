package routers

import (
	"github.com/astaxie/beego"
	"github.com/jsxz/go-s1000d/web/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{}, "*:Index")
	beego.Router("/index", &controllers.HomeController{}, "*:Index")
	beego.Router("/elements", &controllers.HomeController{}, "*:AllElement")
	beego.Router("/xslt", &controllers.HomeController{}, "*:AllXslt")
}
