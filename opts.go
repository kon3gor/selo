package selo

type locatorSettings struct {
	debug bool
}

type Option func(*locatorSettings)

func WithDebug(v bool) Option {
	return func(ls *locatorSettings) {
		ls.debug = v
	}
}
