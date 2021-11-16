//go:build android || ios
//+build android ios

package limao

import (
	masset "golang.org/x/mobile/asset"
)

func openFile(path string) (Asset, error) {
	a, er := masset.Open(path)
	return a, er

}
