package life

import (
	"sync"

	"github.com/divan/goabm/abm"
	"github.com/divan/goabm/worlds/grid2d"
)

// Cell implements Agent interface for game of life's cell.
type Cell struct {
	mx    sync.RWMutex
	alive bool

	abm  *abm.ABM
	grid *grid.Grid

	x, y int
}

// New returns a new cell and places it onto grid.
func New(abm *abm.ABM, x, y int, alive bool) *Cell {
	grid, ok := abm.World().(*grid.Grid)
	if !ok {
		panic("expecting Grid2D world for this model")
	}
	cell := &Cell{
		alive: alive,

		abm:  abm,
		grid: grid,

		x: x,
		y: y,
	}

	return cell
}

// Run runs single iteration over cell.
//
// Any live cell with fewer than two live neighbors dies, as if caused by under population.
// Any live cell with two or three live neighbors lives on to the next generation.
// Any live cell with more than three live neighbors dies, as if by overpopulation.
// Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
func (c *Cell) Run() {
	neighbors := c.CountNeighbors()
	if c.IsAlive() {
		if neighbors < 2 || neighbors > 3 {
			c.Die()
		}
	} else {
		if neighbors == 3 {
			c.Reborn()
		}
	}
}

func (c *Cell) CountNeighbors() int {
	c.mx.RLock()
	defer c.mx.RUnlock()
	return c.grid.CountNeighbors(c.x, c.y, IsAlive)
}

func IsAlive(a abm.Agent) bool {
	if a == nil {
		return false
	}

	c, ok := a.(*Cell)
	return ok && c.IsAlive()
}

func (c *Cell) IsAlive() bool {
	c.mx.RLock()
	defer c.mx.RUnlock()
	return c.alive
}

func (c *Cell) Die() {
	c.mx.Lock()
	c.alive = false
	c.mx.Unlock()
}

func (c *Cell) Reborn() {
	c.mx.Lock()
	c.alive = true
	c.mx.Unlock()
}
