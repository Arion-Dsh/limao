// Copyright 2014 The Go Authors.  All rights reserved.
// license that can be found in the LICENSE file.

//go:build js
// +build js

package gl

import (
	"syscall/js"
)

func uint8Aarray(b []byte) js.Value {
	if b == nil {
		return js.Null()
	}
	data := js.Global().Get("Uint8Array").New(len(b))
	js.CopyBytesToJS(data, b)
	return data
}

func f32Array(f []float32) js.Value {
	data := js.Global().Get("Float32Array").New(len(f))
	// js.CopyBytesToJS(data, b)
	for i, v := range f {
		data.SetIndex(i, v)
	}
	return data
}

func (ctx *context) ActiveTexture(texture Enum) {
	ctx.gl.Call("activeTexture", uint32(texture))
}

func (ctx *context) AttachShader(p Program, s Shader) {
	ctx.gl.Call("attachShader", p, s)
}

func (ctx *context) BindAttribLocation(p Program, a Attrib, name string) {
	ctx.gl.Call("bindAttribLocation", p, a, name)
}

func (ctx *context) BindBuffer(target Enum, b Buffer) {
	ctx.gl.Call("bindBuffer", uint32(target), b)
}

func (ctx *context) BindFramebuffer(target Enum, fb Framebuffer) {
	ctx.gl.Call("bindFramebuffer", uint32(target), fb)
}

func (ctx *context) BindRenderbuffer(target Enum, rb Renderbuffer) {
	ctx.gl.Call("bindRenderbuffer", uint32(target), rb)
}

func (ctx *context) BindTexture(target Enum, t Texture) {
	ctx.gl.Call("bindTexture", uint32(target), t)
}

func (ctx *context) BindVertexArray(va VertexArray) {
	ctx.gl.Call("bindVertexArray", va)
}

func (ctx *context) BlendColor(red, green, blue, alpha float32) {
}

func (ctx *context) BlendEquation(mode Enum) {
}

func (ctx *context) BlendEquationSeparate(modeRGB, modeAlpha Enum) {
}

func (ctx *context) BlendFunc(sfactor, dfactor Enum) {
	ctx.gl.Call("blendFunc", uint32(sfactor), uint32(dfactor))
}

func (ctx *context) BlendFuncSeparate(sfactorRGB, dfactorRGB, sfactorAlpha, dfactorAlpha Enum) {
}

func (ctx *context) BufferData(target Enum, src []byte, usage Enum) {
	ctx.gl.Call("bufferData", uint32(target), uint8Aarray(src), uint32(usage))
}

func (ctx *context) BufferInit(target Enum, size int, usage Enum) {
	ctx.gl.Call("bufferInit", uint32(target), size, uint32(usage))
}

func (ctx *context) BufferSubData(target Enum, offset int, data []byte) {
	ctx.gl.Call("bufferSubData", uint32(target), offset, uint8Aarray(data))
}

func (ctx *context) CheckFramebufferStatus(target Enum) Enum {
	s := ctx.gl.Call("checkFramebufferStatus", int32(target))
	return Enum(s.Int())
}

func (ctx *context) Clear(mask Enum) {
	ctx.gl.Call("clear", uint32(mask))
}

func (ctx *context) ClearColor(red, green, blue, alpha float32) {
	ctx.gl.Call("clearColor", red, green, blue, alpha)
}

func (ctx *context) ClearDepthf(d float32) {
}

func (ctx *context) ClearStencil(s int) {
}

func (ctx *context) ColorMask(red, green, blue, alpha bool) {
}

func (ctx *context) CompileShader(s Shader) {
	ctx.gl.Call("compileShader", s)
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
	b := ctx.gl.Call("createBuffer")
	return Buffer{b}
}

func (ctx *context) CreateFramebuffer() Framebuffer {
	f := ctx.gl.Call("createFramebuffer")
	return Framebuffer{f}
}

func (ctx *context) CreateProgram() Program {
	p := ctx.gl.Call("createProgram")
	return Program{p}

}

func (ctx *context) CreateRenderbuffer() Renderbuffer {
	rb := ctx.gl.Call("createRenderbuffer")
	return Renderbuffer{rb}
}

func (ctx *context) CreateShader(ty Enum) Shader {
	s := ctx.gl.Call("createShader", int(ty))
	return Shader{s}
}

func (ctx *context) CreateTexture() Texture {
	t := ctx.gl.Call("createTexture")
	return Texture{t}
}

func (ctx *context) CreateVertexArray() VertexArray {
	va := ctx.gl.Call("createVertexArray")
	return VertexArray{va}
}

func (ctx *context) CullFace(mode Enum) {
}

func (ctx *context) DeleteBuffer(v Buffer) {
	ctx.gl.Call("deleteBuffer", v)
}

func (ctx *context) DeleteFramebuffer(v Framebuffer) {
	ctx.gl.Call("deleteFramebuffer", v)
}

func (ctx *context) DeleteProgram(p Program) {
	ctx.gl.Call("deleteProgram", p)
}

func (ctx *context) DeleteRenderbuffer(v Renderbuffer) {
	ctx.gl.Call("deleteRenderbuffer", v)
}

func (ctx *context) DeleteShader(s Shader) {
	ctx.gl.Call("deleteShader", s)
}

func (ctx *context) DeleteTexture(v Texture) {
	ctx.gl.Call("deleteTexture", v)
}

func (ctx *context) DeleteVertexArray(v VertexArray) {
	ctx.gl.Call("deleteVertexArray", v)
}

func (ctx *context) DepthFunc(fn Enum) {
	ctx.gl.Call("depthFunc", uint32(fn))
}

func (ctx *context) DepthMask(flag bool) {
}

func (ctx *context) DepthRangef(n, f float32) {
}

func (ctx *context) DetachShader(p Program, s Shader) {
}

func (ctx *context) Disable(cap Enum) {
	ctx.gl.Call("disable", uint32(cap))
}

func (ctx *context) DisableVertexAttribArray(a Attrib) {
	ctx.gl.Call("disableVertexAttribArray", a)
}

func (ctx *context) DrawArrays(mode Enum, first, count int) {
	ctx.gl.Call("drawArrays", uint32(mode), first, count)
}

func (ctx *context) DrawElements(mode Enum, count int, ty Enum, offset int) {
	ctx.gl.Call("drawElements", uint32(mode), count, uint32(ty), offset)
}

func (ctx *context) Enable(cap Enum) {
	ctx.gl.Call("enable", uint32(cap))
}

func (ctx *context) EnableVertexAttribArray(a Attrib) {
	ctx.gl.Call("enableVertexAttribArray", a)
}

func (ctx *context) Finish() {
	ctx.gl.Call("finish")
}

func (ctx *context) Flush() {
	ctx.gl.Call("flush")
}

func (ctx *context) FramebufferRenderbuffer(target, attachment, rbTarget Enum, rb Renderbuffer) {
	ctx.gl.Call("framebufferRenderbuffer", uint32(target), uint32(attachment), uint32(rbTarget), rb)
}

func (ctx *context) FramebufferTexture2D(target, attachment, texTarget Enum, t Texture, level int) {
	ctx.gl.Call("framebufferTexture2D", uint32(target), uint32(attachment), uint32(texTarget), t, level)
}

func (ctx *context) FrontFace(mode Enum) {
}

func (ctx *context) GenerateMipmap(target Enum) {
}

func (ctx *context) GetActiveAttrib(p Program, index uint32) (name string, size int, ty Enum) {
	au := ctx.gl.Call("getActiveAttrib", p, index)
	name = au.Get("name").String()
	size = au.Get("size").Int()
	ty = Enum(au.Get("type").Int())
	return
}

func (ctx *context) GetActiveUniform(p Program, index uint32) (name string, size int, ty Enum) {
	au := ctx.gl.Call("getActiveUniform", p, index)
	name = au.Get("name").String()
	size = au.Get("size").Int()
	ty = Enum(au.Get("type").Int())
	return
}

func (ctx *context) GetAttachedShaders(p Program) []Shader {

	return []Shader{}
}

func (ctx *context) GetAttribLocation(p Program, name string) Attrib {
	al := ctx.gl.Call("getAttribLocation", p, name)
	return Attrib{al}
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
	i := ctx.gl.Call("getBufferParameteri", uint32(target), uint32(value))
	if i.Bool() {
		return 1
	}
	return 0
}

func (ctx *context) GetError() Enum {
	return 0
}

func (ctx *context) GetFramebufferAttachmentParameteri(target, attachment, pname Enum) int {
	return 0
}

func (ctx *context) GetProgrami(p Program, pname Enum) int {
	i := ctx.gl.Call("getProgramParameter", p, uint32(pname))
	if pname == DELETE_STATUS || pname == LINK_STATUS || pname == VALIDATE_STATUS {
		if i.Bool() {
			return 1
		}
		return 0
	}
	return i.Int()
}

func (ctx *context) GetProgramInfoLog(p Program) string {
	info := ctx.gl.Call("getProgramInfoLog", p)
	return info.String()
}

func (ctx *context) GetRenderbufferParameteri(target, pname Enum) int {
	return 0
}

func (ctx *context) GetShaderi(s Shader, pname Enum) int {
	i := ctx.gl.Call("getShaderParameter", s, uint32(pname))
	if i.Bool() {
		return 1
	}
	return 0
}

func (ctx *context) GetShaderInfoLog(s Shader) string {
	info := ctx.gl.Call("getShaderInfoLog", s)
	return info.String()
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
	l := ctx.gl.Call("getUniformLocation", p, name)
	return Uniform{l}
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
	return false
}

func (ctx *context) IsEnabled(cap Enum) bool {
	return false
}

func (ctx *context) IsFramebuffer(fb Framebuffer) bool {
	return false
}

func (ctx *context) IsProgram(p Program) bool {
	return false
}

func (ctx *context) IsRenderbuffer(rb Renderbuffer) bool {
	return false
}

func (ctx *context) IsShader(s Shader) bool {
	return false
}

func (ctx *context) IsTexture(t Texture) bool {
	return false
}

func (ctx *context) LineWidth(width float32) {
}

func (ctx *context) LinkProgram(p Program) {
	ctx.gl.Call("linkProgram", p)
}

func (ctx *context) PixelStorei(pname Enum, param int32) {
}

func (ctx *context) PolygonOffset(factor, units float32) {
}

func (ctx *context) ReadPixels(dst []byte, x, y, width, height int, format, ty Enum) {
}

func (ctx *context) ReleaseShaderCompiler() {
}

func (ctx *context) RenderbufferStorage(target, internalFormat Enum, width, height int) {
	ctx.gl.Call("renderbufferStorage", uint32(target), uint32(internalFormat), width, height)
}

func (ctx *context) SampleCoverage(value float32, invert bool) {
}

func (ctx *context) Scissor(x, y, width, height int32) {
}

func (ctx *context) ShaderSource(s Shader, src string) {
	ctx.gl.Call("shaderSource", s, src)
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
	ctx.gl.Call("texImage2D", uint32(target), level, internalFormat, width, height, 0, uint32(format), uint32(ty), uint8Aarray(data))
}

func (ctx *context) TexSubImage2D(target Enum, level int, x, y, width, height int, format, ty Enum, data []byte) {
	ctx.gl.Call("texSubImage2D", uint32(target), level, x, y, width, height, uint32(format), uint32(ty), uint8Aarray(data))
}

func (ctx *context) TexParameterf(target, pname Enum, param float32) {
	ctx.gl.Call("texParameterf", uint32(target), uint32(pname), param)
}

func (ctx *context) TexParameterfv(target, pname Enum, params []float32) {
}

func (ctx *context) TexParameteri(target, pname Enum, param int) {
	ctx.gl.Call("texParameteri", uint32(target), uint32(pname), param)
}

func (ctx *context) TexParameteriv(target, pname Enum, params []int32) {
}

func (ctx *context) Uniform1f(dst Uniform, v float32) {
	ctx.gl.Call("uniform1f", dst, v)
}

func (ctx *context) Uniform1fv(dst Uniform, src []float32) {

	ctx.gl.Call("uniform1fv", dst, src)
}

func (ctx *context) Uniform1i(dst Uniform, v int) {
	ctx.gl.Call("uniform1i", dst, v)
}

func (ctx *context) Uniform1iv(dst Uniform, src []int32) {

	ctx.gl.Call("uniform1iv", dst, src)
}

func (ctx *context) Uniform2f(dst Uniform, v0, v1 float32) {

	ctx.gl.Call("uniform2f", dst, v0, v1)
}

func (ctx *context) Uniform2fv(dst Uniform, src []float32) {

	ctx.gl.Call("uniform2fv", dst, src)
}

func (ctx *context) Uniform2i(dst Uniform, v0, v1 int) {

	ctx.gl.Call("uniform2i", dst, v0, v1)
}

func (ctx *context) Uniform2iv(dst Uniform, src []int32) {

	ctx.gl.Call("uniform2iv", dst, src)
}

func (ctx *context) Uniform3f(dst Uniform, v0, v1, v2 float32) {

	ctx.gl.Call("uniform3f", dst, v0, v1, v2)
}

func (ctx *context) Uniform3fv(dst Uniform, src []float32) {
	ctx.gl.Call("uniform3fv", dst, f32Array(src))
}

func (ctx *context) Uniform3i(dst Uniform, v0, v1, v2 int32) {

	ctx.gl.Call("uniform3i", dst, v0, v1, v2)
}

func (ctx *context) Uniform3iv(dst Uniform, src []int32) {

	ctx.gl.Call("uniform3iv", dst, src)
}

func (ctx *context) Uniform4f(dst Uniform, v0, v1, v2, v3 float32) {

	ctx.gl.Call("uniform4f", dst, v0, v1, v2, v3)
}

func (ctx *context) Uniform4fv(dst Uniform, src []float32) {

	ctx.gl.Call("uniform4fv", dst, f32Array(src))
}

func (ctx *context) Uniform4i(dst Uniform, v0, v1, v2, v3 int32) {

	ctx.gl.Call("uniform4i", dst, v0, v1, v2, v3)
}

func (ctx *context) Uniform4iv(dst Uniform, src []int32) {
	ctx.gl.Call("uniform4iv", dst, src)
}

func (ctx *context) UniformMatrix2fv(dst Uniform, src []float32) {
	// r := ctx.gl.Get("FALSE")
	ctx.gl.Call("uniformMatrix2fv", dst, false, f32Array(src))
}

func (ctx *context) UniformMatrix3fv(dst Uniform, src []float32) {
	// r := ctx.gl.Get("FALSE")
	ctx.gl.Call("uniformMatrix3fv", dst, false, f32Array(src))
}

func (ctx *context) UniformMatrix4fv(dst Uniform, src []float32) {
	// r := ctx.gl.Get("FALSE")
	ctx.gl.Call("uniformMatrix4fv", dst, false, f32Array(src))
}

func (ctx *context) UseProgram(p Program) {
	ctx.gl.Call("useProgram", p)
}

func (ctx *context) ValidateProgram(p Program) {
}

func (ctx *context) VertexAttrib1f(dst Attrib, x float32) {
}

func (ctx *context) VertexAttrib1fv(dst Attrib, src []float32) {
}

func (ctx *context) VertexAttrib2f(dst Attrib, x, y float32) {
}

func (ctx *context) VertexAttrib2fv(dst Attrib, src []float32) {
}

func (ctx *context) VertexAttrib3f(dst Attrib, x, y, z float32) {
}

func (ctx *context) VertexAttrib3fv(dst Attrib, src []float32) {
}

func (ctx *context) VertexAttrib4f(dst Attrib, x, y, z, w float32) {
}

func (ctx *context) VertexAttrib4fv(dst Attrib, src []float32) {
}

func (ctx *context) VertexAttribPointer(dst Attrib, size int, ty Enum, normalized bool, stride, offset int) {
	ctx.gl.Call("vertexAttribPointer", dst, size, uint32(ty), normalized, stride, offset)
}

func (ctx *context) Viewport(x, y, width, height int) {
	ctx.gl.Call("viewport", x, y, width, height)
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
