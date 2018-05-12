package main

import (
	"bufio"
	"math/rand"
	"os"
	"syscall"
	"time"
	"unsafe"

	"github.com/divan/goabm/abm"
	"github.com/divan/goabm/models/conway_life"
	"github.com/divan/goabm/ui/term_grid"
	"github.com/divan/goabm/worlds/grid2d"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	a := abm.New()
	w, h := getTermSize()
	grid2D := grid.New(w, h)
	a.SetWorld(grid2D)

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			alive := rand.Float64() > 0.5
			cell := life.New(a, x, y, alive)
			a.AddAgent(cell)
			grid2D.SetCell(x, y, cell)
		}
	}

	a.LimitIterations(1000)

	ch := make(chan [][]bool)
	a.SetReportFunc(func(a *abm.ABM) {
		ch <- grid2D.Dump(life.IsAlive)
		b, _ := bufio.NewReader(os.Stdin).ReadBytes('\n')
		if len(b) > 1 && string(b[0]) == "q" {
			close(ch)
		}
	})

	go func() {
		a.StartSimulation()
		close(ch)
	}()

	ui := termgrid.New(w, h, ch)
	defer ui.Stop()
	ui.Loop()
}

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func getTermSize() (int, int) {
	ws := &winsize{}
	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		panic(errno)
	}
	return int(ws.Col) - 1, int(ws.Row) - 1
}
