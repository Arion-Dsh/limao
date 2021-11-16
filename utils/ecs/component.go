package ecs

var ComponentsBox *components

func init() {
	ComponentsBox = &components{}
}

type components struct {
	comps []*component
}

func (cps *components) make(e *Entity, c Component) {
	if cps.has(e, c.Label()) {
		return
	}

	cps.comps = append(cps.comps, &component{e, c})
}

func (cps *components) has(e *Entity, label string) bool {
	ok := false

	for _, component := range cps.GetByEntity(e) {
		if component.Label() == label {
			ok = true
		}
	}

	return ok
}

func (cps *components) GetByLabel(label string) []Component {
	components := []Component{}
	for _, comp := range cps.comps {
		if comp.C.Label() == label {
			components = append(components, comp.C)
		}
	}

	return components

}

func (cps *components) GetByEntity(e *Entity) []Component {
	components := []Component{}

	for _, comp := range cps.comps {
		if e == comp.E {
			components = append(components, comp.C)
		}
	}

	return components
}

type component struct {
	E *Entity
	C Component
}

// Component ...
type Component interface {
	Label() string
}
