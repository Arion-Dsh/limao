package graphics

import (
	"image"
	"limao/geom"
	"limao/graphics/gl"
	"math"
)

type texture2D struct {
	ctx     gl.Context
	vao     gl.VertexArray
	program gl.Program
	mvp     gl.Uniform
	uvp     gl.Uniform

	vert       gl.Buffer
	vertLoc    gl.Attrib
	clrVertLoc gl.Attrib

	gltex gl.Texture
	w, h  int
}

func (t *texture2D) BindData(data []uint8) {
	t.ctx.BlendFunc(gl.ONE, gl.ONE_MINUS_SRC_ALPHA)
	t.ctx.BindTexture(gl.TEXTURE_2D, t.gltex)
	t.ctx.TexSubImage2D(gl.TEXTURE_2D, 0, 0, 0, t.w, t.h, gl.RGBA, gl.UNSIGNED_BYTE, data)

}

func (t *texture2D) verify(geom geom.Geom, rect image.Rectangle) bool {

	rect = image.Rect(0, 0, rect.Dx(), rect.Dy())

	max := image.Point{
		X: rect.Max.X + rect.Dx()*int(math.Round(float64(geom.Scale.X))),
		Y: rect.Max.Y + rect.Dy()*int(math.Round(float64(geom.Scale.Y))),
	}
	rect = image.Rectangle{rect.Min, max}
	rect = rect.Add(image.Pt(geom.Position.X, geom.Position.Y))
	b := image.Rect(0, 0, g.vw, g.vh)

	return b.Overlaps(rect)

}

func (t *texture2D) bindData(ctx gl.Context, vw, vh int, geom geom.Geom, r image.Rectangle) {

	u := [4]float32{
		float32(t.w), float32(t.h),
		float32(r.Min.X), float32(r.Min.Y),
	}
	ctx.UniformMatrix2fv(t.uvp, u[:])

	w := float32(r.Bounds().Dx())
	h := float32(r.Bounds().Dy())

	m := [9]float32{
		float32(vw), w, float32(geom.Position.X),
		float32(vh), h, float32(geom.Position.Y),
		float32(geom.Scale.X), float32(geom.Scale.Y), 0,
	}
	ctx.UniformMatrix3fv(t.mvp, m[:])

}

func (tex *texture2D) release() {
	g.ctx.DeleteTexture(tex.gltex)

}
