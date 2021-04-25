package main

import (
	"encoding/json"
	"fmt"
	"github.com/gobuffalo/packr"
	"github.com/zserge/lorca"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

var needUpdate = false

func getInfo() (size int64, data string, tagName string) {
	size = 0
	url := "https://api.github.com/repos/xr1627119275/batch-google-tran/releases"
	method := "GET"
	client := &http.Client{}
	req, _ := http.NewRequest(method, url, nil)

	res, _ := client.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	data = string(body)
	fmt.Println(data)
	r := regexp.MustCompile(`\"size\":(.+?),`)
	if r.MatchString(data) {

		matchAll := r.FindAllStringSubmatch(data, -1)
		fmt.Println(matchAll)

		var data = matchAll[0][1]
		size, _ = strconv.ParseInt(data, 10, 64)

		stat, _ := os.Stat(os.Args[0])

		for i := 0; i < len(matchAll); i++ {
			var size, _ = strconv.ParseInt(matchAll[i][1], 10, 64)
			if size == stat.Size() {
				var jsonData []map[string]string
				json.Unmarshal(body, &jsonData)
				fmt.Println(jsonData[i]["tag_name"])
				tagName = jsonData[i]["tag_name"]
			}
		}
	}
	return
}
func CheckUpdate(size int64) {
	if size <= 0 {
		return
	}
	stat, _ := os.Stat(os.Args[0])
	if size != stat.Size() {
		needUpdate = true
	}
}
func main() {
	var size, data, tagName = getInfo()
	CheckUpdate(size)

	mux := http.NewServeMux()
	box := packr.NewBox("./html")
	mux.Handle("/", http.FileServer(box))

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}

	defer ln.Close()

	go http.Serve(ln, mux)

	ui, _ := lorca.New("", "", 600, 900)
	defer ui.Close()
	ui.Bind("getSize", func() int64 {
		return size
	})
	ui.Bind("getTagName", func() string {
		return tagName
	})
	ui.Bind("getUpdateInfo", func() string {
		return data
	})
	ui.Bind("gotranslate", func(data map[string]interface{}) string {

		q := data["data"].(string)
		types := data["types"].([]interface{})

		langs := make([]string, len(types))
		resData := &TkRes{
			Success: false,
			Data:    map[string]interface{}{},
		}

		for i, item := range types {
			langs[i] = item.(string)
		}
		fmt.Printf("data :%v, types: %v ", q, langs)

		Tran(q, langs, resData)
		fmt.Println("success")
		resData.Success = true
		res, err := json.Marshal(resData)

		if err != nil {
			return "{'success': false}"
		}

		return string(res)
	})
	var defaultUrl = ""
	if needUpdate {
		defaultUrl = "/update.html"
	}
	ui.Load("data:text/html,<html lang=\"en\">" +
		"<head>" +
		" <meta charset=\"UTF-8\"> <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"> <title>加载中</title>" +
		"</head>" +
		"<body style='padding: 0;margin: 0;height: 100vh;display: flex; justify-content: center;align-items: center;'> <h1 >加载中</h1>" +
		"<script>   location.href = '" + fmt.Sprintf("http://%s%s", ln.Addr(), defaultUrl) + "' </script>" +
		"</body" +
		"></html>")
	<-ui.Done()
}
