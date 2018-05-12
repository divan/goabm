package termgrid

import "fmt"

type UI struct {
	Width, Height int
	ch            chan [][]bool
}

func New(w, h int, ch chan [][]bool) *UI {
	return &UI{
		Width:  w,
		Height: h,
		ch:     ch,
	}
}

func (*UI) Stop() {
}

func (ui *UI) Loop() {
	for dump := range ui.ch {
		for i := 0; i < ui.Height; i++ {
			fmt.Printf("\033[0;0H")
		}
		for i := 0; i < ui.Height; i++ {
			for j := 0; j < ui.Width; j++ {
				if dump[i][j] {
					fmt.Printf("*")
				} else {
					fmt.Printf(" ")
				}
			}
			fmt.Println()
		}
	}
}
