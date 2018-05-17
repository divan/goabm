package grid

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// TestAgent implements abm.Agent interface for tests.
type TestAgent struct{}

func (*TestAgent) Run(i int) {}

func BenchmarkTick(b *testing.B) {
	g := New(100, 100, 100)
	for i := 0; i < b.N; i++ {
		g.Tick()
	}
}

func TestIndexes(t *testing.T) {
	Convey("Indexes should be calculated correctly", t, func() {
		g := New(10, 10, 10)
		So(g.idx(0, 1, 0), ShouldEqual, 10)
		So(g.idx(1, 1, 1), ShouldEqual, 111)
		So(g.idx(9, 9, 9), ShouldEqual, 999)
	})
	Convey("Reverse indexes should be calculated correctly", t, func() {
		g := New(10, 10, 10)
		x, y, z := g.xyz(999)
		So(x, ShouldEqual, 9)
		So(y, ShouldEqual, 9)
		So(z, ShouldEqual, 9)
		x, y, z = g.xyz(10)
		So(x, ShouldEqual, 0)
		So(y, ShouldEqual, 1)
		So(z, ShouldEqual, 0)
		x, y, z = g.xyz(2)
		So(x, ShouldEqual, 2)
		So(y, ShouldEqual, 0)
		So(z, ShouldEqual, 0)
	})
}
