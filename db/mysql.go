package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/jsxz/go-s1000d/common"
	"github.com/jsxz/go-s1000d/models"
)

var DB *sqlx.DB

func init() {
	var err error
	//DB, err = sqlx.Open(`mysql`, `root:root@tcp(127.0.0.1:3306)/go-s1000d?charset=utf8&parseTime=true`)
	DB, err = sqlx.Open(`mysql`, `root:root@tcp(192.168.56.11:3306)/go-s1000d?charset=utf8&parseTime=true`)
	common.CheckError(err)
	if err = DB.Ping(); err != nil {
		panic(err)
	}
}
func ElementAdd(name string, version int) bool {
	obj, e := ElementOne(name)
	common.CheckError(e)
	if e != nil || name == obj.Name {
		fmt.Println(name + "已存在")
		return false
	}

	_, err := DB.Exec("insert into `element` (`name`,`version`) values(?,?)", name, version)
	common.CheckError(err)
	return true
}
func ElementAll() ([]models.Element, error) {
	mods := make([]models.Element, 0)
	err := DB.Select(&mods, "select * from `element` sort by id asc")
	return mods, err
}
func ElementOne(name string) (models.Element, error) {
	mods := make([]models.Element, 0)
	err := DB.Select(&mods, "select * from `element` where name=?", name)
	common.CheckError(err)
	if len(mods) > 0 {
		return mods[0], err
	} else {
		return models.Element{}, err
	}
}
