package limao

import "limao/geom"

type Drawable interface {
	Load()
	Draw(opts *geom.Geom)
}

func LoadDrawable(draws ...Drawable) {
	if defaultApp.isRun != true {
		for _, d := range draws {
			defaultApp.draws = append(defaultApp.draws, d)
		}
		return
	}
	for _, d := range draws {
		d.Load()
	}

}
