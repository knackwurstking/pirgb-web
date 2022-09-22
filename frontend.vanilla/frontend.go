package frontend

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed www/*
var www embed.FS

func GetFS() http.FileSystem {
  fs, err := fs.Sub(www, "www")
  if err != nil {
    panic(err)
  }
  return http.FS(fs)
}
