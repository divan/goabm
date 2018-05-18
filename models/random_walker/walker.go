package walker

import (
	"errors"
	"math/rand"

	"github.com/divan/goabm/abm"
	grid "github.com/divan/goabm/worlds/grid2d"
)

// Walker implements abm.Agent and worlds.XY and
// walks randomly over 2D grid.
type Walker struct {
	x, y         int
	origx, origy int
	grid         *grid.Grid
	trail        bool // leave trail?
}

func New(abm *abm.ABM, x, y int, trail bool) (*Walker, error) {
	world := abm.World()
	if world == nil {
		return nil, errors.New("Walker needs a World defined to operate")
	}
	grid, ok := world.(*grid.Grid)
	if !ok {
		return nil, errors.New("Walker needs a Grid world to operate")
	}
	return &Walker{
		origx: x,
		origy: y,
		x:     x,
		y:     y,
		grid:  grid,
		trail: trail,
	}, nil
}

func (w *Walker) Run() {
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

	var err error
	if w.trail {
		err = w.grid.Copy(oldx, oldy, w.x, w.y)
	} else {
		err = w.grid.Move(oldx, oldy, w.x, w.y)
	}

	if err != nil {
		w.x, w.y = oldx, oldy
	}
}

func (w *Walker) X() int { return w.x }
func (w *Walker) Y() int { return w.y }
