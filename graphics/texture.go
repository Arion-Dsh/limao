package graphics

import (
	"image"
	"limao/graphics/gl"
)

type Texture interface {
	id() gl.Texture
	BindData([]uint8)
	Draw(opt DrawOptions, r image.Rectangle)
	release()
}

type texture2Dv3 struct {
	ctx    gl.Context
	shader *shader
	gltex  gl.Texture
	w, h   int
}

func (t *texture2Dv3) BindPixel(data []uint8) {
	t.ctx.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	t.ctx.BindTexture(gl.TEXTURE_2D, t.gltex)
	t.ctx.TexSubImage2D(gl.TEXTURE_2D, 0, 0, 0, t.w, t.h, gl.RGBA, gl.UNSIGNED_BYTE, data)
}

func CreateTexture2D(w, h int) Texture {
	return g.imageShader.NewTexture2D(w, h)
}
