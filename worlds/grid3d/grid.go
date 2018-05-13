package grid

import (
	"errors"
	"sync"

	"github.com/divan/goabm/abm"
)

type Grid struct {
	mx                   sync.RWMutex
	width, height, depth int

	cells, cellsPrev [][][]abm.Agent

	nmx sync.RWMutex
}

func New(width, height, depth int) *Grid {
	g := &Grid{
		width:  width,
		height: height,
		depth:  depth,
	}

	g.initSlices()

	return g
}

// Tick marks beginning of the new time period.
// Implements World interface.
func (g *Grid) Tick() {
	g.mx.RLock()
	defer g.mx.RUnlock()

	for i := 0; i < g.width; i++ {
		for j := 0; j < g.height; j++ {
			for k := 0; k < g.depth; k++ {
				g.cellsPrev[k][j][i] = abm.CopyAgent(g.cells[k][j][i])
			}
		}
	}
}

func (g *Grid) Move(fromX, fromY, fromZ, toX, toY, toZ int) error {
	if err := g.validateXYZ(fromX, fromY, fromZ); err != nil {
		return err
	}
	if err := g.validateXYZ(toX, toY, toZ); err != nil {
		return err
	}
	g.mx.Lock()
	defer g.mx.Unlock()

	agent := g.cells[fromZ][fromY][fromX]
	g.cells[toZ][toY][toX] = agent
	//g.cells[fromZ][fromY][fromX] = nil
	return nil
}

func (g *Grid) Cell(x, y, z int) abm.Agent {
	if g.validateXYZ(x, y, z) != nil {
		return nil
	}
	g.mx.RLock()
	defer g.mx.RUnlock()
	return g.cellsPrev[z][y][x]
}

func (g *Grid) SetCell(x, y, z int, c abm.Agent) {
	if err := g.validateXYZ(x, y, z); err != nil {
		panic(err)
	}
	g.mx.Lock()
	g.cells[z][y][x] = c
	g.mx.Unlock()
}

func (g *Grid) ClearCell(x, y, z int) {
	g.SetCell(x, y, z, nil)
}

func (g *Grid) Width() int {
	return g.width
}

func (g *Grid) Height() int {
	return g.height
}

func (g *Grid) Depth() int {
	return g.depth
}

func (g *Grid) validateXYZ(x, y, z int) error {
	if x < 0 {
		return errors.New("x < 0")
	}
	if x > g.width-1 {
		return errors.New("x > grid width")
	}
	if y < 0 {
		return errors.New("y < 0")
	}
	if y > g.height-1 {
		return errors.New("y > grid height")
	}
	if z < 0 {
		return errors.New("z < 0")
	}
	if z > g.depth-1 {
		return errors.New("z > grid depth")
	}
	return nil
}

func (g *Grid) Dump(fn func(c abm.Agent) bool) [][]bool {
	g.mx.RLock()
	defer g.mx.RUnlock()

	var ret = make([][]bool, g.height)
	/*
		for i := 0; i < g.height; i++ {
			ret[i] = make([]bool, g.width)
			for j := 0; j < g.width; j++ {
				a := g.cells[i][j]
				ret[i][j] = fn(a)
			}
		}
	*/
	return ret
}

// just move this verbose initialization here for brevity.
func (g *Grid) initSlices() {
	g.cells = make([][][]abm.Agent, g.depth)
	g.cellsPrev = make([][][]abm.Agent, g.depth)
	for j := 0; j < g.depth; j++ {
		g.cells[j] = make([][]abm.Agent, g.height)
		g.cellsPrev[j] = make([][]abm.Agent, g.height)
		for i := 0; i < g.height; i++ {
			g.cells[j][i] = make([]abm.Agent, g.width)
			g.cellsPrev[j][i] = make([]abm.Agent, g.width)
		}
	}
}
