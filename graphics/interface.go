package graphics

type GraphicsI interface {
	SetViewPort(ctx, w, h int)
	Clear()
	Draw()
	Release()
}

type images interface {
	newTexture2D(w, h int) Texture
	release()
}
