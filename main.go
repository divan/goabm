package main

import (
	"flag"
	"fmt"
)

func main() {
	var noui = flag.Bool("noui", false, "Don't show ui, just print to console")
	flag.Parse()

	abm := NewABM()

	humans := 100

	for i := 0; i < humans; i++ {
		h := NewHuman()
		h.abm = abm
		abm.AddAgent(h)
	}

	abm.LimitIterations(1000)
	ch := abm.StartSimulation()

	if *noui {
		fmt.Println("Alive")
		for c := range ch {
			fmt.Printf("%d, ", int(c))
		}
	}
	ui := initUI(ch)
	defer stopUI()

	ui.HandleKeys()
	ui.Loop()
}
