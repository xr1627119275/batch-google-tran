package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"testing"
)

func TestStat(t *testing.T) {
	fmt.Println(os.Args[0])
	stat, err := os.Stat(os.Args[0])
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(stat.Size())
}

func TestGetVersion(t *testing.T) {
	url := "https://api.github.com/repos/xr1627119275/batch-google-tran/releases"
	method := "GET"
	client := &http.Client{}
	req, _ := http.NewRequest(method, url, nil)
	res, _ := client.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	data := string(body)
	fmt.Println(data)
	r := regexp.MustCompile(`\"size\":(.+?),`)
	if r.MatchString(data) {
		matchAll := r.FindAllStringSubmatch(data, -1)
		fmt.Println(matchAll)

		//var data = matchAll[0][1]
		//var size, _ = strconv.ParseInt(data, 10, 64)

		for i := 0; i < len(matchAll); i++ {
			var _size, _ = strconv.ParseInt(matchAll[i][1], 10, 64)
			fmt.Println(_size)
			if _size == msize {
				var jsonData []map[string]string
				json.Unmarshal(body, &jsonData)
				fmt.Println(jsonData[i]["tag_name"])
				return
			}
		}

	}
}

const msize = 6427136
