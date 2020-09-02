package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/masakurapa/gooki/internal/tree"
)

func main() {
	di, opt := parseFlagAndArgs()
	// if the -h option is specified, outputs usage and exit
	if opt.Help {
		usage()
		os.Exit(0)
	}

	nodes, err := tree.Make(di, opt)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, "")
		os.Exit(1)
	}

	nodes.Write(os.Stdout, di, opt)
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

Synopsis:
	gooki [options...] directory

Options:`)

	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "")
}
