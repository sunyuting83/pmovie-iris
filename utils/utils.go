package utils

import (
	"os"
	"path/filepath"
	"strings"
)

// GetDBPath get data file path
func GetDBPath(t string, d bool) string {
	var (
		p   string
		dir string
	)
	if d {
		dir = "/home/sun/Works/gopath/src/porn_movie"
	} else {
		path, err := os.Executable()
		if err != nil {
		}
		dir = filepath.Dir(path)
	}

	p = "movie.db"
	if t == "level" {
		p = "Cache"
	}
	return strings.Join([]string{dir, p}, "/")
}
