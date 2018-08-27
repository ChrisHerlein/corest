package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	tui "github.com/marcusolsson/tui-go"
)

var methodList = []string{"GET", "PUT", "POST", "DELETE"}

type request struct {
	Method  string
	Url     string
	Headers map[string]string
	Body    string
	Answer  string
}

func sendButtonFunc(b *tui.Button) {
	req, _ := http.NewRequest(curReq.Method, curReq.Url, bytes.NewBufferString(curReq.Body))
	for k, v := range curReq.Headers {
		req.Header.Add(k, v)
	}
	res, e := http.DefaultClient.Do(req)
	if e == nil {
		bd, _ := ioutil.ReadAll(res.Body)
		var ansmap map[string]interface{}
		json.Unmarshal(bd, &ansmap)
		indentAns, _ := json.MarshalIndent(ansmap, "", "\t")
		curReq.Answer = string(indentAns)
		ansWidget.SetText(curReq.Answer)
		return
	}
	ansWidget.SetText("Some error ocurred")
}
