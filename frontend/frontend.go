//go:build !dev

package frontend

import (
	"embed"
	"io/fs"
)

var (
	//go:embed dist
	dist embed.FS
	Dist fs.FS
)

func init() {
	var err error
	Dist, err = fs.Sub(dist, "dist")
	if err != nil {
		panic(err)
	}
}
