package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/masakurapa/gooki/pkg/tree"
)

var (
	h = flag.Bool("h", false, "Outputs a usage.")
)

func main() {
	parseFlag()
	path := dir()
	files, err := tree.Walk(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, "")
		os.Exit(1)
	}

	outputTree(path, files)
}

func parseFlag() {
	flag.Usage = func() {
		usage()
		os.Exit(2)
	}
	flag.Parse()

	// if the -h option is specified, outputs usage and exit
	if *h {
		usage()
		os.Exit(0)
	}
}

func dir() string {
	arg := flag.Arg(0)
	if arg == "" {
		return "."
	}
	return arg
}

func outputTree(path string, files []tree.FileInfo) {
	fmt.Fprintln(os.Stdout, path)
	output(files, 0, "", false)
	fmt.Fprintln(os.Stdout, "")
	outputCount(files)

}

func output(files []tree.FileInfo, deap int, basePath string, hideBorder bool) {
	filteredFiles := filterByDeapAndPath(files, deap, basePath)
	max := len(filteredFiles) - 1
	for i, f := range filteredFiles {
		s := ""

		if f.Deap != 0 {
			if hideBorder {
				s += "    " + strings.Repeat(" ", (f.Deap-1)*4)
			} else {
				s += "│   " + strings.Repeat(" ", (f.Deap-1)*4)
			}
		}

		if i != max {
			fmt.Fprintln(os.Stdout, s+"├── "+f.Name)
		} else {
			fmt.Fprintln(os.Stdout, s+"└── "+f.Name)
		}

		if !f.IsDir {
			continue
		}

		if hideBorder {
			output(files, deap+1, f.Path, hideBorder)
		} else {
			output(files, deap+1, f.Path, (deap == 0 && i == max))
		}
	}
}

func filterByDeapAndPath(files []tree.FileInfo, deap int, path string) []tree.FileInfo {
	ret := make([]tree.FileInfo, 0)
	for _, f := range files {
		if f.Deap != deap {
			continue
		}

		if deap != 0 && !strings.HasPrefix(f.Path, path+"/") {
			continue
		}

		ret = append(ret, f)
	}
	return ret
}

func outputCount(files []tree.FileInfo) {
	dn := 0
	fn := 0
	for _, f := range files {
		if f.IsDir {
			dn++
		} else {
			fn++
		}
	}

	fmt.Fprintln(os.Stdout, fmt.Sprintf("%d directories, %d files", dn, fn))
}

func usage() {
	fmt.Fprintln(os.Stderr, `
	  _____             _  ___
	 / ____|           | |/ (_)
	| |  __  ___   ___ | ' / _
	| | |_ |/ _ \ / _ \|  < | |
	| |__| | (_) | (_) | . \| |
	 \_____|\___/ \___/|_|\_\_|

Description:
	'gooki' is list contents of directories in a tree-like format.

Options:`)

	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "")
}
