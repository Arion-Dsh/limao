package graphics

import (
	"limao/graphics/gl"
	"sync"
)

func init() {
	g = &graphicsT{
		mu: sync.Mutex{},
		wr: sync.RWMutex{},
		// vw, vh was hack for webgl
		vw: 4,
		vh: 4,
	}
}

var g *graphicsT

type graphicsT struct {
	ctx gl.Context

	fb     framebuffer
	images images
	vw, vh int

	mu sync.Mutex
	wr sync.RWMutex

	executes chan func()
}

func Load(ctx gl.Context) {
	g.ctx = ctx
	// if g.fb != nil {
	// return
	// }
	// if g.ctx.Version() >= 3.0 {
	// g.fb = newframeBuffer(g.ctx, g.vw, g.vh)
	// }

}
func SetViewPort(w, h int) {

	g.mu.Lock()
	g.vw = w
	g.vh = h
	g.mu.Unlock()
	g.ctx.Viewport(0, 0, w, h)
	if g.fb != nil {
		g.fb.resize(w, h)
	}

}

func Clear() {
	g.ctx.ClearColor(255, 255, 255, 1)
	g.ctx.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func Draw() {
	// if g.fb == nil {
	// return
	// }
	g.mu.Lock()
	// g.fb.draw()
	g.ctx.Flush()
	g.mu.Unlock()
}

func Release() {
	g.mu.Lock()
	defer g.mu.Unlock()
	if g.images != nil {
		g.images.release()
	}
	//  if g.fb != nil {
	// g.fb.release()

	// }
}
