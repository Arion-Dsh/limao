package graphics

import (
	"limao/graphics/gl"
	"sync"
)

type images2Dv3 struct {
	ctx gl.Context

	vao     gl.VertexArray
	program gl.Program
	vert    gl.Buffer

	vertLoc    gl.Attrib
	clrVertLoc gl.Attrib

	uvpLoc       gl.Uniform
	transformLoc gl.Uniform
	texLoc       gl.Uniform
	mvp          gl.Uniform
	uvp          gl.Uniform

	active int
	mu     sync.Mutex

	texs map[Texture]struct{}
}

func newImages2D() images {

	withVao := false
	if g.ctx.Version() >= 3.0 {
		withVao = true
	}

	i := &images2Dv3{}
	ctx := g.ctx

	program, err := CreateProgram(ctx, vs2d, fs2d)
	if err != nil {
		panic(err)
	}
	i.program = program

	i.vert = ctx.CreateBuffer()

	i.vertLoc = ctx.GetAttribLocation(program, "vert")
	i.clrVertLoc = ctx.GetAttribLocation(program, "clr_vert")

	i.uvpLoc = ctx.GetUniformLocation(program, "uvp")
	i.transformLoc = ctx.GetUniformLocation(program, "transform")
	i.texLoc = ctx.GetUniformLocation(program, "tex")

	i.mvp = ctx.GetUniformLocation(program, "mvp")
	i.uvp = ctx.GetUniformLocation(program, "uvp")

	ctx.Enable(gl.BLEND)
	ctx.DepthFunc(gl.LESS)

	ctx.UseProgram(program)

	if withVao {
		i.vao = ctx.CreateVertexArray()
		ctx.BindVertexArray(i.vao)
	}

	ctx.BindBuffer(gl.ARRAY_BUFFER, i.vert)
	ctx.BufferData(gl.ARRAY_BUFFER, vert, gl.STATIC_DRAW)

	if withVao {

		ctx.BindBuffer(gl.ARRAY_BUFFER, i.vert)
		ctx.EnableVertexAttribArray(i.vertLoc)
		ctx.VertexAttribPointer(i.vertLoc, 2, gl.FLOAT, false, 4*4, 0)

		ctx.BindBuffer(gl.ARRAY_BUFFER, i.vert)
		ctx.EnableVertexAttribArray(i.clrVertLoc)
		ctx.VertexAttribPointer(i.clrVertLoc, 2, gl.FLOAT, false, 4*4, 2*4)

	}

	return i
}

func (imgs *images2Dv3) newTexture2D(w, h int) Texture {

	tex := &texture2D{
		ctx:        g.ctx,
		program:    imgs.program,
		vao:        imgs.vao,
		mvp:        imgs.mvp,
		uvp:        imgs.uvp,
		vert:       imgs.vert,
		vertLoc:    imgs.vertLoc,
		clrVertLoc: imgs.clrVertLoc,

		w: w,
		h: h,
	}

	ctx := g.ctx
	tex.gltex = ctx.CreateTexture()
	ctx.BindTexture(gl.TEXTURE_2D, tex.gltex)
	ctx.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, tex.w, tex.h, gl.RGBA, gl.UNSIGNED_BYTE, nil)
	ctx.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	ctx.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	ctx.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	ctx.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	return tex

}

func (i *images2Dv3) release() {
	for t := range i.texs {
		t.release()
	}

	ctx := g.ctx
	ctx.DeleteBuffer(i.vert)
	ctx.DeleteVertexArray(i.vao)
	ctx.DeleteProgram(i.program)
}

var vert = f32Bytes(
	0, 0, 0, 0, // top left
	+1, 0, 1, 0, // top right
	0, -1, 0, 1, // bottom left
	+1, -1, 1, 1, // bottom right
)
