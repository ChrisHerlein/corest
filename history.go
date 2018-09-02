package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	tui "github.com/marcusolsson/tui-go"
)

const directory = "/.herrest"

var useHistory bool
var home string
var history []request

func ensurePath(hm string) {
	home = hm
	fi, e := os.Stat(home + directory)
	if e != nil {
		e = os.Mkdir(home+directory, 0755)
		if e != nil {
			fmt.Println("Cannot create herrest dir to save history")
			return
		}
	} else if !fi.IsDir() {
		fmt.Println("Cannot use herrest dir to save history")
		return
	}
	useHistory = true
}

func loadHistory() {
	if !useHistory {
		return
	}
	fl, e := ioutil.ReadFile(home + directory + "/history.json")
	if e != nil {
		fmt.Println("Cannot read history")
		return
	}

	e = json.Unmarshal(fl, &history)
	if e != nil {
		fmt.Println("Cannot understand history")
	}
}

var hstLst *tui.List

func saveRequest(btn *tui.Button) {
	var found bool
	for i := 0; i < len(history); i++ {
		if history[i].Url == curReq.Url && history[i].Method == curReq.Method {
			history[i].Body = curReq.Body
			history[i].Headers = curReq.Headers
			found = true
		}
	}

	if !found {
		history = append(history, curReq)
		hstLst.AddItems(fmt.Sprintf("%s %s", curReq.Method, curReq.Url))
	}

	hstJson, _ := json.Marshal(history)
	ioutil.WriteFile(home+directory+"/history.json", hstJson, 0644)
}

func deleteRequest(btn *tui.Button) {
	sel := hstLst.SelectedItem()

	var newItems = make([]string, 0)
	var nhst = make([]request, 0)
	for i := 0; i < len(history); i++ {
		if fmt.Sprintf("%s %s", history[i].Method, history[i].Url) == sel {
			continue
		}
		newItems = append(newItems, fmt.Sprintf("%s %s", history[i].Method, history[i].Url))
		nhst = append(nhst, history[i])
	}

	newHistList := tui.NewList()
	newHistList.AddItems(newItems...)
	history = nhst
	hstJson, _ := json.Marshal(history)
	ioutil.WriteFile(home+directory+"/history.json", hstJson, 0644)
}

func getHistoryBox() tui.Widget {
	lst := tui.NewList()
	hstLst = lst
	for i := 0; i < len(history); i++ {
		lst.AddItems(fmt.Sprintf("%s %s", history[i].Method, history[i].Url))
	}
	lst.OnSelectionChanged(switchToHistReq)
	focWdg[" History "] = append(focWdg[" History "], lst)

	lstBox := tui.NewVBox(lst)
	lstBox.SetBorder(true)
	lstBox.SetSizePolicy(tui.Minimum, tui.Minimum)

	btn := tui.NewButton("SAVE")
	btn.OnActivated(saveRequest)
	focWdg[" History "] = append(focWdg[" History "], btn)
	saveBox := tui.NewVBox(btn)
	saveBox.SetBorder(true)
	saveBox.SetSizePolicy(tui.Minimum, tui.Maximum)

	delbtn := tui.NewButton("DEL")
	delbtn.OnActivated(deleteRequest)
	focWdg[" History "] = append(focWdg[" History "], delbtn)
	delBox := tui.NewVBox(delbtn)
	delBox.SetBorder(true)
	delBox.SetSizePolicy(tui.Minimum, tui.Maximum)

	btnBox := tui.NewHBox(saveBox, delBox)
	btnBox.SetSizePolicy(tui.Minimum, tui.Maximum)

	return tui.NewVBox(lstBox, btnBox)
}

func switchToHistReq(lst *tui.List) {
	sel := lst.SelectedItem()

	var found bool
	var hstReq request
	for i := 0; i < len(history) && !found; i++ {
		if fmt.Sprintf("%s %s", history[i].Method, history[i].Url) == sel {
			hstReq = history[i]
			found = true
		}
	}

	if found {
		curReq.Headers = hstReq.Headers
		curReq.Method = hstReq.Method
		curReq.Body = hstReq.Body
		curReq.Url = hstReq.Url

		curReqWidgets.Method.SetSelected(methodToInt[hstReq.Method])
		curReqWidgets.Url.SetText(hstReq.Url)
		curReqWidgets.Body.SetText(hstReq.Body)

		var headerText string
		for k, v := range hstReq.Headers {
			headerText += fmt.Sprintf("%s: %+v\n", k, v)
		}
		curReqWidgets.Headers.SetText(headerText)
	}
}
