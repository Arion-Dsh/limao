package input

type Key uint32

type key struct {
	code  Key
	state int
}

func KeyPress(k Key) {
	i.mu.RLock()
	defer i.mu.RUnlock()
	i.keys[k] = key{k, 1}
}

func KeyRelease(k Key) {
	i.mu.RLock()
	defer i.mu.RUnlock()
	i.keys[k] = key{k, 2}
}

func IsKeyRelease(k Key) bool {
	i.mu.RLock()

	defer i.mu.RUnlock()
	ky, ok := i.keys[k]
	if ok && ky.state == 2 {
		delete(i.keys, k)
		return true
	}
	return false

}

func IsKeyPressed(k Key) bool {
	i.mu.Lock()
	defer i.mu.Unlock()
	ky, ok := i.keys[k]
	if ok && ky.state == 1 {
		// delete(i.keys, k)
		return true
	}
	return false
}
