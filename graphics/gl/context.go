// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux || darwin || android || windows
// +build linux darwin android windows

package gl

import (
	gles "golang.org/x/mobile/gl"
)

type context struct {
	gl      gles.Context
	worker  gles.Worker
	version float32
}

type context3 struct {
	*context
}

func (ctx *context) DoWork() {
	ctx.worker.DoWork()
}
func (ctx *context) WorkAvailable() <-chan struct{} { return ctx.worker.WorkAvailable() }

func (ctx *context) Version() float32 {
	return ctx.version
}
func NewContext() (Context, Worker) {
	// "GL_ES_2_0" or "GL_ES_3_0".
	ctx, worker := gles.NewContext()

	v := getVersion()

	c := &context{ctx, worker, v}
	if v < 3.0 {
		return c, c
	}

	return &context3{c}, c

}

func NewContextWithES(ctx interface{}) Context {
	esctx, _ := ctx.(gles.Context)

	v := getVersion()

	c := &context{esctx, nil, v}

	if v < 3.0 {
		return c
	}

	return &context3{c}
}
