package life

import (
	"fmt"
	"testing"

	"github.com/divan/goabm/abm"
	"github.com/divan/goabm/worlds/grid2d"
	. "github.com/smartystreets/goconvey/convey"
)

func TestBlinker(t *testing.T) {
	Convey("Blinker should blink", t, func() {
		a := abm.New()
		g := grid.New(5, 5)
		a.SetWorld(g)

		points := []struct{ x, y int }{{1, 2}, {2, 2}, {3, 2}}

		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				var alive bool
				for _, point := range points {
					if i == point.x && j == point.y {
						alive = true
					}
				}
				cell := New(a, i, j, alive)
				a.AddAgent(cell)
				g.SetCell(i, j, cell)
			}
		}

		PrintDump(g.Dump(IsAlive), 5, 5)
		a.LimitIterations(1)
		a.StartSimulation()
		PrintDump(g.Dump(IsAlive), 5, 5)

		expectedAlive := []struct{ x, y int }{{2, 1}, {2, 2}, {2, 3}}
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				for _, alive := range expectedAlive {
					if i == alive.x && j == alive.y {
						So(g.Cell(i, j).(*Cell).IsAlive(), ShouldBeTrue)
					}
				}
			}
		}
	})
}

/*
func TestStillLifes(t *testing.T) {
	Convey("Still life figures should live", t, func() {
		g := New(4, 4)
		a := &TestAgent{}
		g.SetCell(1, 1, a)
		g.SetCell(2, 1, a)
		g.SetCell(1, 2, a)
		g.SetCell(2, 2, a)

		g.Tick()
		So(g.Cell(0, 0), ShouldBeNil)
		So(g.Cell(3, 3), ShouldBeNil)
		So(g.Cell(0, 3), ShouldBeNil)
		So(g.Cell(0, 3), ShouldBeNil)
		So(g.Cell(1, 1), ShouldNotBeNil)
		So(g.Cell(2, 1), ShouldNotBeNil)
		So(g.Cell(1, 2), ShouldNotBeNil)
		So(g.Cell(2, 2), ShouldNotBeNil)
	})
}
*/

func PrintDump(dump [][]interface{}, w, h int) {
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			if dump[j][i].(bool) {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}
