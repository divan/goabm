package main

import (
	"fmt"
	"time"

	"github.com/gizak/termui"
)

// UI represents UI layout.
type UI struct {
	alive *termui.LineChart

	newborn *termui.Gauge
	child   *termui.Gauge
	teen    *termui.Gauge
	adult   *termui.Gauge
	aged    *termui.Gauge
	old     *termui.Gauge
}

func initUI(alive <-chan report) *UI {
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
			data = append(data, c.alive)
			ui.updateChart(data, c.state)
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
	lc.Height = termui.TermHeight() - 3*3
	lc.AxesColor = termui.ColorWhite
	lc.LineColor = termui.ColorYellow | termui.AttrBold

	ui.alive = lc

	ui.newborn = newGauge("Newborn")
	ui.child = newGauge("Child")
	ui.teen = newGauge("Teen")
	ui.adult = newGauge("Adult")
	ui.aged = newGauge("Aged")
	ui.old = newGauge("Old")
}

func (ui *UI) createLayout() {
	termui.Body.AddRows(
		termui.NewRow(
			termui.NewCol(6, 0, ui.newborn),
			termui.NewCol(6, 0, ui.child),
		),
		termui.NewRow(
			termui.NewCol(6, 0, ui.teen),
			termui.NewCol(6, 0, ui.adult),
		),
		termui.NewRow(
			termui.NewCol(6, 0, ui.aged),
			termui.NewCol(6, 0, ui.old),
		),
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

func (ui *UI) updateChart(data []float64, state map[string]int) {
	ui.alive.Data = data

	total := float64(state["newborn"]+state["child"]+state["young"]+state["adult"]+state["aged"]+state["old"]) / 6
	ui.newborn.Percent = int(float64(state["newborn"]) * 100.0 / total)
	ui.child.Percent = int(float64(state["child"]) * 100.0 / total)
	ui.teen.Percent = int(float64(state["teen"]) * 100.0 / total)
	ui.adult.Percent = int(float64(state["adult"]) * 100.0 / total)
	ui.aged.Percent = int(float64(state["aged"]) * 100.0 / total)
	ui.old.Percent = int(float64(state["old"]) * 100.0 / total)

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

func newGauge(text string) *termui.Gauge {
	g := termui.NewGauge()
	g.Percent = 0
	g.Width = termui.TermWidth()
	g.Height = 3
	g.BorderLabel = text
	g.BarColor = termui.ColorRed
	g.BorderFg = termui.ColorWhite
	g.BorderLabelFg = termui.ColorCyan
	return g
}
