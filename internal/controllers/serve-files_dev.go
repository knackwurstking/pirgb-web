//go:build dev

package controllers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/knackwurstking/pirgb-web/pkg/log"
)

var (
	DistDir       = filepath.Join("views", "dist")
	FrontentFiles []string
)

func init() {
	CreateFrontentFiles()
}

func CreateFrontentFiles() {
	var createFrontentFiles func(string)

	createFrontentFiles = func(path string) {
		files, err := os.ReadDir(path)
		if err != nil {
			log.Error.Println(err)
			return
		}

		for _, file := range files {
			if file.Type().IsDir() {
				createFrontentFiles(filepath.Join(path, file.Name()))
			} else {
				FrontentFiles = append(FrontentFiles, filepath.Join(path, file.Name()))
			}
		}
	}

	createFrontentFiles(DistDir)
}

func ServeFiles(pattern string, mux *http.ServeMux) *http.ServeMux {
	var indexHandler = func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join("views", "dist", "index.html"))
	}

	pattern = strings.TrimRight(pattern, "/")

	if pattern != "" {
		log.Debug.Printf("Endpoint: \"%s\"", pattern)
		mux.HandleFunc(pattern, indexHandler)
	}

	log.Debug.Printf("Endpoint: \"%s\"", pattern)
	mux.HandleFunc(pattern+"/", indexHandler)

	for _, file := range FrontentFiles {
		log.Debug.Printf("Endpoint: \"%s\"", pattern)
		mux.HandleFunc(pattern+"/"+file, func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, filepath.Join("views", "dist", file))
		})
	}

	return mux
}
