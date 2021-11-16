// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build js
// +build js

package gl

import "syscall/js"

// Enum is equivalent to GLenum, and is normally used with one of the
// constants defined in this package.
type Enum uint32

// Attrib identifies the location of a specific attribute variable.
type Attrib struct {
	js.Value
}

// Program identifies a compiled shader program.
type Program struct {
	js.Value
}

// Shader identifies a GLSL shader.
type Shader struct {
	js.Value
}

// Buffer identifies a GL buffer object.
type Buffer struct {
	js.Value
}

// Framebuffer identifies a GL framebuffer.
type Framebuffer struct {
	js.Value
}

// A Renderbuffer is a GL object that holds an image in an internal format.
type Renderbuffer struct {
	js.Value
}

// A Texture identifies a GL texture unit.
type Texture struct {
	js.Value
}

// Uniform identifies the location of a specific uniform variable.
type Uniform struct {
	js.Value
}

// A VertexArray is a GL object that holds vertices in an internal format.
type VertexArray struct {
	js.Value
}
