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
	// i.Geom().Translate(409, 50)
	// i.Translate(900, 1200)
	// i.Scale(1.2, 1.2)
	// i.Rotate(90)
	// i2.Geom().Translate(589, 0)
	// i2.Scale(1.5, 1.5)
	// i3.Translate(300, 0)
}

func (app *app) Update() {
}
func (app *app) Draw() {

	// i.Fill(color.Black)
	i2.Draw(nil)
	op := &limao.DrawOptions{}
	op.GeoM.Translate(200, 300)
	i.Draw(op)

	op.GeoM.Translate(400, 0)

	i.Draw(op)

	// op := geom.New().Translate(240, 240)
	// i.Draw(op)
	// op2 := geom.New().Translate(0, 240)
	// i2.Draw(op2)
	// i3.Draw(nil)
}

func main() {
	a := new(app)

	i = limao.NewImageFromFile("gopher.png")
	i2 = i.SubImage(image.Rect(100, 100, 150, 150))
	// i3 = i2.SubImage(image.Rect(20, 20, 120, 170))
	o := &limao.RunOpts{Width: 800, Height: 800}
	limao.Run(a, o)

}
