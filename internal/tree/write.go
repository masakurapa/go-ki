package tree

import (
	"fmt"
	"io"

	"github.com/masakurapa/gooki/internal/dirinfo"
	"github.com/masakurapa/gooki/pkg/option"
)

func (nodes *Nodes) Write(out io.Writer, di dirinfo.DirInfo, opt option.Option) {
	w := writer{
		out: out,
		opt: opt,
	}

	fmt.Fprintln(out, di.Path())
	w.write(*nodes, []bool{})
	fmt.Fprintln(out, "")

	if opt.OnlyDirectory {
		fmt.Fprintln(out, fmt.Sprintf("%d directories", w.dirCnt))
	} else {
		fmt.Fprintln(out, fmt.Sprintf("%d directories, %d files", w.dirCnt, w.fileCnt))
	}
}

type writer struct {
	out     io.Writer
	opt     option.Option
	dirCnt  int
	fileCnt int
}

func (w *writer) write(nodes Nodes, showLineFlgs []bool) {
	max := len(nodes) - 1
	for i, node := range nodes {
		if node.fileInfo.IsDir() {
			w.dirCnt++
		} else {
			w.fileCnt++
		}

		lastNode := i == max

		fileName := node.fileInfo.Name()
		if node.fileInfo.IsSymlink() {
			fileName += " -> " + node.fileInfo.SymlinkFileName()
		}

		fmt.Fprintln(w.out, w.getRuledLine(lastNode, showLineFlgs)+fileName)
		if node.HasChild() {
			w.write(node.Child(), append(showLineFlgs, !lastNode))
		}
	}
}

func (*writer) getRuledLine(isLastNode bool, showLineFlgs []bool) string {
	line := ""
	for _, b := range showLineFlgs {
		if b {
			line += "│   "
		} else {
			line += "    "
		}
	}

	if isLastNode {
		line += "└── "
	} else {
		line += "├── "
	}

	return line
}
