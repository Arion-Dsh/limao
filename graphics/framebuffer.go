package graphics

import (
	"errors"
	"limao/graphics/gl"
)

type framebuffer struct {
	ctx gl.Context

	vao gl.VertexArray

	vbo          gl.Buffer
	posLoc       gl.Attrib
	texCoordsLoc gl.Attrib

	// framebuffer
	fb  gl.Framebuffer
	tex gl.Texture
	rbo gl.Renderbuffer

	program gl.Program
}

func (fb *framebuffer) id() gl.Framebuffer {

	return fb.fb
}

func newframeBuffer(ctx gl.Context, w, h int) (*framebuffer, error) {

	program, err := CreateProgram(ctx, vsSc, fsSc)
	if err != nil {
		return nil, err
	}

	fb := &framebuffer{
		ctx: ctx,

		vbo:          ctx.CreateBuffer(),
		posLoc:       ctx.GetAttribLocation(program, "pos"),
		texCoordsLoc: ctx.GetAttribLocation(program, "texCoords"),

		fb:  ctx.CreateFramebuffer(),
		tex: ctx.CreateTexture(),
		rbo: ctx.CreateRenderbuffer(),

		program: program,
	}

	ctx.Enable(gl.DEPTH_TEST)

	ctx.BindBuffer(gl.ARRAY_BUFFER, fb.vbo)
	ctx.BufferData(gl.ARRAY_BUFFER, fbData, gl.STATIC_DRAW)

	// framebuffer
	ctx.BindFramebuffer(gl.FRAMEBUFFER, fb.fb)

	// create a color attachment texture
	ctx.ActiveTexture(gl.TEXTURE0)
	ctx.BindTexture(gl.TEXTURE_2D, fb.tex)
	ctx.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, w, h, gl.RGBA, gl.UNSIGNED_BYTE, nil)
	ctx.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	ctx.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	ctx.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	ctx.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	ctx.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, fb.tex, 0)

	// create a renderbuffer object for depth and stencil attachment
	ctx.BindRenderbuffer(gl.RENDERBUFFER, fb.rbo)

	ctx.RenderbufferStorage(gl.RENDERBUFFER, gl.DEPTH24_STENCIL8, w, h)
	ctx.FramebufferRenderbuffer(gl.FRAMEBUFFER, gl.DEPTH_STENCIL_ATTACHMENT, gl.RENDERBUFFER, fb.rbo) // now actually attach it
	// now that we actually created the framebuffer and added all attachments we want to check if it is actually complete now

	if ctx.CheckFramebufferStatus(gl.FRAMEBUFFER) != gl.FRAMEBUFFER_COMPLETE {
		return nil, errors.New("framebuffer is not complete")
	}

	return fb, nil
}

func (fb *framebuffer) resize(w, h int) {

	fb.ctx.BindTexture(gl.TEXTURE_2D, fb.tex)
	fb.ctx.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int(w), int(h), gl.RGBA, gl.UNSIGNED_BYTE, nil)
	fb.ctx.BindRenderbuffer(gl.RENDERBUFFER, fb.rbo)
	fb.ctx.RenderbufferStorage(gl.RENDERBUFFER, gl.DEPTH24_STENCIL8, w, h)
}

func (fb *framebuffer) draw() {
	ctx := fb.ctx
	ctx.BindFramebuffer(gl.FRAMEBUFFER, gl.Framebuffer{})

	// clear all relevant buffers
	ctx.Disable(gl.DEPTH_TEST)

	ctx.Enable(gl.BLEND)

	ctx.UseProgram(fb.program)

	ctx.ActiveTexture(gl.TEXTURE0)
	ctx.BindTexture(gl.TEXTURE_2D, fb.tex)

	ctx.BindBuffer(gl.ARRAY_BUFFER, fb.vbo)
	ctx.EnableVertexAttribArray(fb.posLoc)
	ctx.VertexAttribPointer(fb.posLoc, 2, gl.FLOAT, false, 4*4, 0)

	ctx.EnableVertexAttribArray(fb.texCoordsLoc)
	ctx.VertexAttribPointer(fb.texCoordsLoc, 2, gl.FLOAT, false, 4*4, 2*4)

	ctx.DrawArrays(gl.TRIANGLE_STRIP, 0, 4)

	ctx.DisableVertexAttribArray(fb.posLoc)
	ctx.DisableVertexAttribArray(fb.texCoordsLoc)

	ctx.Disable(gl.BLEND)

	ctx.BindFramebuffer(gl.FRAMEBUFFER, fb.fb)

}

func (fb *framebuffer) release() {

	ctx := g.ctx
	ctx.DeleteVertexArray(fb.vao)
	ctx.DeleteBuffer(fb.vbo)
	ctx.DeleteFramebuffer(fb.fb)
	ctx.DeleteTexture(fb.tex)
	ctx.DeleteRenderbuffer(fb.rbo)
}

var fbData = f32Bytes(
	-1, +1, 0, 1, // top left
	+1, +1, 1, 1, // top right
	-1, -1, 0, 0, // bottom left
	+1, -1, 1, 0, // bottom right
)

var vsSc = `#version 100

attribute  vec2 pos;
attribute vec2 texCoords;

varying vec2 clr_TexCoords;

void main()
{
    clr_TexCoords = texCoords;
	gl_Position = vec4(pos, 0, 1.0);

}

`

var fsSc = `#version 100

precision mediump float;
varying  vec2 clr_TexCoords;

uniform sampler2D screenTexture;

void main()
{
	gl_FragColor = texture2D(screenTexture, clr_TexCoords);
	
} 
`
