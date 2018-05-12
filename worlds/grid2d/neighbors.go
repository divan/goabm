package grid

import (
	"fmt"

	"github.com/divan/goabm/abm"
)

type Neighborhood uint8

const (
	Euclidian Neighborhood = iota
	Moore
	VonNeumann
)

func (g *Grid) CountNeighbors(x, y int, fn func(a abm.Agent) bool) int {
	if g.validateXY(x, y) != nil {
		return 0
	}
	g.nmx.RLock()
	defer g.nmx.RUnlock()

	var count int
	// assuming Moore neigborhood for now
	corners := []abm.Agent{
		g.left(x, y),
		g.right(x, y),
		g.top(x, y),
		g.bottom(x, y),
		g.topLeft(x, y),
		g.topRight(x, y),
		g.bottomLeft(x, y),
		g.bottomRight(x, y),
	}
	for _, corner := range corners {
		if fn(corner) {
			count++
		}
	}

	return count
}

func (g *Grid) left(x, y int) abm.Agent {
	if x == 0 {
		return nil
	}
	return g.cellsPrev[y][x-1]
}
func (g *Grid) right(x, y int) abm.Agent {
	if x >= g.width-1 {
		return nil
	}
	return g.cellsPrev[y][x+1]
}
func (g *Grid) top(x, y int) abm.Agent {
	if y == 0 {
		return nil
	}
	return g.cellsPrev[y-1][x]
}
func (g *Grid) bottom(x, y int) abm.Agent {
	if y >= g.height-1 {
		return nil
	}
	return g.cellsPrev[y+1][x]
}
func (g *Grid) topLeft(x, y int) abm.Agent {
	if y == 0 || x == 0 {
		return nil
	}
	return g.cellsPrev[y-1][x-1]
}
func (g *Grid) bottomLeft(x, y int) abm.Agent {
	if y >= g.height-1 || x == 0 {
		return nil
	}
	return g.cellsPrev[y+1][x-1]
}
func (g *Grid) topRight(x, y int) abm.Agent {
	if y == 0 || x >= g.width-1 {
		return nil
	}
	return g.cellsPrev[y-1][x+1]
}
func (g *Grid) bottomRight(x, y int) abm.Agent {
	if y >= g.height-1 || x >= g.width-1 {
		return nil
	}
	return g.cellsPrev[y+1][x+1]
}

func (g *Grid) PrintNeighbors(fn func(a abm.Agent) bool) {
	for i := 0; i < g.width; i++ {
		for j := 0; j < g.height; j++ {
			n := g.CountNeighbors(i, j, fn)
			fmt.Printf("%d", n)
		}
		fmt.Println()
	}
}
