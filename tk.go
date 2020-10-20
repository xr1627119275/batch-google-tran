package main

import (
	"fmt"
	gt "github.com/kyai/google-translate-tk"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"sync"
)
var LangsMap = map[string] string {"auto":"Automatic","af":"Afrikaans","sq":"Albanian","ar":"Arabic","hy":"Armenian","az":"Azerbaijani","eu":"Basque","be":"Belarusian","bn":"Bengali","bs":"Bosnian","bg":"Bulgarian","ca":"Catalan","ceb":"Cebuano","ny":"Chichewa","zh-cn":"Chinese Simplified","zh-tw":"Chinese Traditional","co":"Corsican","hr":"Croatian","cs":"Czech","da":"Danish","nl":"Dutch","en":"English","eo":"Esperanto","et":"Estonian","tl":"Filipino","fi":"Finnish","fr":"French","fy":"Frisian","gl":"Galician","ka":"Georgian","de":"German","el":"Greek","gu":"Gujarati","ht":"Haitian Creole","ha":"Hausa","haw":"Hawaiian","iw":"Hebrew","hi":"Hindi","hmn":"Hmong","hu":"Hungarian","is":"Icelandic","ig":"Igbo","id":"Indonesian","ga":"Irish","it":"Italian","ja":"Japanese","jw":"Javanese","kn":"Kannada","kk":"Kazakh","km":"Khmer","ko":"Korean","ku":"Kurdish (Kurmanji)","ky":"Kyrgyz","lo":"Lao","la":"Latin","lv":"Latvian","lt":"Lithuanian","lb":"Luxembourgish","mk":"Macedonian","mg":"Malagasy","ms":"Malay","ml":"Malayalam","mt":"Maltese","mi":"Maori","mr":"Marathi","mn":"Mongolian","my":"Myanmar (Burmese)","ne":"Nepali","no":"Norwegian","ps":"Pashto","fa":"Persian","pl":"Polish","pt":"Portuguese","ma":"Punjabi","ro":"Romanian","ru":"Russian","sm":"Samoan","gd":"Scots Gaelic","sr":"Serbian","st":"Sesotho","sn":"Shona","sd":"Sindhi","si":"Sinhala","sk":"Slovak","sl":"Slovenian","so":"Somali","es":"Spanish","su":"Sudanese","sw":"Swahili","sv":"Swedish","tg":"Tajik","ta":"Tamil","te":"Telugu","th":"Thai","tr":"Turkish","uk":"Ukrainian","ur":"Urdu","uz":"Uzbek","vi":"Vietnamese","cy":"Welsh","xh":"Xhosa","yi":"Yiddish","yo":"Yoruba","zu":"Zulu"}

var tkk string


type TkRes struct {
	Success bool `json:"success"`
	Data map[string] interface {} `json:"data"`
}

func tran(q string, tk string, to string, res *TkRes, wg *sync.WaitGroup) {

	defer func() {
		err := recover()
		if err != nil {
			log.Printf("err : %v \n", err)
		}
	}()

	var  url = "https://translate.google.cn/translate_a/single?client=t&sl=zh-cn&tl="+to+"&hl="+to+"&dt=t&ie=UTF-8&oe=UTF-8&otf=1&ssel=0&tsel=0&kc=7&q="+ url.QueryEscape(q) +"&tk=" + tk

	resp, err := http.Get(url)
	resp.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.122 Safari/537.36")
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	data := string(body)

	fmt.Println(data)

	r := regexp.MustCompile(`\"(.+?)\"`)
	if r.MatchString(data) {
		data = r.FindStringSubmatch(data)[1]
		res.Data[to] = data
	}

	wg.Done()
}

func Tran(q string, langs []string, res *TkRes) {

	if len(langs) == 0 {
		langs = []string{ "Automatic" }
	}
	tk := gt.GetTK(q, tkk)
	var wg = &sync.WaitGroup{}
	for _, item := range langs {
		for key, val := range LangsMap {
			if item == val {
				item = key
			}
		}
		wg.Add(1)
		go tran(q,tk ,item, res, wg)
	}
	wg.Wait()



}


func init() {
	tkk ,_ = gt.GetTKK()
}