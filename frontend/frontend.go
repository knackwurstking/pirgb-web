package frontend

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed dist/*
//go:embed dist/assets/*
var dist embed.FS

func GetFS() http.FileSystem {
	_fs, _ := fs.Sub(dist, "dist")
	return http.FS(_fs)
}
