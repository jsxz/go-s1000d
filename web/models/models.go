package models

import (
	"github.com/astaxie/beego/orm"
	_ `github.com/go-sql-driver/mysql`
)
type Element struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Version int `json:"version"`
}
func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Element))
}
