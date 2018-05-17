package main

import (
	"math/rand"
	"time"

	"github.com/divan/goabm/abm"
	"github.com/divan/goabm/ui/shiny_grid"
	"github.com/divan/goabm/worlds/grid2d"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	a := abm.New()
	w, h := 100, 80
	grid2D := grid.New(w, h)
	a.SetWorld(grid2D)

	cell := NewWalker(a, 50, 40)
	a.AddAgent(cell)
	grid2D.SetCell(cell.x, cell.y, cell)

	a.LimitIterations(10000)

	ch := make(chan [][]interface{})
	a.SetReportFunc(func(a *abm.ABM) {
		ch <- grid2D.Dump(func(a abm.Agent) bool { return a != nil })
	})

	go func() {
		a.StartSimulation()
		close(ch)
	}()

	ui := shiny.New(w, h)
	defer ui.Stop()
	ui.AddGrid(ch)
	ui.Loop()
}
