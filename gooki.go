/*
Package gooki はディレクトリの内容をツリー構造で扱うためのモジュールです
*/
package gooki

import (
	"github.com/masakurapa/gooki/internal/ki"
	"github.com/masakurapa/gooki/pkg/gooki"
)

// MakeByDefaultOption はデフォルトのオプションでディレクトリツリー情報を生成します。
func MakeByDefaultOption(path string) (gooki.Ki, error) {
	return Make(path, DefaultOption())
}

// Make はディレクトリツリー情報を生成します。
func Make(path string, option gooki.Option) (gooki.Ki, error) {
	return ki.Make(path, option)
}

// DefaultOption はデフォルト値を設定したオプションを返します。
func DefaultOption() gooki.Option {
	return gooki.DefaultOption()
}
