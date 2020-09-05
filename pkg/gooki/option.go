package gooki

// Option はディレクトリツリー生成用のオプションを表します
type Option struct {
	// AllFile がtrueの場合、隠しファイルを含めたツリーを作成します
	// デフォルトは隠しファイル（ファイルまたはディレクトリ名が . から始まる）を除くツリーを作成します
	AllFile bool

	// DirectoryOnly がtrueの場合、ディレクトリのみでツリーを作成します
	// デフォルトはファイルも含めたツリーを作成します
	DirectoryOnly bool

	// ShowFullPath がtrueの場合外部出力時のファイル名にファイルのフルパスを付与して出力します
	// デフォルトはファイル名のみを出力します
	ShowFullPath bool

	// ModuleOption はモジュールから使うためのオプションです
	ModuleOption ModuleOption
}

// ModuleOption はモジュールから使うためのオプションです
type ModuleOption struct {
	// IgnoreTest がtrueの場合、_test.goで終わるテストファイルを無視します
	// デフォルトはテストファイルを含みます
	IgnoreTest bool
}

// DefaultOption はデフォルト値を設定したオプションを返します。
func DefaultOption() Option {
	return Option{
		AllFile:       false,
		DirectoryOnly: false,
		ShowFullPath:  false,
		ModuleOption: ModuleOption{
			IgnoreTest: false,
		},
	}
}
