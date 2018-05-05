package main

import (
	"sync"
)

type ABM struct {
	amx    sync.RWMutex
	agents []Agent

	limit int
}

func NewABM() *ABM {
	return &ABM{}
}

func (a *ABM) AddAgent(agent Agent) {
	a.amx.Lock()
	a.agents = append(a.agents, agent)
	a.amx.Unlock()
}

func (a *ABM) LimitIterations(n int) {
	a.limit = n
}

func (a *ABM) StartSimulation() <-chan report {
	ch := make(chan report)
	go func() {
		for i := 0; i < a.limit; i++ {
			a.amx.RLock()
			n := len(a.agents)
			a.amx.RUnlock()
			var wg sync.WaitGroup
			for j := 0; j < n; j++ {
				wg.Add(1)
				go func(wg *sync.WaitGroup, i, j int) {
					a.agents[j].Run(i)
					wg.Done()
				}(&wg, i, j)
			}
			wg.Wait()

			// collect data
			var r report
			r.state = make(map[string]int)
			for k := 0; k < n; k++ {
				h := a.agents[k].(*Human)
				if h.alive {
					r.alive++
				}
				r.state[h.ageRange]++
			}
			ch <- r
		}
		close(ch)
	}()
	return ch
}

type report struct {
	alive float64
	state map[string]int
}
