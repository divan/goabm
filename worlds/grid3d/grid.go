package grid

import (
	"errors"
	"sync"

	"github.com/divan/goabm/abm"
)

type Grid struct {
	mx                   sync.RWMutex
	width, height, depth int

	cells, cellsPrev []abm.Agent

	nmx sync.RWMutex
}

type Point struct {
	X, Y, Z int
	Color   int
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

	g.cellsPrev = append([]abm.Agent{}, g.cells...)
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

	agent := g.cells[g.idx(fromZ, fromY, fromX)]
	g.cells[g.idx(toZ, toY, toX)] = agent
	//g.cells[fromZ][fromY][fromX] = nil
	return nil
}

func (g *Grid) Cell(x, y, z int) abm.Agent {
	if g.validateXYZ(x, y, z) != nil {
		return nil
	}
	g.mx.RLock()
	defer g.mx.RUnlock()
	return g.cellsPrev[g.idx(z, y, x)]
}

func (g *Grid) SetCell(x, y, z int, c abm.Agent) {
	if err := g.validateXYZ(x, y, z); err != nil {
		panic(err)
	}
	g.mx.Lock()
	g.cells[g.idx(z, y, x)] = c
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

func (g *Grid) DumpFull(fn func(c abm.Agent) bool) [][][]interface{} {
	g.mx.RLock()
	defer g.mx.RUnlock()

	var ret = make([][][]interface{}, g.width)
	for i := 0; i < g.width; i++ {
		ret[i] = make([][]interface{}, g.height)
		for j := 0; j < g.height; j++ {
			ret[i][j] = make([]interface{}, g.depth)
			for k := 0; k < g.depth; k++ {
				a := g.cells[g.idx(i, j, k)]
				ret[i][j][k] = fn(a)
			}
		}
	}
	return ret
}

func (g *Grid) Dump(fn func(c abm.Agent) bool) []interface{} {
	g.mx.RLock()
	defer g.mx.RUnlock()

	var ret = make([]interface{}, 0, g.size())
	for i := 0; i < g.size(); i++ {
		a := g.cells[i]
		if fn(a) {
			x, y, z := g.xyz(i)
			point := Point{
				X: x,
				Y: y,
				Z: z,
			}
			ret = append(ret, point)
		}
	}
	return ret
}

func (g *Grid) size() int {
	return g.depth * g.height * g.width
}

func (g *Grid) idx(x, y, z int) int {
	return z*g.width*g.height + y*g.height + x
}

func (g *Grid) xyz(idx int) (int, int, int) {
	z := idx / (g.height * g.width)
	w := idx % (g.height * g.width)
	y := w / g.depth
	x := w % g.depth
	return x, y, z
}

// just move this verbose initialization here for brevity.
func (g *Grid) initSlices() {
	g.cells = make([]abm.Agent, g.size())
	g.cellsPrev = make([]abm.Agent, g.size())
}
