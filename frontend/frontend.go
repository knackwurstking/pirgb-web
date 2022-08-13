//go:build !dev

package frontend

import (
	_ "embed"
	"io/fs"
)

// FIX: "frontend/frontend.go:11:5: go:embed cannot apply to var of type fs.FS"
//
//go:embed dist
var Dist fs.FS
