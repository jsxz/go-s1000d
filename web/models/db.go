package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(192.168.56.11:3306)/go-s1000d?charset=utf8&parseTime=true")
	orm.Debug=true
}
func AllElement() []*Element {
	o := orm.NewOrm()
	var items []*Element
	num, err := o.QueryTable("element").All(&items)
	fmt.Printf("Returned Rows Num: %s, %s", num, err)
	return items
}