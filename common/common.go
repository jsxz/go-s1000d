package common

import (
	"io/ioutil"
	"regexp"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
func ReadToString(f string) string {
	bytes, e := ioutil.ReadFile(f)
	CheckError(e)
	return string((bytes))
}

func ListFile(folder string) []string {
	var res []string
	files, _ := ioutil.ReadDir(folder) //specify the current dir
	for _, file := range files {
		if file.IsDir() {
			ListFile(folder + "/" + file.Name())
		} else {
			//fmt.Println(folder + "/" + file.Name())
			res = append(res, folder+"/"+file.Name())
		}
	}
	return res
}
func GetData(s, pattern string) []string {
	reg := regexp.MustCompile(pattern)
	submatch := reg.FindAllSubmatch([]byte(s), -1)
	var items []string
	for _, v := range submatch {
		for i, v1 := range v {
			s1 := string(v1)
			if i == 1 {
				items = append(items, s1)
			}
		}
	}
	return items
}
