package main

import (
	"github.com/divan/goabm/abm"
	"github.com/divan/goabm/models/human"
	"github.com/divan/goabm/ui/term"
)

func main() {
	a := abm.New()

	for i := 0; i < 100; i++ {
		a.AddAgent(human.New(a))
	}

	a.LimitIterations(200)

	alivesCh := make(chan float64)
	a.SetReportFunc(func(a *abm.ABM) {
		alive := a.Count(func(agent abm.Agent) bool {
			h := agent.(*human.Human)
			return h.IsAlive()
		})
		alivesCh <- float64(alive)
	})

	go a.StartSimulation()

	ui := term.NewUI()
	defer ui.Stop()

	ui.AddChart("Humans Alive", alivesCh)
	ui.Loop()
}
