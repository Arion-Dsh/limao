package limao

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"limao/geom"
	"limao/graphics"
	"limao/graphics/gl"
)

func init() {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("jpg", "jpg", jpeg.Decode, jpeg.DecodeConfig)
}

type images interface {
	NewTexture2D(w, h int) gl.Texture
}

func NewImage(w, h int) *Image {
	rgba := image.NewRGBA(image.Rect(0, 0, w, h))
	return newImage(rgba, w, h)
}

func NewImageFromRGBA(rgba *image.RGBA) *Image {
	w := rgba.Bounds().Dx()
	h := rgba.Bounds().Dy()
	return newImage(rgba, w, h)
}

func NewImageFromFile(path string) *Image {
	f, err := OpenAsset(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	wantSrc, _, err := image.Decode(f)

	if err != nil {
		panic(err)
	}

	w := wantSrc.Bounds().Dx()
	h := wantSrc.Bounds().Dy()
	rgba := image.NewRGBA(image.Rect(0, 0, w, h))
	img := newImage(rgba, w, h)
	draw.Draw(img.RGBA, wantSrc.Bounds(), wantSrc, image.ZP, draw.Src)

	return img
}

func newImage(rgba *image.RGBA, w, h int) *Image {
	r := image.Rect(0, 0, w, h)
	img := &Image{
		RGBA:   image.NewRGBA(r),
		rect:   r,
		width:  w,
		height: h,
	}
	img.geom = geom.New()
	return img
}

type Image struct {
	RGBA   *image.RGBA
	width  int
	height int
	geom   *geom.Geom
	tex    graphics.Texture
	rect   image.Rectangle //rgba position in tex

	subs  []*Image
	isSub bool

	subDraw []*Image
}

func (img *Image) Geom() *geom.Geom {
	return img.geom
}

func (img *Image) Dx() int {
	return img.width
}

func (img *Image) Dy() int {
	return img.height
}

func (img *Image) Fill(c color.Color) {
	draw.Draw(img.RGBA, img.RGBA.Bounds(), &image.Uniform{c}, image.ZP, draw.Src)
	if img.tex != nil {
		img.tex.BindData(img.RGBA.Pix)
	}
}

func (img *Image) Release() {

}

//SubImage  parent image's left top is min(0,0)
func (img *Image) SubImage(r image.Rectangle) *Image {
	if r.Dx() > img.width || r.Dy() > img.height {
		panic("Rectangle is big than image")
	}

	x := img.rect.Min.X + r.Min.X
	y := img.rect.Min.Y + r.Min.Y
	i := &Image{
		rect:   image.Rect(x, y, x+r.Dx(), y+r.Dy()),
		width:  r.Dx(),
		height: r.Dy(),
		geom:   geom.New(),
		isSub:  true,
	}
	if img.tex != nil {
		i.tex = img.tex
	}
	if img.subs == nil {
		img.subs = []*Image{}
	}
	img.subs = append(img.subs, i)
	return i
}

func (img *Image) Load() {
	if img.isSub {
		return
	}
	img.tex = graphics.NewTexture2D(img.width, img.height)
	img.tex.BindData(img.RGBA.Pix)

	var loadSub func(imgs []*Image)

	loadSub = func(imgs []*Image) {
		if imgs != nil {
			for _, i := range imgs {
				i.tex = img.tex
				loadSub(i.subs)
			}
		}
	}
	loadSub(img.subs)
}

func (img *Image) Draw(opts *geom.Geom) {
	if img.tex == nil {
		panic("limao Image must be preload. ")
	}
	g := geom.New().Add(img.geom)
	if opts != nil {
		g.Add(opts)
	}

	graphics.DrawTexture2D(img.tex, *g, img.rect)
}
