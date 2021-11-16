package limao

import (
	"io"
)

type Asset interface {
	io.ReadSeeker
	io.Closer
}

func OpenAsset(path string) (Asset, error) {
	return openFile(path)
}
