//go:build linux || darwin || windows
// +build linux darwin windows

package main

import (
	"image"
	"limao"
)

var i *limao.Image
var i2 *limao.Image
var i3 *limao.Image

type app struct{}

func (app *app) PreLoad() {

	limao.LoadDrawable(i, i2)
	i.Geom().Translate(409, 50)
	// i.Translate(900, 1200)
	// i.Scale(1.2, 1.2)
	// i.Rotate(90)
	i2.Geom().Translate(589, 0)
	// i2.Scale(1.5, 1.5)
	// i3.Translate(300, 0)
}

func (app *app) Update() {
}
func (app *app) Draw() {

	i.Draw(nil)
	i2.Draw(nil)
	// i3.Draw(nil)
}

func main() {
	a := new(app)

	i = limao.NewImageFromFile("gopher.png")
	i2 = i.SubImage(image.Rect(100, 100, 150, 150))
	// i3 = i2.SubImage(image.Rect(20, 20, 120, 170))
	o := &limao.RunOpts{Width: 600, Height: 800}
	limao.Run(a, o)

}
