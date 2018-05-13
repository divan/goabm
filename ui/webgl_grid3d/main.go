package ui

import (
	"log"
)

type Coord struct {
	X, Y, Z int
}

type UI struct {
	ch   chan Coord
	done chan struct{}
	ws   *WSServer
}

func New(w, h, d int, ch chan Coord) *UI {
	ws := NewWSServer()
	log.Printf("Starting web server...")
	ui := &UI{
		ch:   ch,
		done: make(chan struct{}),
		ws:   ws,
	}
	go ui.startWeb(ws)
	return ui
}

func (ui *UI) Stop() {
	ui.done <- struct{}{}
}

func (ui *UI) Loop() {
	<-ui.done
}
