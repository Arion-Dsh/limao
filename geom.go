package limao

type Geom struct {
	px, py   int
	sx, sy   float32
	rotate2D int
}

func (g *Geom) Px() int {
	return g.px
}
func (g *Geom) Py() int {
	return g.py
}

func (g *Geom) Sx() float32 {
	if g.sx == 0 {
		return 1
	}
	return g.sx
}

func (g *Geom) Sy() float32 {
	if g.sy == 0 {
		return 1
	}
	return g.sy
}
func (g *Geom) R2D() int {
	return g.rotate2D
}

func (g *Geom) Identity() {
	g.sx, g.sy = 0, 0
	g.rotate2D = 0
}

func (g *Geom) Add(sub *Geom) *Geom {
	return g
}

func (g *Geom) Translate(x, y int) *Geom {
	g.px += x
	g.py += y
	return g
}

func (g *Geom) Scale2D(x, y float32) *Geom {
	if g.sx == 0 {
		g.sx = 1
	}
	if g.sy == 0 {
		g.sy = 1
	}
	g.sx *= x
	g.sy *= y
	return g
}

func (g *Geom) Rotate2D(r int) *Geom {
	g.rotate2D += r
	return g
}

type Scale struct {
	X, Y float32
}

func (s Scale) IsZore() bool {
	if s.X <= 0 || s.Y <= 0 {
		return true
	}
	return false
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
