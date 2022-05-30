package graphics

import (
	"errors"
	"fmt"
	"limao/graphics/gl"
)

type Shader interface {
	Use()
	NewTexture2D(h, w int) Texture
}

type ShaderOption struct {
	Name      string
	Uniforms  map[string]string
	Fns       []string
	FragColor string
}

type shader struct {
	ctx     gl.Context
	program gl.Program
	active  int
}

func (s *shader) Use() {
	s.ctx.UseProgram(s.program)
}

func (s *shader) BindBuffer(name string) {

	if b, ok := g.buffers[name]; ok {
		g.ctx.BindBuffer(gl.ARRAY_BUFFER, b)
	}

}

func (s *shader) EnableAttrib(name string, size, stride, offset int) {
	loc := s.ctx.GetAttribLocation(s.program, name)
	s.ctx.EnableVertexAttribArray(loc)
	s.ctx.VertexAttribPointer(loc, size, gl.FLOAT, false, stride*4, offset*4)
}

func (s *shader) SetUniform(name string, data []float32) {

	loc := s.ctx.GetUniformLocation(s.program, name)

	if loc.Value < 0 {
		return
	}
	switch len(data) {
	case 2:
	case 4:
		// s.ctx.UniformMatrix2fv(t.uvp, u[:])
	}
}

func formatFs2D(opt *ShaderOption) (string, error) {

	var fs2d = `#version 100

precision mediump float;

varying  vec2 Tex2DCoords;
varying float fillColor;
varying vec4 color;

uniform sampler2D tex;
uniform float alpha;

	%s 

void main(){
	gl_FragColor = texture2D(tex, Tex2DCoords);
	if (fillColor == 1.){
		gl_FragColor =color; 
	}
	if (alpha != 0.) {
		gl_FragColor = vec4(gl_FragColor.xyz, gl_FragColor.w * alpha);
	}

	%s
}`

	ext := ""

	for k, v := range opt.Uniforms {
		switch v {
		case "float", "vec2", "vec3", "vec4", "mat3", "mat4":
		default:
			return "", errors.New("not suports type" + v)

		}

		ext += fmt.Sprintf(`
uniform %s %s;
		`, v, k)
	}

	for _, f := range opt.Fns {
		ext += fmt.Sprintf(`%s

`, f)
	}

	fs2d = fmt.Sprintf(fs2d, ext, opt.FragColor)

	return fs2d, nil
}
