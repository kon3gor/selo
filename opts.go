package selo

type LocatorSettings struct {
	debug bool
}

type Option func(*LocatorSettings)

func WithDebug(v bool) Option {
	return func(ls *LocatorSettings) {
		ls.debug = v
	}
}
