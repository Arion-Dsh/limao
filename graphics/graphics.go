package graphics

import (
	"limao/graphics/gl"
	"sync"
)

func init() {
	g = &graphicsT{
		mu: sync.Mutex{},
		// vw, vh was hack for webgl
		vw: 4,
		vh: 4,

		buffers: map[string]gl.Buffer{},
	}
}

var g *graphicsT

type graphicsT struct {
	ctx gl.Context

	buffers map[string]gl.Buffer

	fb framebuffer

	images images
	vw, vh int

	imageShader Shader

	mu sync.Mutex
}

func Load(ctx gl.Context) {

	g.ctx = ctx

	if g.imageShader == nil {

		s, err := New2Dshader(new(ShaderOption))

		if err != nil {
			panic("create dufualt images shader error")
		}
		g.imageShader = s

	}

}
func SetViewPort(w, h int) {

	g.mu.Lock()
	g.vw = w
	g.vh = h
	g.mu.Unlock()
	g.ctx.Viewport(0, 0, w, h)

}

func Clear() {
	g.ctx.ClearColor(255, 255, 255, 1)
	g.ctx.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func Draw() {
	g.mu.Lock()
	g.ctx.Flush()
	g.mu.Unlock()
}

func Release() {
	g.mu.Lock()
	defer g.mu.Unlock()
	if g.images != nil {
		g.images.release()
	}
}
