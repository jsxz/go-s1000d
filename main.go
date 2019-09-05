package main

import (
	"fmt"
	"github.com/jsxz/go-s1000d/common"
	"github.com/jsxz/go-s1000d/db"
	"io/ioutil"
	"regexp"
)

func main() {
	folder := "xsd/4.2"
	files := listFile(folder)
	for _, f := range files {
		//if strings.HasSuffix(f,"proced.xsd"){
		fmt.Println(f)
		s := readToString(f)
		run(s, `<xs:element\sname="(\w+)"`)
		//}
	}
}
func run(s, pattern string) {
	reg := regexp.MustCompile(pattern)
	submatch := reg.FindAllSubmatch([]byte(s), -1)
	var items []string
	for _, v := range submatch {
		for i, v1 := range v {
			s1 := string(v1)
			if i == 1 {
				fmt.Println(s1)
				items = append(items, s1)
				db.ElementAdd(s1, 4)
			}
		}
	}
	fmt.Println(items)
	fmt.Println(len(items))
}
func readToString(f string) string {
	bytes, e := ioutil.ReadFile(f)
	common.CheckError(e)
	return string((bytes))
}

func listFile(folder string) []string {
	var res []string
	files, _ := ioutil.ReadDir(folder) //specify the current dir
	for _, file := range files {
		if file.IsDir() {
			listFile(folder + "/" + file.Name())
		} else {
			//fmt.Println(folder + "/" + file.Name())
			res = append(res, folder+"/"+file.Name())
		}
	}
	return res
}
