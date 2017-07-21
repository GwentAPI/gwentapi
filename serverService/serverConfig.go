package serverService

import (
	"net/http"
	"os"
)

func mountMedia(fileSystemPath string, mux *http.ServeMux) {
	fs := justFilesFilesystem{http.Dir(fileSystemPath)}
	mux.Handle("/media/", http.StripPrefix("/media/", http.FileServer(fs)))
}

type justFilesFilesystem struct {
	Fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.Fs.Open(name)

	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if stat.IsDir() {
		return nil, os.ErrNotExist
	}
	return f, nil
}
