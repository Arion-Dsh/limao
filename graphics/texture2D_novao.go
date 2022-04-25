//go:build android || ios || js || itest
// +build android ios js itest

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

	ctx.UseProgram(t.program)

	t.bindData(ctx, g.vw, g.vh, geom, r)

	ctx.ActiveTexture(gl.TEXTURE0)
	ctx.BindTexture(gl.TEXTURE_2D, t.gltex)

	ctx.BindBuffer(gl.ARRAY_BUFFER, t.vert)
	ctx.EnableVertexAttribArray(t.vertLoc)
	ctx.VertexAttribPointer(t.vertLoc, 2, gl.FLOAT, false, 4*4, 0)

	ctx.BindBuffer(gl.ARRAY_BUFFER, t.vert)
	ctx.EnableVertexAttribArray(t.clrVertLoc)
	ctx.VertexAttribPointer(t.clrVertLoc, 2, gl.FLOAT, false, 4*4, 2*4)

	ctx.DrawArrays(gl.TRIANGLE_STRIP, 0, 4)

	ctx.DisableVertexAttribArray(t.vertLoc)
	ctx.DisableVertexAttribArray(t.clrVertLoc)

	ctx.Disable(gl.BLEND)

}
