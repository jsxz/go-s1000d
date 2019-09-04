package routers

import (
	"github.com/jsxz/go-s1000d/web/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.HomeController{},"*:Index")
    beego.Router("/index", &controllers.HomeController{},"*:Index")
    beego.Router("/elements", &controllers.HomeController{},"*:AllElement")
}
