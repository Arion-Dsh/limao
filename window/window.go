package window

type Window interface {
	SetWindowSize(w, h int)
	SetTitle(t string)
	IsFullScreen() bool
	SetFullScreen() bool
}

type Opts struct {
	running bool

	Title  string
	Width  int
	Height int

	IsFullScreen bool

	MaxFPS int
}
