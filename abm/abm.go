package abm

import (
	"sync"
)

type ABM struct {
	mx     sync.RWMutex
	agents []Agent

	i     int // current iteration
	limit int

	world      World
	reportFunc func(*ABM)
}

// New creates new ABM simulation engine with default
// parameters.
func New() *ABM {
	return &ABM{
		limit: 1000,
	}
}

func (a *ABM) SetWorld(w World) {
	if a.world == nil {
		a.world = w
	}
}

func (a *ABM) World() World {
	return a.world
}

func (a *ABM) SetReportFunc(fn func(*ABM)) {
	a.reportFunc = fn
}

func (a *ABM) AddAgent(agent Agent) {
	a.mx.Lock()
	a.agents = append(a.agents, agent)
	a.mx.Unlock()
}

func (a *ABM) AddAgents(spawnFunc func(*ABM) Agent, n int) {
	for i := 0; i < n; i++ {
		agent := spawnFunc(a)
		a.AddAgent(agent)
	}
}

func (a *ABM) LimitIterations(n int) {
	a.mx.Lock()
	a.limit = n
	a.mx.Unlock()
}

func (a *ABM) Limit() int {
	a.mx.RLock()
	defer a.mx.RUnlock()
	return a.limit
}

// Iteration returns current iteration (age, generation).
func (a *ABM) Iteration() int {
	a.mx.RLock()
	defer a.mx.RUnlock()
	return a.i
}

func (a *ABM) StartSimulation() {
	for i := 0; i < a.Limit(); i++ {
		a.i = i
		if a.World() != nil {
			a.World().Tick()
		}

		var wg sync.WaitGroup
		for j := 0; j < a.AgentsCount(); j++ {
			wg.Add(1)
			go func(wg *sync.WaitGroup, i, j int) {
				a.agents[j].Run()
				wg.Done()
			}(&wg, i, j)
		}
		wg.Wait()

		if a.reportFunc != nil {
			a.reportFunc(a)
		}
	}
}

func (a *ABM) AgentsCount() int {
	a.mx.RLock()
	defer a.mx.RUnlock()
	return len(a.agents)
}

func (a *ABM) Agents() []Agent {
	a.mx.RLock()
	defer a.mx.RUnlock()
	return a.agents
}

func (a *ABM) Count(condition func(agent Agent) bool) int {
	var count int
	for _, agent := range a.Agents() {
		if condition(agent) {
			count++
		}
	}
	return count
}
