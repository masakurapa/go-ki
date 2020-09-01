package tree

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type walk struct {
	// root path
	root string
	// a list of all directory and file info under the root path
	files []FileInfo
}

type FileInfo struct {
	Path  string
	Name  string
	IsDir bool
	Deap  int
}

func Walk(path string) ([]FileInfo, error) {
	if ok, err := isDir(path); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("%q is not directory", path)
	}

	w := walk{root: path}

	err := filepath.Walk(path, w.walkFunc)
	if err != nil {
		panic(err)
	}

	return w.files, nil
}

func isDir(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fi.IsDir(), nil
}

func (w *walk) walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	name := info.Name()
	// if the file starts with ".", ignore it as a hidden file
	if strings.HasPrefix(path, ".") || strings.HasPrefix(name, ".") {
		return nil
	}

	w.files = append(w.files, FileInfo{
		Path:  path,
		Name:  name,
		IsDir: info.IsDir(),
		Deap:  strings.Count(path, "/"),
	})

	return nil
}
