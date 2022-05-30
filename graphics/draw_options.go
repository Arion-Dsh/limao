package graphics

import (
	"image/color"
)

type geoM interface {
	Px() int
	Py() int
	Sx() float32
	Sy() float32
	R2D() int
}
type DrawOptions struct {
	GeoM     geoM
	Uniforms map[string][]float32
	Color    color.Color
	Alpha    float32
}
