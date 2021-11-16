//go:build linux || darwin || windows
// +build linux darwin windows

package main

import (
	"image"
	"image/color"
	"io/ioutil"
	"limao"
	"limao/geom"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

var i *limao.Image
var i2 *limao.Image
var i3 *limao.Image
var tf font.Face

type app struct{}

func (app *app) PreLoad() {

	limao.LoadDrawable(i, i2)
	i.Geom().Translate(120, 120)
	// i.Translate(900, 1200)
	// i.Scale(1.2, 1.2)
	// i.Rotate(90)
	// i2.Translate(40, 40)
	// i2.Scale(1.5, 1.5)
	// i3.Translate(300, 0)
}

func (app *app) Update() {
}
func (app *app) Draw() {

	i.Draw(nil)
	i2.Draw(nil)
	i3.Draw(nil)
	col := color.RGBA{200, 100, 0, 255}
	limao.DrawText(geom.New(), "test", col, tf)

}

func main() {
	a := new(app)

	ttfData, err := ioutil.ReadFile("/usr/share/fonts/truetype/noto/NotoMono-Regular.ttf")

	f, err := truetype.Parse(ttfData)
	if err != nil {
	}
	tf = truetype.NewFace(f, &truetype.Options{
		Size: 80,
		DPI:  72,
	})
	i = limao.NewImageFromFile("gopher.png")
	i2 = i.SubImage(image.Rect(90, 90, 200, 240))
	i3 = i2.SubImage(image.Rect(20, 20, 120, 170))
	o := &limao.RunOpts{Width: 600, Height: 800}
	limao.Run(a, o)

}
