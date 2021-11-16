//go:build linux || darwin || (darwin && !ios) || windows
// +build linux darwin darwin,!ios windows

package limao

import (
	"os"
	"path/filepath"
)

func openFile(path string) (Asset, error) {
	if !filepath.IsAbs(path) {
		path = filepath.Join("assets", path)
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return f, nil
}
