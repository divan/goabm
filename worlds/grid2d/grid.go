package grid

import (
	"errors"
	"sync"

	"github.com/divan/goabm/abm"
)

type Grid struct {
	mx            sync.RWMutex
	width, height int

	cells, cellsPrev []abm.Agent
}

func New(width, height int) *Grid {
	g := &Grid{
		width:  width,
		height: height,
	}

	g.initSlices()

	return g
}

// Tick marks beginning of the new time period.
// Implements World interface.
func (g *Grid) Tick() {
	g.mx.RLock()
	defer g.mx.RUnlock()

	for i := 0; i < g.size(); i++ {
		g.cellsPrev[i] = abm.CopyAgent(g.cells[i])
	}

}

func (g *Grid) Move(fromX, fromY, toX, toY int) error {
	if err := g.validateXY(fromX, fromY); err != nil {
		return err
	}
	if err := g.validateXY(toX, toY); err != nil {
		return err
	}
	g.mx.Lock()
	defer g.mx.Unlock()

	agent := g.cells[g.idx(fromX, fromY)]
	g.cells[g.idx(toX, toY)] = agent
	g.cells[g.idx(fromX, fromY)] = nil
	return nil
}

func (g *Grid) Copy(fromX, fromY, toX, toY int) error {
	if err := g.validateXY(fromX, fromY); err != nil {
		return err
	}
	if err := g.validateXY(toX, toY); err != nil {
		return err
	}
	g.mx.Lock()
	defer g.mx.Unlock()

	agent := g.cells[g.idx(fromX, fromY)]
	g.cells[g.idx(toX, toY)] = agent
	return nil
}

func (g *Grid) Cell(x, y int) abm.Agent {
	if g.validateXY(x, y) != nil {
		return nil
	}
	g.mx.RLock()
	defer g.mx.RUnlock()
	return g.cells[g.idx(x, y)]
}

func (g *Grid) SetCell(x, y int, c abm.Agent) {
	if err := g.validateXY(x, y); err != nil {
		panic(err)
	}
	g.mx.Lock()
	g.cells[g.idx(x, y)] = c
	g.mx.Unlock()
}

func (g *Grid) ClearCell(x, y int) {
	g.SetCell(x, y, nil)
}

func (g *Grid) Width() int {
	return g.width
}

func (g *Grid) Height() int {
	return g.height
}

func (g *Grid) validateXY(x, y int) error {
	if x < 0 {
		return errors.New("x < 0")
	}
	if y < 0 {
		return errors.New("y < 0")
	}
	if x > g.width-1 {
		return errors.New("x > grid width")
	}
	if y > g.height-1 {
		return errors.New("y > grid height")
	}
	return nil
}

func (g *Grid) Dump(fn func(c abm.Agent) bool) [][]interface{} {
	g.mx.RLock()
	defer g.mx.RUnlock()

	var ret = make([][]interface{}, g.width)
	for i := 0; i < g.width; i++ {
		ret[i] = make([]interface{}, g.height)
		for j := 0; j < g.height; j++ {
			a := g.cells[g.idx(i, j)]
			ret[i][j] = fn(a)
		}
	}
	return ret
}

// just move this verbose initialization here for brevity.
func (g *Grid) initSlices() {
	g.cells = make([]abm.Agent, g.size())
	g.cellsPrev = make([]abm.Agent, g.size())
}

func (g *Grid) size() int {
	return g.height * g.width
}

func (g *Grid) idx(x, y int) int {
	return y*g.width + x
}
