//go:build linux || darwin || windows
// +build linux darwin windows

package graphics

import (
	"image"
	"limao/geom"
	"limao/graphics/gl"
)

func NewTexture2D(w, h int) Texture {
	if g.images == nil {
		g.images = newImages2D()
	}

	return g.images.newTexture2D(w, h)

}

func DrawTexture2D(t2d Texture, geom geom.Geom, r image.Rectangle) {

	t, _ := t2d.(*texture2D)

	if !t.verify(geom, r) {
		return
	}

	ctx := g.ctx

	ctx.Enable(gl.BLEND)

	// ctx.BindFramebuffer(gl.FRAMEBUFFER, g.fb.id())

	ctx.UseProgram(t.program)

	t.bindData(ctx, g.vw, g.vh, geom, r)

	ctx.BindVertexArray(t.vao)

	ctx.ActiveTexture(gl.TEXTURE0)
	ctx.BindTexture(gl.TEXTURE_2D, t.gltex)

	ctx.DrawArrays(gl.TRIANGLE_STRIP, 0, 4)

	ctx.Disable(gl.BLEND)

}
