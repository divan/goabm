package main

import (
	"log"
	"math/rand"

	"github.com/divan/goabm/abm"
	"github.com/divan/goabm/models/random_walker"
)

type Prey struct {
	*walker.Walker

	abm *abm.ABM
}

func (p *Prey) Run() {
	p.Walker.Run()
}

func NewPrey(a *abm.ABM, w, h int) *Prey {
	wcell, err := walker.New(a, rand.Intn(w-1), rand.Intn(h-1), false)
	if err != nil {
		log.Fatal(err)
	}
	return &Prey{
		Walker: wcell,
		abm:    a,
	}
}

type Predator struct {
	*walker.Walker

	abm *abm.ABM
}

func (p *Predator) Run() {
	p.Walker.Run()
}

func NewPredator(a *abm.ABM, w, h int) *Predator {
	wcell, err := walker.New(a, rand.Intn(w-1), rand.Intn(h-1), false)
	if err != nil {
		log.Fatal(err)
	}
	return &Predator{
		Walker: wcell,
		abm:    a,
	}
}

type Grass struct {
	abm *abm.ABM
}

func (p *Grass) Run() {
}

func NewGrass(a *abm.ABM, w, h int) *Grass {
	return &Grass{
		abm: a,
	}
}
