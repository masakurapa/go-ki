package option

var (
	Default = Option{
		ShowHiddenFile: false,
		OnlyDirectory:  false,
		ShowFullPath:   false,
		Help:           false,
	}
)

type Option struct {
	ShowHiddenFile bool
	OnlyDirectory  bool
	ShowFullPath   bool
	Help           bool
}
