package tree

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/masakurapa/gooki/internal/dirinfo"
	"github.com/masakurapa/gooki/internal/fileinfo"
	"github.com/masakurapa/gooki/pkg/option"
)

type Nodes []Node

type Node struct {
	fileInfo fileinfo.FileInfo
	child    Nodes
}

func (n *Node) FileInfo() fileinfo.FileInfo {
	return n.fileInfo
}

func (n *Node) HasChild() bool {
	return len(n.Child()) > 0
}

func (n *Node) Child() Nodes {
	return n.child
}

func Make(di dirinfo.DirInfo, opt option.Option) (Nodes, error) {
	files := make([]fileinfo.FileInfo, 0)
	err := filepath.Walk(di.Abs(), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// ignore because it includes the root directory
		if path == di.Abs() {
			return nil
		}

		fi := fileinfo.New(strings.TrimPrefix(path, di.Abs()+"/"), info)

		if !opt.ShowHiddenFile && fi.IsHiddenFile() {
			return nil
		}

		files = append(files, fi)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return toNode(files, ""), nil
}

func toNode(files []fileinfo.FileInfo, base string) Nodes {
	nodes := make(Nodes, 0)
	for _, f := range files {
		// if base path is empty, skip all but the first level files.
		if base == "" && !f.IsFirstLevel() {
			continue
		}
		// skip if the file is not directly under the base path.
		if !f.IsParentDir(base) {
			continue
		}

		node := Node{fileInfo: f}
		if f.IsDir() {
			node.child = toNode(files, f.Path())
		}
		nodes = append(nodes, node)
	}

	return nodes
}
