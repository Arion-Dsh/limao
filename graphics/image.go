package graphics

import (
	"image"
	"limao/graphics/gl"
)

func New2Dshader(opt *ShaderOption) (Shader, error) {

	ctx := g.ctx

	fs, err := formatFs2D(opt)

	if err != nil {
		return nil, err
	}

	program, err := CreateProgram(ctx, vs2d, fs)
	if err != nil {
		return nil, err
	}

	shader := &d2Shader{
		shader: &shader{
			ctx:     g.ctx,
			program: program,
		},
		vbo:     ctx.CreateBuffer(),
		vert:    ctx.GetAttribLocation(program, "vert"),
		clrVert: ctx.GetAttribLocation(program, "clr_vert"),

		mvp:   ctx.GetUniformLocation(program, "mvp"),
		uvp:   ctx.GetUniformLocation(program, "uvp"),
		tex:   ctx.GetUniformLocation(program, "tex"),
		alpha: ctx.GetUniformLocation(program, "alpha"),
	}

	shader.vbo = ctx.CreateBuffer()
	ctx.BindBuffer(gl.ARRAY_BUFFER, shader.vbo)
	ctx.BufferData(gl.ARRAY_BUFFER, defaultVert2d, gl.STATIC_DRAW)

	return shader, nil
}

type d2Shader struct {
	*shader
	vbo     gl.Buffer
	vert    gl.Attrib
	clrVert gl.Attrib

	mvp   gl.Uniform
	uvp   gl.Uniform
	tex   gl.Uniform
	alpha gl.Uniform
}

func (s *d2Shader) NewTexture2D(w, h int) Texture {

	tex := tex2D{
		shader: s,
		ctx:    s.ctx,
		w:      w,
		h:      h,
	}

	ctx := s.ctx
	tex.gltex = ctx.CreateTexture()
	ctx.BindTexture(gl.TEXTURE_2D, tex.gltex)
	ctx.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, tex.w, tex.h, gl.RGBA, gl.UNSIGNED_BYTE, nil)
	ctx.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	ctx.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	ctx.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	ctx.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	s.active++
	return tex
}

type tex2D struct {
	shader Shader

	ctx gl.Context

	gltex gl.Texture

	w, h int
}

func (t tex2D) id() gl.Texture {
	return t.gltex
}

func (t tex2D) BindData(data []uint8) {
	ctx := t.ctx
	ctx.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	ctx.BindTexture(gl.TEXTURE_2D, t.gltex)
	ctx.TexSubImage2D(gl.TEXTURE_2D, 0, 0, 0, t.w, t.h, gl.RGBA, gl.UNSIGNED_BYTE, data)
}

func (t tex2D) Draw(opt DrawOptions, r image.Rectangle) {

	shader, _ := t.shader.(*d2Shader)

	ctx := t.ctx

	ctx.Enable(gl.BLEND)

	shader.Use()

	u := [9]float32{
		float32(t.w), float32(t.h), 0,
		float32(r.Min.X), float32(r.Min.Y), 0,
		0, 0, 0,
	}
	if opt.Color != nil {
		r, g, b, a := opt.Color.RGBA()
		u[2] = float32(1)
		u[5] = float32(r)
		u[6] = float32(g)
		u[7] = float32(b)
		u[8] = float32(a)
	}
	ctx.UniformMatrix3fv(shader.uvp, u[:])

	geom := opt.GeoM

	w := float32(r.Bounds().Dx())
	h := float32(r.Bounds().Dy())

	m := [9]float32{
		float32(g.vw), w, float32(opt.GeoM.Px()),
		float32(g.vh), h, float32(geom.Py()),
		float32(geom.Sx()), float32(geom.Sy()), float32(geom.R2D()),
	}

	ctx.UniformMatrix3fv(shader.mvp, m[:])

	ctx.Uniform1f(shader.alpha, opt.Alpha)

	for k, v := range opt.Uniforms {
		shader.SetUniform(k, v)
	}

	ctx.ActiveTexture(gl.TEXTURE0)
	ctx.BindTexture(gl.TEXTURE_2D, t.gltex)

	ctx.BindBuffer(gl.ARRAY_BUFFER, shader.vbo)

	ctx.EnableVertexAttribArray(shader.vert)
	ctx.VertexAttribPointer(shader.vert, 2, gl.FLOAT, false, 4*4, 0)

	ctx.EnableVertexAttribArray(shader.clrVert)
	ctx.VertexAttribPointer(shader.clrVert, 2, gl.FLOAT, false, 4*4, 2*4)

	ctx.DrawArrays(gl.TRIANGLE_STRIP, 0, 4)

	ctx.DisableVertexAttribArray(shader.vert)
	ctx.DisableVertexAttribArray(shader.clrVert)

	ctx.Disable(gl.BLEND)

}

func (t tex2D) release() {
	g.ctx.DeleteTexture(t.gltex)
}
