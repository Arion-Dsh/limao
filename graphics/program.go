package graphics

import (
	"fmt"

	"limao/graphics/gl"
)

/* func (p program) EnableAttrib(name string, size, stride, offset int) { */
/* loc := g.ctx.GetAttribLocation(p.id, name) */
/* g.ctx.EnableVertexAttribArray(loc) */
/* g.ctx.VertexAttribPointer(loc, size, gl.FLOAT, false, stride*4, offset*4) */
/* } */

// func (p program) SetUniform(name string, data []float32) {

// }

// CreateProgram creates, compiles, and links a gl.Program.
func CreateProgram(glctx gl.Context, vertexSrc, fragmentSrc string) (gl.Program, error) {

	p := glctx.CreateProgram()
	vertexShader, err := loadShader(glctx, gl.VERTEX_SHADER, vertexSrc)
	if err != nil {
		return p, err
	}
	fragmentShader, err := loadShader(glctx, gl.FRAGMENT_SHADER, fragmentSrc)
	if err != nil {
		return p, err
	}
	glctx.AttachShader(p, vertexShader)
	glctx.AttachShader(p, fragmentShader)
	glctx.LinkProgram(p)

	// Flag shaders for deletion when program is unlinked.
	glctx.DeleteShader(vertexShader)
	glctx.DeleteShader(fragmentShader)

	if glctx.GetProgrami(p, gl.LINK_STATUS) == 0 {
		defer glctx.DeleteProgram(p)
		return p, fmt.Errorf("shader complie: %s", glctx.GetProgramInfoLog(p))
	}

	return p, nil
}

func loadShader(glctx gl.Context, shaderType gl.Enum, src string) (gl.Shader, error) {
	shader := glctx.CreateShader(shaderType)
	glctx.ShaderSource(shader, src)
	glctx.CompileShader(shader)
	if glctx.GetShaderi(shader, gl.COMPILE_STATUS) == 0 {
		defer glctx.DeleteShader(shader)
		return gl.Shader{}, fmt.Errorf("shader compile: %s", glctx.GetShaderInfoLog(shader))
	}
	return shader, nil
}
