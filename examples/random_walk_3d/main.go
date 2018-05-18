package main

import (
	"flag"
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

func (w *Walker) Run() {
	rx := rand.Intn(6)
	oldx, oldy, oldz := w.x, w.y, w.z
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
	err := w.abm.World().(*grid.Grid).Copy(oldx, oldy, oldz, w.x, w.y, w.z)
	if err != nil {
		w.x, w.y, w.z = oldx, oldy, oldz
	}
}

func main() {
	var n = flag.Int("n", 1, "Number of agents so start")
	var size = flag.Int("size", 100, "3D grid edge size")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
	a := abm.New()
	w, h, d := *size, *size, *size
	g := grid.New(w, h, d)
	a.SetWorld(g)

	for i := 0; i < *n; i++ {
		cell := NewWalker(a, rand.Intn(w), rand.Intn(h), rand.Intn(d))
		a.AddAgent(cell)
		g.SetCell(cell.x, cell.y, cell.z, cell)
	}

	a.LimitIterations(10000)

	ch := make(chan []interface{})
	a.SetReportFunc(func(a *abm.ABM) {
		ch <- g.Dump(func(a abm.Agent) bool { return a != nil })
	})

	go func() {
		time.Sleep(1 * time.Second)
		a.StartSimulation()
		close(ch)
	}()

	ui3d := ui.New(w, h, d)
	defer ui3d.Stop()
	ui3d.AddGrid3D(ch)
	ui3d.Loop()
}
