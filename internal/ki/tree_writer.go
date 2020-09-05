package ki

import (
	"fmt"
	"io"
	"strings"

	"github.com/masakurapa/gooki/pkg/gooki"
)

type writer struct {
	out    io.Writer
	option gooki.Option
}

type treeWriter struct {
	writer
	basePath  string // 起点ディレクトリのパス
	dirCount  int
	fileCount int
}

func (w *treeWriter) Write(ed []gooki.Eda) error {
	fmt.Fprintln(w.writer.out, w.basePath)
	w.write(ed, []bool{})
	fmt.Fprintln(w.writer.out, "")

	if w.writer.option.DirectoryOnly {
		fmt.Fprintln(w.writer.out, fmt.Sprintf("%d directories", w.dirCount))
	} else {
		fmt.Fprintln(w.writer.out, fmt.Sprintf("%d directories, %d files", w.dirCount, w.fileCount))
	}

	return nil
}

func (w *treeWriter) write(ed []gooki.Eda, showLineFlgs []bool) {
	max := len(ed) - 1
	for i, e := range ed {
		if e.Happa().IsDir() {
			w.dirCount++
		} else {
			w.fileCount++
		}

		lastNode := i == max
		fmt.Fprintln(w.out, w.genRuledLine(lastNode, showLineFlgs)+w.genFileName(e))
		if e.HasChild() {
			w.write(e.Child(), append(showLineFlgs, !lastNode))
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

	if w.option.ShowFullPath {
		if strings.HasSuffix(w.basePath, "/") {
			name = w.basePath + ha.RelPath()
		} else {
			name = w.basePath + "/" + ha.RelPath()
		}
	}

	if ha.IsSymlink() {
		return name + " -> " + ha.RealName()
	}

	return name
}
