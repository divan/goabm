package main

import (
	"github.com/divan/goabm/abm"
	"github.com/divan/goabm/models/human"
	"github.com/divan/goabm/ui/term"
)

func main() {
	a := abm.New()

	a.AddAgents(human.New, 100)

	a.LimitIterations(1000)

	alivesCh := make(chan int)
	a.SetReportFunc(func(a *abm.ABM) {
		alive := a.Count(func(agent abm.Agent) bool {
			h := agent.(*human.Human)
			return h.IsAlive()
		})
		alivesCh <- alive
	})

	go a.StartSimulation()

	ui := term.NewUI(alivesCh)
	defer term.StopUI()

	ui.HandleKeys()
	ui.Loop()
}
