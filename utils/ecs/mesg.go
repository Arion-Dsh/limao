package ecs

import "sync"

var MesgBox *PubSub

func init() {
	MesgBox = &PubSub{subs: map[string][]MesgExec{}}
}

//Mesg  a message used to send
type Mesg interface {
	Topic() string
}

//MesgExec dispatch a message
type MesgExec func(m Mesg)

// PubSub manager of the message
type PubSub struct {
	mu   sync.RWMutex
	subs map[string][]MesgExec
}

//Subscribe  to the specified message type and calls the execute when fired
func (ps *PubSub) Subscribe(topic string, exec MesgExec) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ps.subs[topic] = append(ps.subs[topic], exec)
}

//PubSub Public a message
func (ps *PubSub) Public(m Mesg) {
	ps.mu.RLock()
	execs := ps.subs[m.Topic()]
	es := make([]MesgExec, len(execs))

	for i, exec := range execs {
		es[i] = exec
	}

	ps.mu.RUnlock()

	for _, exec := range es {
		exec(m)
	}
}

//UnSubcribe  #todo remove exec
func (ps *PubSub) UnSubcribe(topic string) {

}
