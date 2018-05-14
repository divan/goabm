package grid

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// TestAgent implements abm.Agent interface for tests.
type TestAgent struct{}

func (*TestAgent) Run(i int) {}

func TestTick(t *testing.T) {
	Convey("Tick should update double buffer", t, func() {
		g := New(100, 100)
		a := &TestAgent{}

		So(g.Cell(50, 50), ShouldBeNil)
		g.SetCell(50, 50, a)
		So(g.Cell(50, 50), ShouldBeNil)
		g.Tick()
		So(g.Cell(50, 50), ShouldEqual, a)
	})
}

func BenchmarkTick(b *testing.B) {
	g := New(100, 100)
	for i := 0; i < b.N; i++ {
		g.Tick()
	}
}
