package option

var (
	Default = Option{
		ShowHiddenFile: false,
		OnlyDirectory:  false,
		Help:           false,
	}
)

type Option struct {
	ShowHiddenFile bool
	OnlyDirectory  bool
	Help           bool
}
