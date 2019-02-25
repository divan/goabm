package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/divan/goabm/abm"
	"github.com/divan/goabm/ui/shiny_grid"
	"github.com/divan/goabm/worlds/grid2d"
)

func main() {
	nprey := flag.Int("prey", 16, "Number of prey specimen")
	npredator := flag.Int("predator", 16, "Number of predator specimen")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	a := abm.New()
	w, h := 100, 100 //getTermSize()
	g := grid.New(w, h)
	a.SetWorld(g)

	for i := 0; i < *nprey; i++ {
		addPrey(a, g, w, h)
	}
	for i := 0; i < *npredator; i++ {
		addPredator(a, g, w, h)
	}
	addGrass(a, g, w, h)

	ch := make(chan [][]interface{})
	a.SetReportFunc(func(a *abm.ABM) {
		ch <- g.Dump(func(a abm.Agent) bool { return a != nil })
	})

	go func() {
		a.StartSimulation()
		close(ch)
	}()

	ui := shiny.New(100, 100)
	defer ui.Stop()
	ui.AddGrid(ch)
	ui.Loop()
}

func addGrass(a *abm.ABM, g *grid.Grid, w, h int) {
	for i := 0; i < w-1; i++ {
		for j := 0; j < h-1; j++ {
			if g.Cell(i, j) == nil {
				cell := NewGrass(a, w, h)
				a.AddAgent(cell)
				g.SetCell(i, j, cell)
			}
		}
	}
}
func addPrey(a *abm.ABM, g *grid.Grid, w, h int) {
	cell := NewPrey(a, w, h)
	a.AddAgent(cell)
	g.SetCell(cell.X(), cell.Y(), cell)
}
func addPredator(a *abm.ABM, g *grid.Grid, w, h int) {
	cell := NewPredator(a, w, h)
	a.AddAgent(cell)
	g.SetCell(cell.X(), cell.Y(), cell)
}
