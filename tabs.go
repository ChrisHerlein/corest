package main

import tui "github.com/marcusolsson/tui-go"

type uiTab struct {
	label *tui.Label
	view  tui.Widget
}

type tabWidget struct {
	*tui.Box

	views  []*uiTab
	active int
}

func newTabWidget(views ...*uiTab) *tabWidget {
	topbar := tui.NewHBox(tui.NewLabel("> "))
	topbar.SetSizePolicy(tui.Minimum, tui.Maximum)
	view := &tabWidget{views: views}

	for i := 0; i < len(views); i++ {
		topbar.Append(views[i].label)
	}

	topbar.Append(tui.NewSpacer())
	view.style()

	vbox := tui.NewVBox(topbar, views[0].view)
	vbox.SetSizePolicy(tui.Maximum, tui.Preferred)
	view.Box = vbox
	return view
}

func (t *tabWidget) OnKeyEvent(ev tui.KeyEvent) {
	switch ev.Key {
	case tui.KeyCtrlN:
		t.Next()
	case tui.KeyCtrlP:
		t.Previous()
	}

	t.Box.OnKeyEvent(ev)
}

func (t *tabWidget) setView(view tui.Widget) {
	t.Box.Remove(1)
	t.Box.Append(view)
}

func (t *tabWidget) style() {
	for i := 0; i < len(t.views); i++ {
		if i == t.active {
			t.views[i].label.SetStyleName("highlight")
			continue
		}
		t.views[i].label.SetStyleName("tab")
	}
}

func (t *tabWidget) Next() {
	t.active = clamp(t.active+1, 0, len(t.views)-1)
	t.style()
	t.setView(t.views[t.active].view)
	fc.Set(focWdg[t.views[t.active].label.Text()]...)
	ui.SetFocusChain(cfc)
}

func (t *tabWidget) Previous() {
	t.active = clamp(t.active-1, 0, len(t.views)-1)
	t.style()
	t.setView(t.views[t.active].view)
	fc.Set(focWdg[t.views[t.active].label.Text()]...)
	ui.SetFocusChain(cfc)
}

func clamp(n, min, max int) int {
	if n < min {
		return max
	}
	if n > max {
		return min
	}
	return n
}
