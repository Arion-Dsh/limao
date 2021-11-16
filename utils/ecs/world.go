package ecs

import (
	"errors"
	"sort"
)

// World is collection of enity and system
type World struct {
	systems  []System
	needSort bool
}

// AddSystem and the system to the word
func (w *World) AddSystem(s System) error {
	if s == nil {
		return errors.New("s can't be nil")
	}
	ok := false
	for _, sys := range w.systems {
		if sys.String() == s.String() {
			ok = true

		}
	}
	if !ok {
		w.systems = append(w.systems, s)
		w.needSort = true
	}
	return nil
}

//GetSystem ...
func (w *World) GetSystem(str string) (System, error) {
	for _, s := range w.systems {
		if s.String() == str {
			return s, nil
		}
	}
	return nil, errors.New("no this system")
}

//System return the world's systems
func (w *World) System() []System {
	return w.systems
}

// Update run in every frame
func (w *World) Update(dt float64) {
	if w.needSort {
		sort.Sort(sort.Reverse(systemSort(w.systems)))
		w.needSort = false
	}

	for _, system := range w.systems {
		system.Update(dt)
	}

}
