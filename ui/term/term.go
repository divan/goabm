package term

import (
	"github.com/divan/goabm/ui"
	"github.com/gizak/termui"
)

// UI represents UI layout. Implements ui.UI, ui.Charts, ui.Grid
type UI struct {
	charts []*termui.LineChart
}

var _ ui.UI = &UI{}
var _ ui.Charts = &UI{}

func NewUI() *UI {
	err := termui.Init()
	if err != nil {
		panic(err)
	}

	ui := &UI{}

	ui.handleKeys()
	ui.Align()
	ui.Render()

	return ui
}

// Stop shuts down terminal graphics. Implements ui.UI.
func (*UI) Stop() {
	termui.Close()
}

// AddChart adds new chart to ui. Implements ui.Chart.
func (ui *UI) AddChart(name string, values <-chan float64) {
	lc := termui.NewLineChart()
	lc.BorderLabel = name
	lc.Data = []float64{}
	lc.AxesColor = termui.ColorWhite
	lc.LineColor = termui.ColorYellow | termui.AttrBold

	ui.charts = append(ui.charts, lc)
	ui.resizeCharts()

	go ui.updateChartData(lc, values)
}

func (ui *UI) resizeCharts() {
	count := len(ui.charts)
	for i := range ui.charts {
		ui.charts[i].Height = termui.TermHeight() / count
	}
}

func (ui *UI) createLayout() {
	var rows []*termui.Row
	for _, chart := range ui.charts {
		row := termui.NewRow(
			termui.NewCol(12, 0, chart),
		)
		rows = append(rows, row)
	}
	termui.Body.AddRows(rows...)
}

// Render rerenders UI.
func (ui *UI) Render() {
	termui.Body.Align()

	termui.Render(termui.Body)
}

func (ui *UI) updateChartData(chart *termui.LineChart, values <-chan float64) {
	var data []float64
	for c := range values {
		// TODO: use circular buffer to evict old data
		data = append(data, float64(c))
		chart.Data = data
		ui.Render()
	}
}

func (ui *UI) handleKeys() {
	// handle key q pressing
	termui.Handle("/sys/kbd/q", func(termui.Event) {
		termui.StopLoop()
	})
}

// Loop starts event loop and blocks.
func (ui *UI) Loop() {
	ui.createLayout()
	termui.Loop()
}

// Align recalculates layout and aligns widgets.
func (ui *UI) Align() {
	termui.Body.Align()
}
