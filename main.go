package main

import (
	"encoding/json"
	"fmt"
	"github.com/zserge/lorca"
	"log"
	"net"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("html")))
	mux.HandleFunc("/ajax", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("bye bye ,this is v3 httpServer"))
	})
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}

	defer ln.Close()

	go http.Serve(ln,  mux)


	ui, _ := lorca.New("", "", 600, 900)
	defer ui.Close()



	ui.Bind("gotranslate", func(data map[string] interface{})  string {

		q := data["data"].(string)
		types := data["types"].([]interface{})

		langs := make([]string, len(types))
		resData := &TkRes{
			Success: false,
			Data: map[string]interface{}{},
		}

		for i, item := range types {
			langs[i] = item.(string)
		}
		fmt.Printf("data :%v, types: %v \n", q, langs)

		Tran(q,langs, resData)
		fmt.Println("success")
		resData.Success = true
		res , err := json.Marshal(resData)

		if err != nil {
			return "{'success': false}"
		}

		return string(res)

	})





	ui.Load(fmt.Sprintf("http://%s", ln.Addr()))


	<-ui.Done()
}

