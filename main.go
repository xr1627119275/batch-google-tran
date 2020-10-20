package main

import (
	"encoding/json"
	"fmt"
	"github.com/gobuffalo/packr"
	"github.com/zserge/lorca"
	"log"
	"net"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	box := packr.NewBox("./html")
	//mux.Handle("/", http.FileServer(http.Dir("html")))
	mux.Handle("/", http.FileServer(box))
	//mux.HandleFunc("/ajax", func(writer http.ResponseWriter, request *http.Request) {
	//	writer.Write([]byte("bye bye ,this is v3 httpServer"))
	//})
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}

	defer ln.Close()

	go http.Serve(ln, mux)

	ui, _ := lorca.New("", "", 600, 900)
	defer ui.Close()

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

	ui.Load("data:text/html,<html lang=\"en\">" +
		"<head>" +
		" <meta charset=\"UTF-8\"> <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"> <title>加载中</title>" +
		"</head>" +
		"<body style='padding: 0;margin: 0;height: 100vh;display: flex; justify-content: center;align-items: center;'> <h1 >加载中</h1>" +
		"<script>   location.href = '" + fmt.Sprintf("http://%s", ln.Addr()) + "' </script>" +
		"</body" +
		"></html>")

	//ui.Load(fmt.Sprintf("http://%s", ln.Addr()))

	<-ui.Done()
}
