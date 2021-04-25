package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestAPI(t *testing.T) {
	url := "https://translate.googleapis.com/translate_a/single?client=gtx&sl=auto&tl=ko&dj=1&ie=UTF-8&oe=UTF-8&source=icon&dt=t&dt=bd&dt=qc&dt=rm&dt=ex&dt=at&dt=ss&dt=rw&dt=ld&q=%25E8%25BF%2599%25E6%2598%25AF%25E4%25B8%2580%25E6%259D%25A1%25E7%25BF%25BB%25E8%25AF%2591"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
