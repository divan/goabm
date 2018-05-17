package main

import (
	"math/rand"

	"github.com/divan/goabm/abm"
	grid "github.com/divan/goabm/worlds/grid2d"
)

// Walker implements abm.Agent and worlds.XY
type Walker struct {
	x, y         int
	origx, origy int
	abm          *abm.ABM
}

func NewWalker(abm *abm.ABM, x, y int) *Walker {
	return &Walker{
		origx: x,
		origy: y,
		x:     x,
		y:     y,
		abm:   abm,
	}
}

func (w *Walker) Run(i int) {
	rx := rand.Intn(4)
	oldx, oldy := w.x, w.y
	switch rx {
	case 0:
		w.x++
	case 1:
		w.y++
	case 2:
		w.x--
	case 3:
		w.y--
	}
	w.abm.World().(*grid.Grid).Move(oldx, oldy, w.x, w.y)
}

func (w *Walker) X() int { return w.x }
func (w *Walker) Y() int { return w.y }
