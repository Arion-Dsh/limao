//go:build js
// +build js

package main

import (
	"limao"
)

var i *limao.Image

type app struct{}

func (app *app) PreLoad() {
	limao.LoadDrawable(i)
}

func (app *app) Update() {

}
func (app *app) Draw() {
	i.Draw(nil)

}

func main() {

	i = limao.NewImageFromFile("/gopher.png")
	a := new(app)
	o := &limao.RunOpts{Width: 800, Height: 600}
	limao.Run(a, o)

}
