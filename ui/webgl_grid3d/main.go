package ui

import (
	"log"

	"github.com/divan/goabm/ui"
)

type Coord struct {
	X, Y, Z int
}

type UI struct {
	ch   <-chan []interface{}
	done chan struct{}
	ws   *WSServer
}

var _ ui.UI = &UI{}
var _ ui.Grid3D = &UI{}

func New(w, h, d int) *UI {
	ws := NewWSServer()
	ui := &UI{
		done: make(chan struct{}),
		ws:   ws,
	}
	return ui
}

func (ui *UI) Stop() {
	ui.done <- struct{}{}
}

func (ui *UI) Loop() {
	log.Printf("Starting web server...")
	go ui.startWeb(ui.ws)
	<-ui.done
}

func (ui *UI) AddGrid3D(ch <-chan []interface{}) {
	ui.ch = ch
}
