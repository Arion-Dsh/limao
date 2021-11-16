// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build js
// +build js

package gl

import (
	"syscall/js"
)

type context struct {
	gl      js.Value
	version float32
}

type context3 struct {
	*context
}

func (ctx *context) Version() float32 {
	return ctx.version
}

func NewContext(canvasEl js.Value) (Context, error) {
	gl := canvasEl.Call("getContext", "webgl2")
	version := float32(2.0)
	if gl.IsNull() {
		gl = canvasEl.Call("getContext", "webgl")
	}
	webgl := &context{gl, version}
	return webgl, nil
}
