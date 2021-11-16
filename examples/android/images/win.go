package main

import (
	"limao"
)

var i *limao.Image
var i2 *limao.Image

type app struct{}

func (app *app) PreLoad() {
	limao.LoadDrawable(i, i2)
}

func (app *app) Update() {
	// fmt.Println("111")
}
func (app *app) Draw() {
	i.Draw(nil)
	i2.Draw(nil)
}

func main() {
	a := new(app)

	i = limao.NewImageFromFile("gopher.png")
	i2 = limao.NewImageFromFile("gopher.png")
	o := &limao.RunOpts{Width: 600, Height: 800}
	limao.Run(a, o)

}
