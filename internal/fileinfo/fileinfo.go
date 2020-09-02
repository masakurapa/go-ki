package fileinfo

import (
	"os"
	"path/filepath"
	"strings"
)

type FileInfo struct {
	abs        string
	path       string
	dir        string
	info       os.FileInfo
	hiddenFile bool
	depth      int
}

func New(abs, path string, fi os.FileInfo) FileInfo {
	return FileInfo{
		abs:        abs,
		path:       path,
		dir:        filepath.Dir(path),
		info:       fi,
		hiddenFile: (strings.HasPrefix(path, ".") || strings.HasPrefix(fi.Name(), ".")),
		depth:      strings.Count(path, "/"),
	}
}

func (fi *FileInfo) Path() string {
	return fi.path
}

func (fi *FileInfo) Dir() string {
	return fi.dir
}

func (fi *FileInfo) IsParentDir(path string) bool {
	if path == "" {
		path = "."
	}
	if path == fi.path {
		return false
	}
	return fi.dir == path
}

func (fi *FileInfo) Name() string {
	return fi.info.Name()
}

func (fi *FileInfo) IsDir() bool {
	return fi.info.IsDir()
}

func (fi *FileInfo) IsHiddenFile() bool {
	return fi.hiddenFile
}

func (fi *FileInfo) Depth() int {
	return fi.depth
}

func (fi *FileInfo) IsFirstLevel() bool {
	return fi.Depth() == 0
}

func (fi *FileInfo) IsSymlink() bool {
	return fi.info.Mode()&os.ModeSymlink == os.ModeSymlink
}

func (fi *FileInfo) SymlinkFileName() string {
	realPath, err := os.Readlink(fi.abs)
	if err != nil {
		//TODO: error handling
		return ""
	}
	return filepath.Base(realPath)
}
