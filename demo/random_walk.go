package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/divan/goabm/abm"
	"github.com/divan/goabm/models/random_walker"
	"github.com/divan/goabm/ui/shiny_grid"
	"github.com/divan/goabm/worlds/grid2d"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	a := abm.New()
	w, h := 300, 200
	grid2D := grid.New(w, h)
	a.SetWorld(grid2D)

	for i := 0; i < 10; i++ {
		cell, err := walker.New(a, rand.Intn(w-1), rand.Intn(h-1), true)
		if err != nil {
			log.Fatal(err)
		}
		a.AddAgent(cell)
		grid2D.SetCell(cell.X(), cell.Y(), cell)
	}

	ch := make(chan [][]interface{})
	a.SetReportFunc(func(a *abm.ABM) {
		ch <- grid2D.Dump(func(a abm.Agent) bool { return a != nil })
	})

	a.LimitIterations(10000)
	go func() {
		a.StartSimulation()
		close(ch)
	}()

	ui := shiny.New(w, h)
	defer ui.Stop()
	ui.AddGrid(ch)
	ui.Loop()
}
