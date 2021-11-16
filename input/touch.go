package input

type touch struct {
	id     int
	ox, oy int // begin x, y
	x, y   int
}

func TouchMove(id, x, y int) {

	t, ok := i.touchs[id]
	if !ok {
		t = &touch{id: id, ox: x, oy: y, x: x, y: y}
		i.touchs[id] = t

		return
	}
	t.x, t.y = x, y
}

func TouchEnd(id int) {

	_, ok := i.touchs[id]
	if !ok {
		return
	}
	delete(i.touchs, id)
}
