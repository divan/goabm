package human

import (
	"github.com/atgjack/prob"
	"github.com/divan/goabm/abm"
)

var (
	childAgeDist prob.Distribution
	childrenDist prob.Distribution
	deathDist    prob.Distribution
)

func init() {
	var err error
	childAgeDist, err = prob.NewNormal(25, 7.5)
	if err != nil {
		panic(err)
	}
	childrenDist, err = prob.NewNormal(1.5, 3.0)
	if err != nil {
		panic(err)
	}
	deathDist, err = prob.NewNormal(80, 1.5)
	if err != nil {
		panic(err)
	}
}

// Human implements Agent interface for human that can replicate and age.
// Implements Agent interface.
type Human struct {
	age int

	alive    bool
	ageRange string

	childAge int
	deathAge int

	abm *abm.ABM
}

// New inits new human agent.
func New(abm *abm.ABM) abm.Agent {
	return &Human{
		age:      0,
		alive:    true,
		ageRange: "newborn",

		childAge: firstChildAge(),
		deathAge: deathAge(),

		abm: abm,
	}
}

func (h *Human) IsAlive() bool {
	return h.alive
}

func (h *Human) Age() int {
	return h.age
}

// Run satisfies Agent interface.
func (h *Human) Run() {
	h.age++
	h.updateAgeRange()

	if h.age == h.childAge {
		for i := 0; i < numChilds(); i++ {
			h.BornNewChild()
		}
	} else if h.age == h.deathAge {
		h.Die()
	}
}

func firstChildAge() int {
	return int(childAgeDist.Random())
}
func numChilds() int {
	return int(childrenDist.Random())
}
func deathAge() int {
	return int(deathDist.Random())
}

func (h *Human) updateAgeRange() {
	if !h.alive {
		h.ageRange = "dead"
		return
	}

	switch {
	case h.age < 3:
		h.ageRange = "newborn"
		break
	case h.age < 10:
		h.ageRange = "child"
		break
	case h.age < 18:
		h.ageRange = "teen"
		break
	case h.age < 50:
		h.ageRange = "adult"
		break
	case h.age < 60:
		h.ageRange = "aged"
		break
	case h.age >= 60:
		h.ageRange = "old"
		break
	default:
		h.ageRange = "N/A"
	}
}

func (h *Human) BornNewChild() {
	child := New(h.abm)
	h.abm.AddAgent(child)
}

func (h *Human) Die() {
	h.alive = false
}
