package limao

type DrawOptions struct {
	GeoM     Geom
	Uniforms map[string][]float32
	Alpha    float32
}

type Drawable interface {
	Load()
	Draw(opts *DrawOptions)
}

func LoadDrawable(draws ...Drawable) {
	if !defaultApp.isRun {
		for _, d := range draws {
			defaultApp.draws = append(defaultApp.draws, d)
		}
		return
	}
	for _, d := range draws {
		d.Load()
	}

}
