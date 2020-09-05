/*
Package gooki はディレクトリの内容をツリー構造で扱うためのモジュールです
*/
package gooki

import (
	"fmt"
	"os"

	"github.com/masakurapa/gooki/internal/ki"
	"github.com/masakurapa/gooki/pkg/gooki"
)

// MakeByDefaultOption はデフォルトのオプションでディレクトリツリー情報を生成します。
func MakeByDefaultOption(path string) (gooki.Ki, error) {
	return Make(path, DefaultOption())
}

// Make はディレクトリツリー情報を生成します。
func Make(path string, option gooki.Option) (gooki.Ki, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if !fi.IsDir() {
		return nil, fmt.Errorf("%q is not directory", path)
	}
	return ki.Make(path, option)
}

/// DefaultOption はデフォルト値を設定したオプションを返します。
func DefaultOption() gooki.Option {
	return gooki.DefaultOption()
}
