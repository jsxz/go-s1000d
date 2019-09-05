package main

import (
	"fmt"
	"github.com/jsxz/go-s1000d/common"
	"github.com/jsxz/go-s1000d/db"
	"strings"
)

func main() {
	folder := `D:\work\git\xml-translator\4.2`
	files := common.ListFile(folder)
	for _, f := range files {
		//if strings.HasSuffix(f, "procedure.xslt") {
		if strings.HasSuffix(f, ".xslt") {
			fmt.Println(f)
			s := common.ReadToString(f)
			items := common.GetData(s, `<xsl:template\s*match="(.*)?"`)
			fmt.Println(items)
			for _, v := range items {
				split := strings.Split(f, "/")
				if len(split) > 1 {
					file := split[1]
					db.XsltAdd(v, file)
				} else {
					db.XsltAdd(v, "")
				}
			}
		}
	}
}
