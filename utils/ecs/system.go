package ecs

//System is a tool  for working with a collection of entities that have one or
//more of the same components
type System interface {
	Update(dt float64)
	String() string
	Priority() int
}

type systemSort []System

func (l systemSort) Len() int { return len(l) }

func (l systemSort) Less(i, j int) bool {
	return l[i].Priority() < l[j].Priority()
}

func (l systemSort) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
