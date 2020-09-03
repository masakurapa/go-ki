package gooki

import (
	"github.com/masakurapa/gooki/internal/dirinfo"
	"github.com/masakurapa/gooki/internal/tree"
	"github.com/masakurapa/gooki/pkg/option"
)

func Make(path string, opt option.Option) (tree.Nodes, error) {
	di, err := dirinfo.New(path)
	if err != nil {
		return nil, err
	}
	return tree.Make(*di, opt)
}

func DefaultOption() option.Option {
	return option.Default
}
