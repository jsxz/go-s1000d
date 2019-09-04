package main

import (
	"fmt"
	"regexp"
)
//https://github.com/google/re2/wiki/Syntax	 不支持
func main() {
	reg := regexp.MustCompile(`(?=l)`)
	allString := reg.FindAllString("hello",-1)
	fmt.Println(allString)
}
