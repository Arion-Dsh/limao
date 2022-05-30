package graphics

import "limao/graphics/gl"

func ArrayBuffer(name string, data []float32) {
	b := g.ctx.CreateBuffer()
	g.ctx.BindBuffer(gl.ARRAY_BUFFER, b)
	g.ctx.BufferData(gl.ARRAY_BUFFER, f32Bytes(data...), gl.STATIC_DRAW)
	g.buffers[name] = b
}

/* func CreateTexture2D(w, h int) Texture { */

/* tex := tex2D{ */
/* w: w, */
/* h: h, */
/* } */

/* ctx := g.ctx */
/* tex.gltex = ctx.CreateTexture() */

/* ctx.BindTexture(gl.TEXTURE_2D, tex.gltex) */
/* ctx.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, tex.w, tex.h, gl.RGBA, gl.UNSIGNED_BYTE, nil) */
/* ctx.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR) */
/* ctx.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR) */
/* ctx.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE) */
/* ctx.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE) */

/* return tex */
/* } */

type glShader struct {
	glctx   gl.Context
	program gl.Program

	active int
}

func BindBuffer(name string) {

	if b, ok := g.buffers[name]; ok {
		g.ctx.BindBuffer(gl.ARRAY_BUFFER, b)
	}

}
