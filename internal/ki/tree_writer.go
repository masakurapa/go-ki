package ki

import (
	"fmt"
	"io"
	"strings"

	"github.com/masakurapa/gooki/pkg/gooki"
)

type treeWriter struct {
	out io.Writer
	ki  ki
}

func (w *treeWriter) write() error {
	fmt.Fprintln(w.out, w.ki.originalPath)
	w.output(w.ki.Eda(), []bool{})
	fmt.Fprintln(w.out, "")

	if w.ki.option.DirectoryOnly {
		fmt.Fprintln(w.out, fmt.Sprintf("%d directories", w.ki.dirCount))
	} else {
		fmt.Fprintln(w.out, fmt.Sprintf("%d directories, %d files", w.ki.dirCount, w.ki.fileCount))
	}

	return nil
}

func (w *treeWriter) output(ed []gooki.Eda, showLineFlgs []bool) {
	max := len(ed) - 1
	for i, e := range ed {
		lastNode := i == max
		fmt.Fprintln(w.out, w.genRuledLine(lastNode, showLineFlgs)+w.genFileName(e))
		if e.HasChild() {
			w.output(e.Child(), append(showLineFlgs, !lastNode))
		}
	}
}

// ファイルの前に出力する罫線の組み立て
func (w *treeWriter) genRuledLine(isLastNode bool, showLineFlgs []bool) string {
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

// ファイル名の組み立て
func (w *treeWriter) genFileName(e gooki.Eda) string {
	ha := e.Happa()
	name := ha.Name()

	if w.ki.option.ShowFullPath {
		if strings.HasSuffix(w.ki.originalPath, "/") {
			name = w.ki.originalPath + ha.RelPath()
		} else {
			name = w.ki.originalPath + "/" + ha.RelPath()
		}
	}

	if ha.IsSymlink() {
		return name + " -> " + ha.RealName()
	}

	return name
}
