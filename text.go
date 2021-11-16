package limao

import (
	"image"
	"image/color"
	"limao/geom"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type text struct {
	t rune
	d font.Drawer
}

var textCache map[text]*Image = map[text]*Image{}

func DrawText(opts *geom.Geom, texts string, clr color.Color, face font.Face) {

	fonth := face.Metrics().Height
	var dx, dy fixed.Int26_6

	for _, t := range texts {
		if t == '\n' {
			dx = 0
			dy += fonth
			continue
		}

		fontw := font.MeasureString(face, string(t))
		w := fontw.Ceil()
		h := fonth.Ceil()
		dot := fixed.Point26_6{
			X: 0,
			Y: face.Metrics().Ascent,
		}

		d := font.Drawer{
			Src:  image.NewUniform(clr),
			Face: face,
			Dot:  dot,
		}
		textImg, ok := textCache[text{t, d}]
		if !ok {
			rgba := image.NewRGBA(image.Rect(0, 0, w, h))
			textImg = newImage(rgba, w, h)
			textCache[text{t, d}] = textImg
			d.Dst = textImg.RGBA
			d.DrawString(string(t))
			textImg.Load()

		}
		// geom.Translate(x+dx.Ceil(), y+dy.Ceil())
		opt := geom.New().Translate(dx.Ceil(), dy.Ceil()).Add(opts)
		textImg.Draw(opt)
		dx += fontw

	}

}
