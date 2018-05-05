package main

import (
	"fmt"
	"time"

	"github.com/gizak/termui"
)

// UI represents UI layout.
type UI struct {
	alive *termui.LineChart
}

func initUI(alive <-chan float64) *UI {
	err := termui.Init()
	if err != nil {
		panic(err)
	}

	ui := &UI{}
	ui.createCharts()
	ui.createLayout()

	ui.Align()
	ui.Render()

	go func() {
		var data []float64
		for c := range alive {
			data = append(data, c)
			ui.updateChart(data)
		}
	}()

	return ui
}

func stopUI() {
	termui.Close()
}

func (ui *UI) createCharts() {
	// alive
	lc := termui.NewLineChart()
	lc.BorderLabel = "Humans Alive"
	lc.Data = []float64{}
	lc.Height = termui.TermHeight()
	lc.AxesColor = termui.ColorWhite
	lc.LineColor = termui.ColorYellow | termui.AttrBold

	ui.alive = lc
}

func (ui *UI) createLayout() {
	termui.Body.AddRows(
		termui.NewRow(
			termui.NewCol(12, 0, ui.alive),
		),
	)
}

// Render rerenders UI.
func (ui *UI) Render() {
	termui.Body.Align()

	termui.Render(termui.Body)
}

func (ui *UI) updateChart(data []float64) {
	ui.alive.Data = data
	ui.Render()
}

func (ui *UI) HandleKeys() {
	// handle key q pressing
	termui.Handle("/sys/kbd/q", func(termui.Event) {
		termui.StopLoop()
	})
}

// AddTimer adds handler for repeatable functions that interact with UI.
func (ui *UI) AddTimer(d time.Duration, fn func(e termui.Event)) {
	durationStr := fmt.Sprintf("/timer/%s", d)
	termui.Handle(durationStr, fn)
}

// Loop starts event loop and blocks.
func (ui *UI) Loop() {
	termui.Loop()
}

// Align recalculates layout and aligns widgets.
func (ui *UI) Align() {
	termui.Body.Align()
}
