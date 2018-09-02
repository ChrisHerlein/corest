package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	tui "github.com/marcusolsson/tui-go"
)

var methodList = []string{"GET", "PUT", "POST", "DELETE"}
var methodToInt = map[string]int{"GET": 0, "PUT": 1, "POST": 2, "DELETE": 3}
var curReqWidgets = curReqMatch{}

type curReqMatch struct {
	Method  *tui.List
	Url     *tui.Entry
	Headers *tui.TextEdit
	Body    *tui.TextEdit
}

type request struct {
	Method  string            `json:"method"`
	Url     string            `json:"url"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
	Answer  string            `json:"-"`
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
		indentAns, _ := json.MarshalIndent(ansmap, "", "    ")
		curReq.Answer = string(indentAns)
		ansBody.SetText(curReq.Answer)

		var headerText string
		for k, v := range res.Header {
			headerText += fmt.Sprintf("%s: %+v\n", k, v)
		}
		ansHeaders.SetText(headerText)
		return
	}
	ansBody.SetText("Error ocurred: " + e.Error())
}

type showBody struct {
	Fields map[string]interface{}
	Maps   map[string]showBody
	Arrays map[string][]interface{}
	Wdg    tui.Widget
}

func generateAnsWidgets(body []byte) tui.Widget {
	if len(body) == 0 {
		return nil
	}
	var bdarr []interface{}
	e := json.Unmarshal(body, &bdarr)
	if e != nil || len(bdarr) == 0 {
		var bdmap map[string]interface{}
		e = json.Unmarshal(body, &bdmap)
		if e != nil || len(bdmap) == 0 {
			return nil
		}
		return generateFromMap(bdmap)
	}
	return generateFromArray(bdarr)
}

func generateFromMap(body map[string]interface{}) tui.Widget {
	var ans tui.Widget
	return ans
}

func generateFromArray(body []interface{}) tui.Widget {
	var ans tui.Widget
	return ans
}
