package limao

import (
	"limao/device"
	"limao/window"
	"sync"
	"time"
)

type Engine interface {
	PreLoad()
	Update()
	Draw()
}

var defaultApp *app

func init() {
	defaultApp = new(app)
}

type app struct {
	eng Engine

	isRun   bool
	startAt time.Time

	draws []Drawable
	mu    sync.Mutex
}

func (a *app) OnStart() {

	if a.isRun {
		return
	}

	ChangeScene(defaultApp.eng)
	a.mu.Lock()

	defaultApp.startAt = time.Now()
	for _, d := range a.draws {
		d.Load()
	}
	a.isRun = true
	a.mu.Unlock()
}

func (a *app) Update() {
	a.eng.Update()
}
func (a *app) Draw() {
	a.eng.Draw()
}

func (a *app) OnStop() {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.draws = []Drawable{}
	a.isRun = false
}

func ChangeScene(scenne Engine) {
	defaultApp.mu.Lock()
	defer defaultApp.mu.Unlock()
	defaultApp.eng = scenne
	scenne.PreLoad()
}

// Run ...
func Run(eng Engine, o *RunOpts) {

	opts := o.opts()
	defaultApp.eng = eng
	device.Run(defaultApp, opts)
}

type RunOpts struct {
	Title  string
	Width  int
	Height int
	MaxFPS int
}

func (o *RunOpts) verify() {

	if o == nil {
		o = new(RunOpts)
	}

	if o.Title == "" {
		o.Title = "owo"
	}
	if o.Width <= 0 {
		o.Width = 124
	}
	if o.Height <= 0 {
		o.Height = 124
	}
	if o.MaxFPS == 0 {
		o.MaxFPS = 60
	}
}
func (o *RunOpts) opts() *window.Opts {
	o.verify()
	return &window.Opts{Title: o.Title, Width: o.Width, Height: o.Height, MaxFPS: o.MaxFPS}
}
