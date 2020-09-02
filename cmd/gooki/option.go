package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/masakurapa/gooki/internal/dirinfo"
	"github.com/masakurapa/gooki/pkg/option"
)

func parseFlagAndArgs() (dirinfo.DirInfo, option.Option) {
	flag.Usage = func() {
		usage()
		os.Exit(2)
	}

	opt := option.Default
	parseFlags(&opt)

	return getDirInfo(), opt
}

func getDirInfo() dirinfo.DirInfo {
	arg := flag.Arg(0)
	if arg == "" {
		arg = "."
	}

	di, err := dirinfo.New(arg)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	return *di
}

func parseFlags(opt *option.Option) {
	flag.BoolVar(&opt.ShowHiddenFile, "a", opt.ShowHiddenFile, "Outputs all files. By default does not hidden files.")
	flag.BoolVar(&opt.Help, "help", opt.Help, "Outputs a usage.")

	flag.Parse()
}
