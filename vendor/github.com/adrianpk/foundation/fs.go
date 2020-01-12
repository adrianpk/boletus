package kabestan

import (
	"net/http"
	"strings"
)

// FileSystem custom file system handler
type FileSystem struct {
	FS http.FileSystem
}

// Open opens file
func (fs FileSystem) Open(path string) (http.File, error) {
	f, err := fs.FS.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := strings.TrimSuffix(path, "/") + "/index.html"
		if _, err := fs.FS.Open(index); err != nil {
			return nil, err
		}
	}

	return f, nil
}
