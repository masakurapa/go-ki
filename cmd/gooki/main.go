/*
Package gooki はディレクトリの内容をツリー構造で出力するコマンドラインツールです


*/
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/masakurapa/gooki"
	"github.com/masakurapa/gooki/internal/opt"
)

var (
	help bool
)

func main() {
	path, option := parseFlagAndArgs()
	// if the -h option is specified, outputs usage and exit
	if help {
		usage()
		os.Exit(0)
	}

	k, err := gooki.Make(path, option)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, "")
		os.Exit(1)
	}

	if err := k.WriteTree(os.Stdout, option); err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, "")
		os.Exit(1)
	}
}

func parseFlagAndArgs() (string, opt.Option) {
	flag.Usage = func() {
		usage()
		os.Exit(2)
	}

	option := parseFlags()
	arg := flag.Arg(0)
	if arg == "" {
		arg = "."
	}

	return arg, option
}

func parseFlags() opt.Option {
	o := gooki.DefaultOption()

	flag.BoolVar(&o.AllFile, "a", o.AllFile, "Outputs all files. By default does not hidden files.")
	flag.BoolVar(&o.DirectoryOnly, "d", o.DirectoryOnly, "Outputs only directories.")

	flag.BoolVar(&o.ShowFullPath, "f", o.ShowFullPath, "Outputs full path prefix for each file.")

	flag.BoolVar(&help, "help", false, "Outputs a usage.")

	flag.Parse()
	return o
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
