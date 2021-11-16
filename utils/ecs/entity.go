package ecs

import (
	"errors"
	"sync/atomic"
)

var entityIndex int64

func init() {
	entityIndex = 0
}

// Entity just a index
type Entity uint64

//NewEntity return a entity with id
func NewEntity() *Entity {
	e := Entity(atomic.AddInt64(&entityIndex, 1))
	return &e
}

func (e *Entity) GetComponent(label string) (Component, error) {
	comps := ComponentsBox.GetByEntity(e)

	for _, cp := range comps {
		if cp.Label() == label {
			return cp, nil
		}
	}
	return nil, errors.New("no this Component")
}

//GetComponents ...
func (e *Entity) GetComponents() []Component {
	return ComponentsBox.GetByEntity(e)
}

func (e *Entity) AddComponents(comps ...Component) {
	for _, c := range comps {
		ComponentsBox.make(e, c)
	}

}
