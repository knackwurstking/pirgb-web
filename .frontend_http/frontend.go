//go:build !dev

package frontend

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/sirupsen/logrus"
)

//go:embed dist/*
//go:embed dist/assets/*
//go:embed dist/fonts/*
var dist embed.FS

func GetFS() http.FileSystem {
	fs, err := fs.Sub(dist, "dist")
	if err != nil {
		logrus.Fatalln(err.Error())
	}

	return http.FS(fs)
}