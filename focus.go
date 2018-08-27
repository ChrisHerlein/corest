package main

import tui "github.com/marcusolsson/tui-go"

//var fc = customFC{
//	wdg: make([]tui.Widget, 0),
//	hst: make([]int, 0),
//}

var fc = tui.SimpleFocusChain{}
var cfc = customFC{}

type customFC struct {
}

func (c customFC) FocusNext(w tui.Widget) tui.Widget {
	return fc.FocusNext(w)
}

func (c customFC) FocusPrev(w tui.Widget) tui.Widget {
	return fc.FocusPrev(w)
}

func (c customFC) FocusDefault() tui.Widget {
	return fc.FocusDefault()
}
