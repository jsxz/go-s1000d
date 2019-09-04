package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/jsxz/go-s1000d/web/models"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) AllElement()  {
	element := models.AllElement()
	bytes, _ := json.Marshal(element)
	c.Data["json"]=string(bytes)
	c.ServeJSON()
}
func (c *HomeController) Index() {

	c.Data["json"]="home index"
	c.ServeJSON()
}
