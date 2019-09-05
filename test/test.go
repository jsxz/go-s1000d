package main

import (
	"fmt"
	"strings"
)

//https://github.com/google/re2/wiki/Syntax	 不支持
func main() {
	//reg := regexp.MustCompile(`(.*)`)
	//allString := reg.FindAllString("hello/fsfs",-1)
	//fmt.Println(allString)
	test()
}
func test() {
	s := `d:\f/a.fx`
	split := strings.Split(s, "/")
	//after := strings.SplitAfter(s, "/")
	fmt.Println(strings.Split(s, "/"))
	print(len(split))
	fmt.Println(split[1])
}
