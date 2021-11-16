package geom

type Size struct {
	W, H int
}

type Geom struct {
	Position Point
	Scale    Scale
	Rotate2d float32
}

func New() *Geom {
	return &Geom{
		Position: Point{X: 0, Y: 0},
		Scale:    Scale{X: 1, Y: 1},
		Rotate2d: 0,
	}
}

func (g *Geom) Identity() {

}

func (g *Geom) Add(sub *Geom) *Geom {
	if sub == nil {
		return g
	}
	g.Position.X += sub.Position.X
	g.Position.Y += sub.Position.Y
	g.Scale.X *= sub.Scale.X
	g.Scale.Y *= sub.Scale.Y
	g.Rotate2d += sub.Rotate2d
	return g
}

func (g *Geom) Translate(x, y int) *Geom {
	g.Position.X += x
	g.Position.Y += y
	return g
}

func (g *Geom) Scale2D(x, y float32) *Geom {
	g.Scale.X *= x
	g.Scale.Y *= y
	return g
}

func (g *Geom) Rotate2D(r float32) *Geom {
	g.Rotate2d += r
	return g
}

// Point the point on screen
type Point struct {
	X, Y, Z int
}

// Rectangle ...
type Rectangle struct {
	Min, Max Point
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
