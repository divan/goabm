package abm

type World interface {
	Tick() // mark the beginning of the next time period
}

// BorderRule represents rule of what happens when agent
// reaches the edge of the world.
type BorderRule uint8

const (
	BorderPeriodic  BorderRule = iota // wrap around the other side
	BorderAPeriodic                   // hit the wall
)
