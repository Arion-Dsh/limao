//go:build js
//+build js

package limao

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

type jsfile struct {
	*bytes.Reader
}

func (f *jsfile) Close() error {
	return nil
}

func openFile(path string) (Asset, error) {
	if !filepath.IsAbs(path) {
		path = filepath.Join("assets", path)
	}
	res, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	f := &jsfile{bytes.NewReader(body)}
	return f, nil
}
