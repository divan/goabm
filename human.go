package main

import (
	"github.com/atgjack/prob"
)

var (
	childAgeDist prob.Distribution
	childrenDist prob.Distribution
	deathDist    prob.Distribution
)

func init() {
	var err error
	childAgeDist, err = prob.NewNormal(30, 3.5)
	if err != nil {
		panic(err)
	}
	childrenDist, err = prob.NewNormal(1.1, 1.0)
	if err != nil {
		panic(err)
	}
	deathDist, err = prob.NewNormal(80, 1.5)
	if err != nil {
		panic(err)
	}
}

// Human implements Agent interface for human that can replicate and age.
type Human struct {
	age int

	alive    bool
	ageRange string

	abm *ABM
}

func NewHuman() *Human {
	return &Human{
		age:      0,
		alive:    true,
		ageRange: "newborn",
	}
}

// Run satisfies Agent interface.
//
// i effectively means years.
func (h *Human) Run(i int) {
	h.age++
	h.updateAgeRange()

	rndAge := int(childAgeDist.Random())
	if h.age == rndAge {
		rndChildren := int(childrenDist.Random())
		for i := 0; i < rndChildren; i++ {
			h.BornNewChild()
		}
	}

	rndDeathAge := int(deathDist.Random())
	if h.age == rndDeathAge {
		h.Die()
	}
}

func (h *Human) updateAgeRange() {
	if !h.alive {
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
	case h.age < 30:
		h.ageRange = "young"
		break
	case h.age < 45:
		h.ageRange = "middle-age"
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
	child := NewHuman()
	child.abm = h.abm
	h.abm.AddAgent(child)
}

func (h *Human) Die() {
	h.alive = false
}
