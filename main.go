package main

import (
	"fmt"
	"github.com/jsxz/go-s1000d/common"
	"github.com/jsxz/go-s1000d/db"
)

func main() {
	folder := "xsd/4.2"
	files := common.ListFile(folder)
	for _, f := range files {
		//if strings.HasSuffix(f,"proced.xsd"){
		fmt.Println(f)
		s := common.ReadToString(f)
		items := common.GetData(s, `<xs:element\sname="(\w+)"`)
		for _, v := range items {
			db.ElementAdd(v, 4)
		}
		//}
	}
}
