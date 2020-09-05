package gooki

import (
	"io"
)

// Ki はディレクトリツリー全体を扱うインタフェースです
type Ki interface {
	// Eda はファイルまたはディレクトリを返します
	Eda() []Eda
	// Write はディレクトリツリーの情報をツリー形式で書き込みを行います
	Write(out io.Writer) error
}

// Eda はファイルまたはディレクトリ情報を階層化するためのインタフェースです
type Eda interface {
	// Child は子のファイルまたはディレクトリを返します
	Child() []Eda
	// Happa はファイルまたはディレクトリ情報を返します
	Happa() Happa
	// HasChild はEdaが子のファイルまたはディレクトリを持つかを返します。
	HasChild() bool
}

// Happa はファイルまたはディレクトリ情報を扱うためのインタフェースです
type Happa interface {
	// AbsPath はファイルまたはディレクトリ名を絶対パスを返します。
	AbsPath() string
	// RelPath はファイルまたはディレクトリ名をディレクトリツリーの起点からの相対パスを返します。
	RelPath() string
	// Dir はファイル名を除いたパスを返します
	Dir() string
	// Name はファイルまたはディレクトリ名を返します。
	Name() string
	// IsDir はディレクトリの場合にtrueを返します
	IsDir() bool
	// IsHiddenFile は隠しファイルの場合にtrueを返します
	IsHiddenFile() bool
	// IsSymlink はシンボリックリンクの場合にtrueを返します
	IsSymlink() bool
	// RealName はリンク元のファイル名を取得します
	RealName() string
}
