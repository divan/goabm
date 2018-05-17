package main

import (
	"bufio"
	"math/rand"
	"os"
	"time"

	"github.com/divan/goabm/abm"
	"github.com/divan/goabm/models/conway_life"
	"github.com/divan/goabm/ui/term_grid"
	"github.com/divan/goabm/worlds/grid2d"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	a := abm.New()
	w, h := termSize()
	grid2D := grid.New(w, h)
	a.SetWorld(grid2D)

	// populate grid randomly
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			alive := rand.Float64() > 0.5
			cell := life.New(a, x, y, alive)
			a.AddAgent(cell)
			grid2D.SetCell(x, y, cell)
		}
	}

	a.LimitIterations(1000)

	ch := make(chan [][]interface{})
	a.SetReportFunc(func(a *abm.ABM) {
		ch <- grid2D.Dump(life.IsAlive)
		if waitForEnter() {
			close(ch)
		}
	})

	go func() {
		a.StartSimulation()
		close(ch)
	}()

	ui := termgrid.New()
	defer ui.Stop()
	ui.AddGrid(ch)
	ui.Loop()
}

func waitForEnter() bool {
	b, _ := bufio.NewReader(os.Stdin).ReadBytes('\n')
	return len(b) > 1 && string(b[0]) == "q"
}
