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

	alivesCh, infantsCh, retiredCh := make(chan float64), make(chan float64), make(chan float64)
	a.SetReportFunc(func(a *abm.ABM) {
		alive := a.Count(func(agent abm.Agent) bool {
			h := agent.(*human.Human)
			return h.IsAlive()
		})
		infants := a.Count(func(agent abm.Agent) bool {
			h := agent.(*human.Human)
			return h.Age() < 5
		})
		retired := a.Count(func(agent abm.Agent) bool {
			h := agent.(*human.Human)
			return h.Age() > 60
		})
		alivesCh <- float64(alive)
		infantsCh <- float64(infants)
		retiredCh <- float64(retired)
	})

	go a.StartSimulation()

	ui := term.NewUI()
	defer ui.Stop()

	ui.AddChart("Humans Alive", alivesCh)
	ui.AddChart("Infants", infantsCh)
	ui.AddChart("Retired", retiredCh)
	ui.Loop()
}
