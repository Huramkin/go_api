package domain

import (
	"io/ioutil"
	"net/http"
)

func GetEmail() {
	resp, err := http.Get("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	defer resp.Body.Close()
	pageBytes, err := ioutil.ReadAll(resp.Body)
	pageStr := string(pageBytes)
	println(pageStr)
	println(err)
}
