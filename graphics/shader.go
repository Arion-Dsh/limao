package graphics

import (
	"fmt"

	"limao/graphics/gl"
)

// CreateProgram creates, compiles, and links a gl.Program.
func CreateProgram(glctx gl.Context, vertexSrc, fragmentSrc string) (gl.Program, error) {

	program := glctx.CreateProgram()
	vertexShader, err := loadShader(glctx, gl.VERTEX_SHADER, vertexSrc)
	if err != nil {
		return gl.Program{}, err
	}
	fragmentShader, err := loadShader(glctx, gl.FRAGMENT_SHADER, fragmentSrc)
	if err != nil {
		return gl.Program{}, err
	}
	glctx.AttachShader(program, vertexShader)
	glctx.AttachShader(program, fragmentShader)
	glctx.LinkProgram(program)

	// Flag shaders for deletion when program is unlinked.
	glctx.DeleteShader(vertexShader)
	glctx.DeleteShader(fragmentShader)

	if glctx.GetProgrami(program, gl.LINK_STATUS) == 0 {
		defer glctx.DeleteProgram(program)
		return gl.Program{}, fmt.Errorf("shader complie: %s", glctx.GetProgramInfoLog(program))
	}

	return program, nil
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
