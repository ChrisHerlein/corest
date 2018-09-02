package main

import (
	"fmt"
	"os"
	"strings"

	tui "github.com/marcusolsson/tui-go"
)

var curReq = request{
	Method: "GET",
}

var (
	focWdg = map[string][]tui.Widget{
		" History ": make([]tui.Widget, 0),
		" Request ": make([]tui.Widget, 0),
		" Answer ":  make([]tui.Widget, 0),
	}
	scr        *tui.ScrollArea
	ansHeaders *tui.TextEdit
	ansBody    *tui.TextEdit
	ui         tui.UI
)

func init() {
	home := os.Getenv("HOME")
	ensurePath(home)
	loadHistory()
}

func main() {
	initTheme()
	var methBox = getMethodBox()
	var reqBox = getReqBox()
	var ansBox = getAnswerBox()
	var hstBox = getHistoryBox()

	tabLayout := newTabWidget(
		&uiTab{label: tui.NewLabel(" Request "), view: reqBox},
		&uiTab{label: tui.NewLabel(" Answer "), view: ansBox},
		&uiTab{label: tui.NewLabel(" History "), view: hstBox},
	)

	var mainBox = tui.NewVBox(methBox, tabLayout)

	ui, _ = tui.New(mainBox)

	fc.Set(focWdg[" Request "]...)
	ui.SetTheme(myTheme)
	ui.SetFocusChain(cfc)
	ui.SetKeybinding("Esc", func() { fmt.Println("quit"); ui.Quit() })
	ui.SetKeybinding("P", func() {
		if scr.IsFocused() {
			scr.Scroll(0, -1)
		}
	})
	ui.SetKeybinding("N", func() {
		if scr.IsFocused() {
			scr.Scroll(0, 1)
		}
	})

	e := ui.Run()
	if e != nil {
		fmt.Println("e:", e.Error())
	}

}

func getMethodBox() tui.Widget {

	lst := tui.NewList()
	lst.AddItems(methodList...)
	lst.OnSelectionChanged(changeReqMethod)
	lst.SetSelected(0)
	lst.SetSizePolicy(tui.Minimum, tui.Minimum)
	curReqWidgets.Method = lst
	focWdg[" History "] = append(focWdg[" History "], lst)
	focWdg[" Request "] = append(focWdg[" Request "], lst)
	focWdg[" Answer "] = append(focWdg[" Answer "], lst)

	lstBox := tui.NewHBox(lst)
	lstBox.SetTitle("Method")
	lstBox.SetBorder(true)
	lstBox.SetSizePolicy(tui.Maximum, tui.Maximum)

	entry := tui.NewEntry()
	entry.OnChanged(changeReqUrl)
	curReqWidgets.Url = entry
	focWdg[" History "] = append(focWdg[" History "], entry)
	focWdg[" Request "] = append(focWdg[" Request "], entry)
	focWdg[" Answer "] = append(focWdg[" Answer "], entry)

	entryBox := tui.NewHBox(entry)
	entryBox.SetTitle("Url")
	entryBox.SetBorder(true)
	entryBox.SetSizePolicy(tui.Minimum, tui.Maximum)

	btn := tui.NewButton("SEND")
	btn.OnActivated(sendButtonFunc)
	focWdg[" History "] = append(focWdg[" History "], btn)
	focWdg[" Request "] = append(focWdg[" Request "], btn)
	focWdg[" Answer "] = append(focWdg[" Answer "], btn)

	btnBox := tui.NewHBox(btn)
	btnBox.SetBorder(true)
	btnBox.SetSizePolicy(tui.Maximum, tui.Maximum)

	box := tui.NewHBox(lstBox, entryBox, btnBox)
	box.SetBorder(true)
	box.SetSizePolicy(tui.Minimum, tui.Maximum)
	return box
}

func changeReqMethod(lst *tui.List) {
	curReq.Method = methodList[lst.Selected()]
}

func changeReqHeaders(headers *tui.TextEdit) {
	curReq.Headers = make(map[string]string)
	txt := headers.Text()
	lines := strings.Split(txt, "\n")
	for i := 0; i < len(lines); i++ {
		parts := strings.Split(lines[i], " ")
		if len(parts) == 2 {
			curReq.Headers[parts[0]] = parts[1]
		}
	}
}

func getReqBox() tui.Widget {

	headers := tui.NewTextEdit()
	headers.OnTextChanged(changeReqHeaders)
	curReqWidgets.Headers = headers
	focWdg[" Request "] = append(focWdg[" Request "], headers)

	headersBox := tui.NewVBox(headers)
	headersBox.SetSizePolicy(tui.Minimum, tui.Maximum)
	headersBox.SetTitle("Headers")
	headersBox.SetBorder(true)

	tedit := tui.NewTextEdit()
	tedit.OnTextChanged(changeReqBody)
	curReqWidgets.Body = tedit
	focWdg[" Request "] = append(focWdg[" Request "], tedit)

	var bodyBox = tui.NewVBox(tedit)
	bodyBox.SetSizePolicy(tui.Maximum, tui.Minimum)
	bodyBox.SetTitle("Body")
	bodyBox.SetBorder(true)

	var box = tui.NewVBox(headersBox, bodyBox)
	return box
}

func changeReqUrl(entry *tui.Entry) {
	curReq.Url = entry.Text()
}

func changeReqBody(entry *tui.TextEdit) {
	curReq.Body = entry.Text()
}

func getAnswerBox() tui.Widget {

	ansHeaders = tui.NewTextEdit()

	headersBox := tui.NewVBox(ansHeaders)
	headersBox.SetSizePolicy(tui.Minimum, tui.Maximum)
	headersBox.SetTitle("Headers")
	headersBox.SetBorder(true)

	ansBody = tui.NewTextEdit()
	scr = tui.NewScrollArea(ansBody)
	focWdg[" Answer "] = append(focWdg[" Answer "], scr)

	var bodyBox = tui.NewVBox(scr)
	bodyBox.SetSizePolicy(tui.Maximum, tui.Minimum)
	bodyBox.SetTitle("Body")
	bodyBox.SetBorder(true)

	var box = tui.NewVBox(headersBox, bodyBox)
	return box
}
