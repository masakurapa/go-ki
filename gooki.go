/*
Package gooki はディレクトリの内容をツリー構造で扱うためのモジュールです
*/
package gooki

import (
	"fmt"
	"os"

	"github.com/masakurapa/gooki/internal/ki"
	"github.com/masakurapa/gooki/internal/opt"
)

// MakeByDefaultOption はデフォルトのオプションでディレクトリツリー情報を生成します。
func MakeByDefaultOption(path string) (ki.Ki, error) {
	return Make(path, DefaultOption())
}

// Make はディレクトリツリー情報を生成します。
func Make(path string, option opt.Option) (ki.Ki, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if !fi.IsDir() {
		return nil, fmt.Errorf("%q is not directory", path)
	}
	return ki.Make(path, option)
}

/*
DefaultOption はデフォルト値を設定したオプションを返します。

Default Values:
	AllFile:       false    隠しファイルを出力しない
	DirectoryOnly: false    ファイルを出力
	ShowFullPath:  false    ファイル名のみ出力
*/
func DefaultOption() opt.Option {
	return opt.Option{
		AllFile:       false,
		DirectoryOnly: false,
		ShowFullPath:  false,
	}
}
