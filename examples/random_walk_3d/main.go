package main

import (
	"math/rand"
	"time"

	"github.com/divan/goabm/abm"
	ui "github.com/divan/goabm/ui/webgl_grid3d"
	"github.com/divan/goabm/worlds/grid3d"
)

// Walker implements abm.Agent
type Walker struct {
	x, y, z             int
	origx, origy, origz int
	abm                 *abm.ABM
}

func NewWalker(abm *abm.ABM, x, y, z int) *Walker {
	return &Walker{
		origx: x,
		origy: y,
		origz: z,
		x:     x,
		y:     y,
		z:     z,
		abm:   abm,
	}
}

func (w *Walker) Run(i int) {
	rx := rand.Intn(6)
	//	oldx, oldy, oldz := w.x, w.y, w.z
	switch rx {
	case 0:
		w.x++
	case 1:
		w.y++
	case 2:
		w.x--
	case 3:
		w.y--
	case 4:
		w.z++
	case 5:
		w.z--
	}
	//w.abm.World().(*grid.Grid).Move(oldx, oldy, oldz, w.x, w.y, w.z)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	a := abm.New()
	w, h, d := 100, 100, 100
	g := grid.New(w, h, d)
	//a.SetWorld(g)

	cell := NewWalker(a, 50, 50, 50)
	a.AddAgent(cell)
	g.SetCell(cell.x, cell.y, cell.z, cell)

	a.LimitIterations(1000000)

	ch := make(chan ui.Coord)
	a.SetReportFunc(func(a *abm.ABM) {
		ch <- ui.Coord{
			X: cell.x,
			Y: cell.y,
			Z: cell.z,
		}
	})

	go func() {
		time.Sleep(5 * time.Second)
		a.StartSimulation()
		close(ch)
	}()

	ui3d := ui.New(w, h, d, ch)
	defer ui3d.Stop()
	ui3d.Loop()
}
