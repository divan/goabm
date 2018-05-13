package grid

import (
	"errors"
	"reflect"
	"sync"

	"github.com/divan/goabm/abm"
)

type Grid struct {
	mx            sync.RWMutex
	width, height int

	cells, cellsPrev [][]abm.Agent

	nmx sync.RWMutex
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

	for i := 0; i < g.width; i++ {
		for j := 0; j < g.height; j++ {
			g.cellsPrev[j][i] = copyAgent(g.cells[j][i])
		}
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

	agent := g.cells[fromY][fromX]
	g.cells[toY][toX] = agent
	//g.cells[fromY][fromX] = nil
	return nil
}

func copyAgent(src abm.Agent) abm.Agent {
	if src == nil {
		return nil
	}
	typ := reflect.TypeOf(src)
	val := reflect.ValueOf(src)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}
	elem := reflect.New(typ).Elem()
	elem.Set(val)
	return elem.Addr().Interface().(abm.Agent)
}

func (g *Grid) Cell(x, y int) abm.Agent {
	if g.validateXY(x, y) != nil {
		return nil
	}
	g.mx.RLock()
	defer g.mx.RUnlock()
	return g.cellsPrev[y][x]
}

func (g *Grid) SetCell(x, y int, c abm.Agent) {
	if err := g.validateXY(x, y); err != nil {
		panic(err)
	}
	g.mx.Lock()
	g.cells[y][x] = c
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

func (g *Grid) Dump(fn func(c abm.Agent) bool) [][]bool {
	g.mx.RLock()
	defer g.mx.RUnlock()

	var ret = make([][]bool, g.height)
	for i := 0; i < g.height; i++ {
		ret[i] = make([]bool, g.width)
		for j := 0; j < g.width; j++ {
			a := g.cells[i][j]
			ret[i][j] = fn(a)
		}
	}
	return ret
}

// just move this verbose initialization here for brevity.
func (g *Grid) initSlices() {
	g.cells = make([][]abm.Agent, g.height)
	g.cellsPrev = make([][]abm.Agent, g.height)
	for i := 0; i < g.height; i++ {
		g.cells[i] = make([]abm.Agent, g.width)
		g.cellsPrev[i] = make([]abm.Agent, g.width)
	}
}
