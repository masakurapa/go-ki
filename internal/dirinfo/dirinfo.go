package dirinfo

import (
	"fmt"
	"os"
	"path/filepath"
)

type DirInfo struct {
	path string
	abs  string
}

func New(val string) (*DirInfo, error) {
	fi, err := os.Stat(val)
	if err != nil {
		return nil, err
	}
	if !fi.IsDir() {
		return nil, fmt.Errorf("%q is not directory", val)
	}

	abs, err := filepath.Abs(val)
	if err != nil {
		return nil, err
	}

	return &DirInfo{
		path: val,
		abs:  abs,
	}, nil
}

func (d *DirInfo) Path() string {
	return d.path
}

func (d *DirInfo) Abs() string {
	return d.abs
}
