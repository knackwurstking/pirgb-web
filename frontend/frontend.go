//go:build !dev

package frontend

import (
	_ "embed"
	"io/fs"
)

//go:embed dist
var Dist fs.FS
