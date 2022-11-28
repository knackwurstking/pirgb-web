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
				FrontentFiles = append(FrontentFiles, strings.Replace(filepath.Join(path, file.Name()), DistDir, "", 1))
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

	log.Debug.Printf("Endpoint: \"%s\"", pattern+"/")
	mux.HandleFunc(pattern+"/", indexHandler)

	for _, file := range FrontentFiles {
		fp := pattern + "/" + strings.TrimLeft(file, "/")
		log.Debug.Printf("Endpoint: \"%s\"", fp)
		mux.HandleFunc(fp, func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, filepath.Join("views", "dist", file))
		})
	}

	return mux
}
