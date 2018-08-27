package main

import (
	"fmt"
	"strings"

	tui "github.com/marcusolsson/tui-go"
)

var curReq = request{}
var focWdg = make([]tui.Widget, 0)
var ansWidget *tui.TextEdit
var ui tui.UI

func main() {
	var methBox = getMethodBox()
	var urlBox = getUrlBox()
	var ansBox = getAnswerBox()
	var mainBox = tui.NewHBox(methBox, urlBox, ansBox)

	ui, _ = tui.New(mainBox)

	fc.Set(focWdg...)
	ui.SetFocusChain(cfc)
	ui.SetKeybinding("Esc", func() { fmt.Println("quit"); ui.Quit() })
	e := ui.Run()
	if e != nil {
		fmt.Println("e:", e.Error())
	}

	fmt.Println("selected method:", curReq.Method)
	fmt.Println("url:", curReq.Url)
	fmt.Println("ans:", curReq.Answer)

}

func getMethodBox() tui.Widget {
	lbl := tui.NewLabel("Method")
	lst := tui.NewList()

	lst.AddItems(methodList...)
	lst.OnSelectionChanged(changeReqMethod)
	lst.SetSelected(0)
	changeReqMethod(lst)
	focWdg = append(focWdg, lst)

	lblHead := tui.NewLabel("Headers")
	headers := tui.NewTextEdit()
	headers.OnTextChanged(changeReqHeaders)
	focWdg = append(focWdg, headers)

	box := tui.NewVBox(lbl, lst, lblHead, headers)
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

func getUrlBox() tui.Widget {
	lbl := tui.NewLabel("Url")
	entry := tui.NewEntry()
	entry.OnChanged(changeReqUrl)
	focWdg = append(focWdg, entry)

	lbl2 := tui.NewLabel("Body")
	tedit := tui.NewTextEdit()
	tedit.OnTextChanged(changeReqBody)
	focWdg = append(focWdg, tedit)

	btn := tui.NewButton("SEND")
	btn.OnActivated(sendButtonFunc)
	focWdg = append(focWdg, btn)

	var box = tui.NewVBox(lbl, entry, lbl2, tedit, btn)
	return box
}

func changeReqUrl(entry *tui.Entry) {
	curReq.Url = entry.Text()
}

func changeReqBody(entry *tui.TextEdit) {
	curReq.Body = entry.Text()
}

func getAnswerBox() tui.Widget {
	lbl2 := tui.NewLabel("Answer")
	tedit := tui.NewTextEdit()
	tedit.SetText(curReq.Answer)
	ansWidget = tedit
	var box = tui.NewVBox(lbl2, tedit)
	return box
}
