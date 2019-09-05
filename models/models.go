package models

type Element struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Version int    `json:"version"`
}
type Xslt struct {
	Id    int    `json:"id"`
	Match string `json:"match"`
	File  string `json:"file"`
}
