package termgrid

import (
	"fmt"
	"log"
	"syscall"
	"unsafe"

	"github.com/divan/goabm/ui"
)

type UI struct {
	width, height int
	ch            <-chan [][]interface{}
}

var _ ui.UI = &UI{}
var _ ui.Grid = &UI{}

func New() *UI {
	w, h := TermSize()
	return &UI{
		width:  w,
		height: h,
	}
}

func (*UI) Stop() {}

func (ui *UI) AddGrid(ch <-chan [][]interface{}) {
	ui.ch = ch
}

func (ui *UI) Loop() {
	for dump := range ui.ch {
		for i := 0; i < ui.height; i++ {
			fmt.Printf("\033[0;0H")
		}
		for i := 0; i < ui.height; i++ {
			for j := 0; j < ui.width; j++ {
				val, ok := dump[j][i].(bool)
				if !ok {
					log.Fatal("Expecting bool for grid dump")
				}
				if val {
					fmt.Printf("*")
				} else {
					fmt.Printf(" ")
				}
			}
			fmt.Println()
		}
	}
}

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func TermSize() (int, int) {
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
