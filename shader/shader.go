package shader

import (
	"errors"
	"fmt"
	"limao/graphics"
	"limao/graphics/gl"
)

type ShaderOption struct {
	Uniforms  map[string]string
	Fns       []string
	FragColor string
}

type Shader interface{}

func New2Dshader(ctx gl.Context, opt *ShaderOption) (Shader, error) {

	useVao := false
	if ctx.Version() >= 3.0 {
		useVao = true
	}

	shader := &shader{useVao: useVao}

	fs, err := formatfs(opt)

	if err != nil {
		return nil, err
	}
	program, err := graphics.CreateProgram(ctx, vs2d, fs)
	if err != nil {
		return nil, err
	}
	shader.program = program

	shader.vert = ctx.CreateBuffer()

	shader.vertLoc = ctx.GetAttribLocation(program, "vert")
	shader.clrVertLoc = ctx.GetAttribLocation(program, "clr_vert")

	shader.mvp = ctx.GetUniformLocation(program, "mvp")
	shader.uvp = ctx.GetUniformLocation(program, "uvp")
	shader.tex = ctx.GetUniformLocation(program, "tex")

	ctx.UseProgram(program)

	ctx.BindBuffer(gl.ARRAY_BUFFER, shader.vert)
	ctx.BufferData(gl.ARRAY_BUFFER, vert, gl.STATIC_DRAW)

	if useVao {
		shader.vao = ctx.CreateVertexArray()
		ctx.BindVertexArray(shader.vao)

		ctx.BindBuffer(gl.ARRAY_BUFFER, shader.vert)
		ctx.EnableVertexAttribArray(shader.vertLoc)
		ctx.VertexAttribPointer(shader.vertLoc, 2, gl.FLOAT, false, 4*4, 0)

		ctx.BindBuffer(gl.ARRAY_BUFFER, shader.vert)
		ctx.EnableVertexAttribArray(shader.clrVertLoc)
		ctx.VertexAttribPointer(shader.clrVertLoc, 2, gl.FLOAT, false, 4*4, 2*4)

	}

	return shader, nil
}

type shader struct {
	ctx gl.Context

	program gl.Program

	vao        gl.VertexArray
	vert       gl.Buffer
	vertLoc    gl.Attrib
	clrVertLoc gl.Attrib

	mvp gl.Uniform
	uvp gl.Uniform
	tex gl.Uniform

	active int

	useVao bool
}

func (s *shader) NewTexture(h, w int) {
	s.active++
}

func formatfs(opt *ShaderOption) (string, error) {

	var fs2d = `#version 100pe

precision mediump float;

uniform sampler2D tex;
varying  vec2 Tex2DCoords;
	%s 

void main()
{
	gl_FragColor = texture2D(tex, Tex2DCoords);

	%s
}`

	ext := ""

	for k, v := range opt.Uniforms {
		switch v {
		case "float", "vec2", "vec3", "mat2", "mat3", "mat4":
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

var vert = f32Bytes(
	0, 0, 0, 0, // top left
	+1, 0, 1, 0, // top right
	0, -1, 0, 1, // bottom left
	+1, -1, 1, 1, // bottom right
)

var vs2d = `#version 100

#define PI 3.14159265359

uniform mat3 mvp;
uniform mat2 uvp;

attribute  vec2 vert;
attribute vec2 clr_vert;

varying vec2 Tex2DCoords;


mat4 scale(float x, float y, float z) {
	return mat4(
	x, 0, 0, 0, 
	0, y, 0, 0, 
	0, 0, z, 0, 
	0, 0, 0, 1);
}

mat4 translate(float x, float y, float z){
	return mat4(
	1, 0, 0, 0, 
	0, 1, 0, 0, 
	0, 0, 1, 0, 
	x, y, z, 1);
}

mat4 ortho(float left, float right, float bottom, float top, float near, float far)  {
	float rml = right - left; 
	float tmb = top - bottom;
	float fmn = far - near;
	return mat4(
		2.0 / rml, 0, 0, 0,
		0, 2.0 / tmb, 0, 0, 
		0, 0, -2.0 / fmn, 0,
		-(right + left) / rml, -(top + bottom) / tmb, -(far + near) / fmn, 1
	);
}

mat4 lookat(vec3 eye, vec3 center, vec3 up) {

	vec3 d = normalize(eye-center);
	vec3 r = normalize(cross(up, d));
	vec3 u = cross(d, r);

	mat4 m = mat4(
		r[0], u[0], -d[0], 0,
		r[1], u[1], -d[1], 0,
		r[2], u[2], -d[2], 0,
		0, 0, 0, 1
	);

	return translate(-eye[0], -eye[1], -eye[2]) * m;
}

mat2 scale2d(vec2 _scale){
	return mat2(_scale.x,0.0,0.0, _scale.y);
}

mat4 rotate2d(float _angle){
	_angle = (_angle * PI) / 180.0;
    return mat4(cos(_angle), sin(_angle), 0, 0,
	                -sin(_angle), cos(_angle), 0 ,0,
					0,0,1,0,
					0,0,0,1);
}


void main()
{
	
	
	float vw = mvp[0].x;
	float vh = mvp[1].x;
	float w = mvp[0].y; 
	float h = mvp[1].y; 
	float px = mvp[0].z;
	float py = mvp[1].z;
	float sx = mvp[2].x;
	float sy = mvp[2].y;
	float angle = mvp[2].z;

	float tw = uvp[0].x; 
	float th = uvp[0].y; 
	float minx = uvp[1].x; 
	float miny = uvp[1].y; 

	float upx = (minx)/tw;
	float upy = (miny)/tw;       
	vec2 clr = clr_vert * scale2d(vec2(w/tw, h/th));
	Tex2DCoords = vec2(clr.x + upx, clr.y + upy);

	// Rectangle 
	mat2 size = mat2(w/vw , 0, 0, h/vh);;
	gl_Position = vec4(vert *size, 0, 1);
	gl_Position = rotate2d(angle) * scale(sx, sy, 1.0) * gl_Position;
	gl_Position = translate(px/vw, -py/vh, 0.0) * gl_Position;
	gl_Position = ortho(0.0, 1.0, -1.0, 0.0, -1.0, 0.1) * gl_Position;
}

`
