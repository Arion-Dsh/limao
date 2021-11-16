// Copyright 2014 The Go Authors.  All rights reserved.
// license that can be found in the LICENSE file.

//go:build !js
// +build !js

package gl

import (
	gles "golang.org/x/mobile/gl"
)

func (ctx *context) ActiveTexture(t Enum) {
	ctx.gl.ActiveTexture(gles.Enum(t))

	// ctx.gl.Call("activeTexture", uint32(texture))
}

func (ctx *context) AttachShader(p Program, s Shader) {
	ctx.gl.AttachShader(gles.Program(p), gles.Shader(s))
}

func (ctx *context) BindAttribLocation(p Program, a Attrib, name string) {
	ctx.gl.BindAttribLocation(gles.Program(p), gles.Attrib(a), name)
}

func (ctx *context) BindBuffer(target Enum, b Buffer) {
	ctx.gl.BindBuffer(gles.Enum(target), gles.Buffer(b))
}

func (ctx *context) BindFramebuffer(target Enum, fb Framebuffer) {
	ctx.gl.BindFramebuffer(gles.Enum(target), gles.Framebuffer(fb))
}

func (ctx *context) BindRenderbuffer(target Enum, rb Renderbuffer) {
	ctx.gl.BindRenderbuffer(gles.Enum(target), gles.Renderbuffer(rb))
}

func (ctx *context) BindTexture(target Enum, t Texture) {
	ctx.gl.BindTexture(gles.Enum(target), gles.Texture(t))
}

func (ctx *context) BindVertexArray(va VertexArray) {
	ctx.gl.BindVertexArray(gles.VertexArray(va))
}

func (ctx *context) BlendColor(red, green, blue, alpha float32) {
	ctx.gl.BlendColor(red, green, blue, alpha)
}

func (ctx *context) BlendEquation(mode Enum) {
}

func (ctx *context) BlendEquationSeparate(modeRGB, modeAlpha Enum) {
}

func (ctx *context) BlendFunc(sfactor, dfactor Enum) {
	ctx.gl.BlendFunc(gles.Enum(sfactor), gles.Enum(dfactor))
}

func (ctx *context) BlendFuncSeparate(sfactorRGB, dfactorRGB, sfactorAlpha, dfactorAlpha Enum) {
}

func (ctx *context) BufferData(target Enum, src []byte, usage Enum) {
	ctx.gl.BufferData(gles.Enum(target), src, gles.Enum(usage))
}

func (ctx *context) BufferInit(target Enum, size int, usage Enum) {
	ctx.gl.BufferInit(gles.Enum(target), size, gles.Enum(usage))

}

func (ctx *context) BufferSubData(target Enum, offset int, data []byte) {
	ctx.gl.BufferSubData(gles.Enum(target), offset, data)
}

func (ctx *context) CheckFramebufferStatus(target Enum) Enum {
	t := ctx.gl.CheckFramebufferStatus(gles.Enum(target))
	return Enum(t)
}

func (ctx *context) Clear(mask Enum) {
	ctx.gl.Clear(gles.Enum(mask))
}

func (ctx *context) ClearColor(r, g, b, a float32) {
	ctx.gl.ClearColor(r, g, b, a)
}

func (ctx *context) ClearDepthf(d float32) {
	ctx.gl.ClearDepthf(d)
}

func (ctx *context) ClearStencil(s int) {
}

func (ctx *context) ColorMask(r, g, b, a bool) {
	ctx.gl.ColorMask(r, g, b, a)
}

func (ctx *context) CompileShader(s Shader) {
	ctx.gl.CompileShader(gles.Shader(s))
}

func (ctx *context) CompressedTexImage2D(target Enum, level int, internalformat Enum, width, height, border int, data []byte) {
}

func (ctx *context) CompressedTexSubImage2D(target Enum, level, xoffset, yoffset, width, height int, format Enum, data []byte) {
}

func (ctx *context) CopyTexImage2D(target Enum, level int, internalformat Enum, x, y, width, height, border int) {
}

func (ctx *context) CopyTexSubImage2D(target Enum, level, xoffset, yoffset, x, y, width, height int) {
}

func (ctx *context) CreateBuffer() Buffer {
	b := ctx.gl.CreateBuffer()
	return Buffer(b)
}

func (ctx *context) CreateFramebuffer() Framebuffer {
	f := ctx.gl.CreateFramebuffer()
	return Framebuffer(f)
}

func (ctx *context) CreateProgram() Program {
	p := ctx.gl.CreateProgram()
	return Program(p)

}

func (ctx *context) CreateRenderbuffer() Renderbuffer {
	rb := ctx.gl.CreateRenderbuffer()
	return Renderbuffer(rb)
}

func (ctx *context) CreateShader(ty Enum) Shader {
	s := ctx.gl.CreateShader(gles.Enum(ty))
	return Shader(s)
}

func (ctx *context) CreateTexture() Texture {
	t := ctx.gl.CreateTexture()
	return Texture(t)
}

func (ctx *context) CreateVertexArray() VertexArray {
	va := ctx.gl.CreateVertexArray()
	return VertexArray(va)
}

func (ctx *context) CullFace(mode Enum) {
}

func (ctx *context) DeleteBuffer(v Buffer) {
	ctx.gl.DeleteBuffer(gles.Buffer(v))
}

func (ctx *context) DeleteFramebuffer(v Framebuffer) {
	ctx.gl.DeleteFramebuffer(gles.Framebuffer(v))
}

func (ctx *context) DeleteProgram(p Program) {
	ctx.gl.DeleteProgram(gles.Program(p))
}

func (ctx *context) DeleteRenderbuffer(v Renderbuffer) {
	ctx.gl.DeleteRenderbuffer(gles.Renderbuffer(v))
}

func (ctx *context) DeleteShader(s Shader) {
	ctx.gl.DeleteShader(gles.Shader(s))
}

func (ctx *context) DeleteTexture(v Texture) {
	ctx.gl.DeleteTexture(gles.Texture(v))
}

func (ctx *context) DeleteVertexArray(v VertexArray) {
	ctx.gl.DeleteVertexArray(gles.VertexArray(v))
}

func (ctx *context) DepthFunc(fn Enum) {
	ctx.gl.DepthFunc(gles.Enum(fn))
}

func (ctx *context) DepthMask(flag bool) {
}

func (ctx *context) DepthRangef(n, f float32) {
}

func (ctx *context) DetachShader(p Program, s Shader) {
}

func (ctx *context) Disable(cap Enum) {
	ctx.gl.Disable(gles.Enum(cap))
}

func (ctx *context) DisableVertexAttribArray(a Attrib) {
	ctx.gl.DisableVertexAttribArray(gles.Attrib(a))
}

func (ctx *context) DrawArrays(mode Enum, first, count int) {
	ctx.gl.DrawArrays(gles.Enum(mode), first, count)
}

func (ctx *context) DrawElements(mode Enum, count int, ty Enum, offset int) {
	ctx.gl.DrawElements(gles.Enum(mode), count, gles.Enum(ty), offset)
}

func (ctx *context) Enable(cap Enum) {
	ctx.gl.Enable(gles.Enum(cap))
}

func (ctx *context) EnableVertexAttribArray(a Attrib) {
	ctx.gl.EnableVertexAttribArray(gles.Attrib(a))
}

func (ctx *context) Finish() {
	ctx.gl.Finish()
}

func (ctx *context) Flush() {
	ctx.gl.Flush()
}

func (ctx *context) FramebufferRenderbuffer(target, attachment, rbTarget Enum, rb Renderbuffer) {
	ctx.gl.FramebufferRenderbuffer(gles.Enum(target), gles.Enum(attachment), gles.Enum(rbTarget), gles.Renderbuffer(rb))
}

func (ctx *context) FramebufferTexture2D(target, attachment, texTarget Enum, t Texture, level int) {
	ctx.gl.FramebufferTexture2D(gles.Enum(target), gles.Enum(attachment), gles.Enum(texTarget), gles.Texture(t), level)
}

func (ctx *context) FrontFace(mode Enum) {
}

func (ctx *context) GenerateMipmap(target Enum) {
}

func (ctx *context) GetActiveAttrib(p Program, index uint32) (name string, size int, ty Enum) {
	name, size, tty := ctx.gl.GetActiveAttrib(gles.Program(p), index)
	ty = Enum(tty)
	return
}

func (ctx *context) GetActiveUniform(p Program, index uint32) (name string, size int, ty Enum) {
	name, size, tty := ctx.gl.GetActiveUniform(gles.Program(p), index)
	ty = Enum(tty)
	return
}

func (ctx *context) GetAttachedShaders(p Program) []Shader {

	return []Shader{}
}

func (ctx *context) GetAttribLocation(p Program, name string) Attrib {
	a := ctx.gl.GetAttribLocation(gles.Program(p), name)
	return Attrib(a)
}

func (ctx *context) GetBooleanv(dst []bool, pname Enum) {
}

func (ctx *context) GetFloatv(dst []float32, pname Enum) {
}

func (ctx *context) GetIntegerv(dst []int32, pname Enum) {
}

func (ctx *context) GetInteger(pname Enum) int {
	return 0
}

func (ctx *context) GetBufferParameteri(target, value Enum) int {
	return ctx.gl.GetBufferParameteri(gles.Enum(target), gles.Enum(value))
}

func (ctx *context) GetError() Enum {
	return 0
}

func (ctx *context) GetFramebufferAttachmentParameteri(target, attachment, pname Enum) int {
	return 0
}

func (ctx *context) GetProgrami(p Program, pname Enum) int {
	return ctx.gl.GetProgrami(gles.Program(p), gles.Enum(pname))
}

func (ctx *context) GetProgramInfoLog(p Program) string {
	return ctx.gl.GetProgramInfoLog(gles.Program(p))
}

func (ctx *context) GetRenderbufferParameteri(target, pname Enum) int {
	return 0
}

func (ctx *context) GetShaderi(s Shader, pname Enum) int {
	return ctx.gl.GetShaderi(gles.Shader(s), gles.Enum(pname))
}

func (ctx *context) GetShaderInfoLog(s Shader) string {
	return ctx.gl.GetShaderInfoLog(gles.Shader(s))
}

func (ctx *context) GetShaderPrecisionFormat(shadertype, precisiontype Enum) (rangeLow, rangeHigh, precision int) {
	return 0, 0, 0
}

func (ctx *context) GetShaderSource(s Shader) string {
	return ""
}

func (ctx *context) GetString(pname Enum) string {
	return ""
}

func (ctx *context) GetTexParameterfv(dst []float32, target, pname Enum) {
}

func (ctx *context) GetTexParameteriv(dst []int32, target, pname Enum) {
}

func (ctx *context) GetUniformfv(dst []float32, src Uniform, p Program) {
}

func (ctx *context) GetUniformiv(dst []int32, src Uniform, p Program) {
}

func (ctx *context) GetUniformLocation(p Program, name string) Uniform {
	u := ctx.gl.GetUniformLocation(gles.Program(p), name)
	return Uniform(u)
}

func (ctx *context) GetVertexAttribf(src Attrib, pname Enum) float32 {
	return 0
}

func (ctx *context) GetVertexAttribfv(dst []float32, src Attrib, pname Enum) {
}

func (ctx *context) GetVertexAttribi(src Attrib, pname Enum) int32 {
	return 0
}

func (ctx *context) GetVertexAttribiv(dst []int32, src Attrib, pname Enum) {
}

func (ctx *context) Hint(target, mode Enum) {
}

func (ctx *context) IsBuffer(b Buffer) bool {
	return ctx.gl.IsBuffer(gles.Buffer(b))
}

func (ctx *context) IsEnabled(cap Enum) bool {
	return false
}

func (ctx *context) IsFramebuffer(fb Framebuffer) bool {
	return ctx.gl.IsFramebuffer(gles.Framebuffer(fb))
}

func (ctx *context) IsProgram(p Program) bool {
	return ctx.gl.IsProgram(gles.Program(p))
}

func (ctx *context) IsRenderbuffer(rb Renderbuffer) bool {
	return ctx.gl.IsRenderbuffer(gles.Renderbuffer(rb))
}

func (ctx *context) IsShader(s Shader) bool {
	return ctx.gl.IsShader(gles.Shader(s))
}

func (ctx *context) IsTexture(t Texture) bool {
	return ctx.gl.IsTexture(gles.Texture(t))

}

func (ctx *context) LineWidth(width float32) {
}

func (ctx *context) LinkProgram(p Program) {
	ctx.gl.LinkProgram(gles.Program(p))
}

func (ctx *context) PixelStorei(pname Enum, param int32) {
	ctx.gl.PixelStorei(gles.Enum(pname), param)
}

func (ctx *context) PolygonOffset(factor, units float32) {
}

func (ctx *context) ReadPixels(dst []byte, x, y, width, height int, format, ty Enum) {
	ctx.gl.ReadPixels(dst, x, y, width, height, gles.Enum(format), gles.Enum(ty))
}

func (ctx *context) ReleaseShaderCompiler() {
}

func (ctx *context) RenderbufferStorage(target, internalFormat Enum, width, height int) {
	ctx.gl.RenderbufferStorage(gles.Enum(target), gles.Enum(internalFormat), width, height)
}

func (ctx *context) SampleCoverage(value float32, invert bool) {
}

func (ctx *context) Scissor(x, y, width, height int32) {
}

func (ctx *context) ShaderSource(s Shader, src string) {
	ctx.gl.ShaderSource(gles.Shader(s), src)
}

func (ctx *context) StencilFunc(fn Enum, ref int, mask uint32) {
}

func (ctx *context) StencilFuncSeparate(face, fn Enum, ref int, mask uint32) {
}

func (ctx *context) StencilMask(mask uint32) {
}

func (ctx *context) StencilMaskSeparate(face Enum, mask uint32) {
}

func (ctx *context) StencilOp(fail, zfail, zpass Enum) {
}

func (ctx *context) StencilOpSeparate(face, sfail, dpfail, dppass Enum) {
}

func (ctx *context) TexImage2D(target Enum, level int, internalFormat int, width, height int, format Enum, ty Enum, data []byte) {
	ctx.gl.TexImage2D(
		gles.Enum(target),
		level,
		internalFormat,
		width,
		height,
		gles.Enum(format),
		gles.Enum(ty),
		data,
	)
}

func (ctx *context) TexSubImage2D(target Enum, level int, x, y, width, height int, format, ty Enum, data []byte) {
	ctx.gl.TexSubImage2D(
		gles.Enum(target),
		level,
		x,
		y,
		width,
		height,
		gles.Enum(format),
		gles.Enum(ty),
		data,
	)
}

func (ctx *context) TexParameterf(target, pname Enum, param float32) {
	ctx.gl.TexParameterf(gles.Enum(target), gles.Enum(pname), param)
}

func (ctx *context) TexParameterfv(target, pname Enum, params []float32) {
}

func (ctx *context) TexParameteri(target, pname Enum, param int) {
	ctx.gl.TexParameteri(gles.Enum(target), gles.Enum(pname), param)
}

func (ctx *context) TexParameteriv(target, pname Enum, params []int32) {
}

func (ctx *context) Uniform1f(dst Uniform, v float32) {
	ctx.gl.Uniform1f(gles.Uniform(dst), v)
}

func (ctx *context) Uniform1fv(dst Uniform, src []float32) {
	ctx.gl.Uniform1fv(gles.Uniform(dst), src)
}

func (ctx *context) Uniform1i(dst Uniform, v int) {
	ctx.gl.Uniform1i(gles.Uniform(dst), v)
}

func (ctx *context) Uniform1iv(dst Uniform, src []int32) {
	ctx.gl.Uniform1iv(gles.Uniform(dst), src)
}

func (ctx *context) Uniform2f(dst Uniform, v0, v1 float32) {
	ctx.gl.Uniform2f(gles.Uniform(dst), v0, v1)

}

func (ctx *context) Uniform2fv(dst Uniform, src []float32) {
	ctx.gl.Uniform2fv(gles.Uniform(dst), src)
}

func (ctx *context) Uniform2i(dst Uniform, v0, v1 int) {
	ctx.gl.Uniform2i(gles.Uniform(dst), v0, v1)
}

func (ctx *context) Uniform2iv(dst Uniform, src []int32) {
	ctx.gl.Uniform2iv(gles.Uniform(dst), src)
}

func (ctx *context) Uniform3f(dst Uniform, v0, v1, v2 float32) {
	ctx.gl.Uniform3f(gles.Uniform(dst), v0, v1, v2)
}

func (ctx *context) Uniform3fv(dst Uniform, src []float32) {
	ctx.gl.Uniform3fv(gles.Uniform(dst), src)
}

func (ctx *context) Uniform3i(dst Uniform, v0, v1, v2 int32) {
	ctx.gl.Uniform3i(gles.Uniform(dst), v0, v1, v2)
}

func (ctx *context) Uniform3iv(dst Uniform, src []int32) {
	ctx.gl.Uniform3iv(gles.Uniform(dst), src)
}

func (ctx *context) Uniform4f(dst Uniform, v0, v1, v2, v3 float32) {
	ctx.gl.Uniform4f(gles.Uniform(dst), v0, v1, v2, v3)
}

func (ctx *context) Uniform4fv(dst Uniform, src []float32) {
	ctx.gl.Uniform4fv(gles.Uniform(dst), src)
}

func (ctx *context) Uniform4i(dst Uniform, v0, v1, v2, v3 int32) {
	ctx.gl.Uniform4i(gles.Uniform(dst), v0, v1, v2, v3)
}

func (ctx *context) Uniform4iv(dst Uniform, src []int32) {
	ctx.gl.Uniform4iv(gles.Uniform(dst), src)
}

func (ctx *context) UniformMatrix2fv(dst Uniform, src []float32) {
	ctx.gl.UniformMatrix2fv(gles.Uniform(dst), src)
}

func (ctx *context) UniformMatrix3fv(dst Uniform, src []float32) {
	ctx.gl.UniformMatrix3fv(gles.Uniform(dst), src)
}

func (ctx *context) UniformMatrix4fv(dst Uniform, src []float32) {
	ctx.gl.UniformMatrix4fv(gles.Uniform(dst), src)
}

func (ctx *context) UseProgram(p Program) {
	ctx.gl.UseProgram(gles.Program(p))
}

func (ctx *context) ValidateProgram(p Program) {
	ctx.gl.ValidateProgram(gles.Program(p))
}

func (ctx *context) VertexAttrib1f(dst Attrib, x float32) {
	ctx.gl.VertexAttrib1f(gles.Attrib(dst), x)
}

func (ctx *context) VertexAttrib1fv(dst Attrib, src []float32) {
	ctx.gl.VertexAttrib1fv(gles.Attrib(dst), src)
}

func (ctx *context) VertexAttrib2f(dst Attrib, x, y float32) {
	ctx.gl.VertexAttrib2f(gles.Attrib(dst), x, y)
}

func (ctx *context) VertexAttrib2fv(dst Attrib, src []float32) {
	ctx.gl.VertexAttrib2fv(gles.Attrib(dst), src)
}

func (ctx *context) VertexAttrib3f(dst Attrib, x, y, z float32) {
	ctx.gl.VertexAttrib3f(gles.Attrib(dst), x, y, z)
}

func (ctx *context) VertexAttrib3fv(dst Attrib, src []float32) {
	ctx.gl.VertexAttrib3fv(gles.Attrib(dst), src)
}

func (ctx *context) VertexAttrib4f(dst Attrib, x, y, z, w float32) {
	ctx.gl.VertexAttrib4f(gles.Attrib(dst), x, y, z, w)
}

func (ctx *context) VertexAttrib4fv(dst Attrib, src []float32) {
	ctx.gl.VertexAttrib4fv(gles.Attrib(dst), src)
}

func (ctx *context) VertexAttribPointer(dst Attrib, size int, ty Enum, normalized bool, stride, offset int) {
	ctx.gl.VertexAttribPointer(
		gles.Attrib(dst),
		size,
		gles.Enum(ty),
		normalized,
		stride,
		offset,
	)
}

func (ctx *context) Viewport(x, y, width, height int) {
	ctx.gl.Viewport(x, y, width, height)
}

func (ctx context3) UniformMatrix2x3fv(dst Uniform, src []float32) {
}

func (ctx context3) UniformMatrix3x2fv(dst Uniform, src []float32) {
}

func (ctx context3) UniformMatrix2x4fv(dst Uniform, src []float32) {
}

func (ctx context3) UniformMatrix4x2fv(dst Uniform, src []float32) {
}

func (ctx context3) UniformMatrix3x4fv(dst Uniform, src []float32) {
}

func (ctx context3) UniformMatrix4x3fv(dst Uniform, src []float32) {
}

func (ctx context3) BlitFramebuffer(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1 int, mask uint, filter Enum) {
}

func (ctx context3) Uniform1ui(dst Uniform, v uint32) {
}

func (ctx context3) Uniform2ui(dst Uniform, v0, v1 uint32) {
}

func (ctx context3) Uniform3ui(dst Uniform, v0, v1, v2 uint) {
}

func (ctx context3) Uniform4ui(dst Uniform, v0, v1, v2, v3 uint32) {
}
