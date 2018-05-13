package grid

import ()

// TestAgent implements abm.Agent interface for tests.
type TestAgent struct{}

func (*TestAgent) Run(i int) {}

/*
func TestGrid(t *testing.T) {
	Convey("Grid 3.3 should init properly", t, func() {
		g := New(3, 3)
		So(g.Width(), ShouldEqual, 3)
		So(g.Height(), ShouldEqual, 3)
		So(len(g.cells), ShouldEqual, 3)

		Convey("SetCell should update neigbors properly", func() {
			a := &TestAgent{}
			g.SetCell(1, 1, a)
			fn := func(a abm.Agent) bool {
				if a == nil {
					return false
				}

			}
			So(g.NumNeighbors(0, 0), ShouldEqual, 1)
			So(g.NumNeighbors(1, 0), ShouldEqual, 1)
			So(g.NumNeighbors(2, 0), ShouldEqual, 1)
			So(g.NumNeighbors(0, 1), ShouldEqual, 1)
			So(g.NumNeighbors(1, 1), ShouldEqual, 0)
			So(g.NumNeighbors(2, 1), ShouldEqual, 1)
			So(g.NumNeighbors(1, 2), ShouldEqual, 1)
			So(g.NumNeighbors(1, 2), ShouldEqual, 1)
			So(g.NumNeighbors(2, 2), ShouldEqual, 1)

			g.Tick()
			g.SetCell(2, 2, a)
			So(g.NumNeighbors(1, 1), ShouldEqual, 1)
			So(g.NumNeighbors(2, 1), ShouldEqual, 2)
			So(g.NumNeighbors(1, 2), ShouldEqual, 2)
			So(g.NumNeighbors(2, 2), ShouldEqual, 1)
		})
	})
}
*/
