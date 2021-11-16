package graphics

import "limao/graphics/gl"

type GraphicsI interface {
	SetViewPort(ctx, w, h int)
	Clear()
	Draw()
	Release()
}

type framebuffer interface {
	id() gl.Framebuffer
	resize(w, h int)
	draw()
	release()
}

type images interface {
	newTexture2D(w, h int) Texture
	release()
}

type Texture interface {
	BindData([]uint8)
	release()
}
