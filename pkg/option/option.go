package option

var (
	Default = Option{
		ShowHiddenFile: false,
		Help:           false,
	}
)

type Option struct {
	ShowHiddenFile bool
	Help           bool
}
