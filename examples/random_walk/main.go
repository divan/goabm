package main

import (
	"log"
	"math/rand"
	"syscall"
	"time"
	"unsafe"

	"github.com/divan/goabm/abm"
	"github.com/divan/goabm/models/random_walker"
	"github.com/divan/goabm/ui/term_grid"
	"github.com/divan/goabm/worlds/grid2d"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	a := abm.New()
	w, h := getTermSize()
	grid2D := grid.New(w, h)
	a.SetWorld(grid2D)

	cell, err := walker.New(a, rand.Intn(w-1), rand.Intn(h-1), true)
	if err != nil {
		log.Fatal(err)
	}
	a.AddAgent(cell)
	grid2D.SetCell(cell.X(), cell.Y(), cell)

	a.LimitIterations(1000)

	ch := make(chan [][]interface{})
	a.SetReportFunc(func(a *abm.ABM) {
		ch <- grid2D.Dump(func(a abm.Agent) bool { return a != nil })
	})

	go func() {
		a.StartSimulation()
		close(ch)
	}()

	ui := termgrid.New()
	defer ui.Stop()
	ui.AddGrid(ch)
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
