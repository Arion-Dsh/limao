package input

import (
	"sync"
)

var i *input

func init() {

	i = &input{
		cursorX: 0,
		cursorY: 0,
		keys:    map[Key]key{},
		mouse:   map[MouseButton]bool{},
		touchs:  map[int]*touch{},
		mu:      sync.RWMutex{},
	}
}

type input struct {
	Rune    []rune
	cursorX int
	cursorY int

	keys  map[Key]key
	mouse map[MouseButton]bool

	touchs map[int]*touch

	mu sync.RWMutex
}

func SetCursor(x, y int) {
	i.mu.RLock()
	defer i.mu.RUnlock()
	i.cursorX = x
	i.cursorY = y
}

func SetRuns(r []rune) {
	i.mu.Lock()
	defer i.mu.Unlock()
	copy(i.Rune, r)

}
