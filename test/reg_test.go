package test

import (
	"log"
	"regexp"
	"testing"
)

func TestReg(t *testing.T) {
	data := "[[[\"This is a translation\",\"这是一条翻译\",null,null,0]\n]\n,null,\"zh-CN\",null,null,null,null,[]\n]\n"
	r := regexp.MustCompile(`\"(.+?)\"`)
	if r.MatchString(data) {
		log.Println(r.FindStringSubmatch(data)[1])
	}
}
