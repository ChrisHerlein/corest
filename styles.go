package main

import (
	tui "github.com/marcusolsson/tui-go"
)

var styleLblMethod = tui.Style{
	Fg:      tui.ColorWhite,
	Bg:      tui.ColorBlack,
	Reverse: tui.DecorationOn,
	Bold:    tui.DecorationOn,
}

var sst2 = tui.Style{
	Bg: tui.ColorBlack,
	Fg: tui.ColorWhite,
}

var myTheme *tui.Theme

func initTheme() {
	stl1 := tui.DefaultTheme.Style("list.item.selected")
	stl2 := tui.DefaultTheme.Style("table.cell.selected")
	stl3 := tui.DefaultTheme.Style("button.focused")
	myTheme = tui.NewTheme()
	myTheme.SetStyle("label.highlight", styleLblMethod)
	myTheme.SetStyle("list.item.selected", stl1)
	myTheme.SetStyle("table.cell.selected", stl2)
	myTheme.SetStyle("button.focused", stl3)
	myTheme.SetStyle("box.focused.border", tui.Style{Fg: tui.ColorGreen, Bg: tui.ColorDefault})
}
