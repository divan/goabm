package main

import (
	"math/rand"
	"time"

	"github.com/divan/goabm/abm"
	"github.com/divan/goabm/models/conway_life"
	"github.com/divan/goabm/ui/shiny_grid"
	"github.com/divan/goabm/worlds/grid2d"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	a := abm.New()
	w, h := 100, 80
	//w, h := termgrid.TermSize()
	g := grid.New(w, h)
	a.SetWorld(g)

	// populate grid randomly
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			alive := rand.Float64() > 0.5
			cell := life.New(a, x, y, alive)
			a.AddAgent(cell)
			g.SetCell(x, y, cell)
		}
	}

	ch := make(chan [][]interface{})
	a.SetReportFunc(func(a *abm.ABM) {
		ch <- g.Dump(life.IsAlive)
		time.Sleep(10 * time.Millisecond)
	})

	go a.StartSimulation()

	//ui := termgrid.New()
	ui := shiny.New(w, h)
	defer ui.Stop()
	ui.AddGrid(ch)
	ui.Loop()
}
