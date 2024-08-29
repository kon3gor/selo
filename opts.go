package selo

type locatorSettings struct {
	logger *proxyLogger
}

func newLocatorSettings() *locatorSettings {
	return &locatorSettings{
		logger: new(proxyLogger),
	}
}

type Option func(*locatorSettings)

func WithDebug(v bool) Option {
	return func(ls *locatorSettings) {
		if v {
			ls.logger.SetLevel(Debug)
		} else {
			ls.logger.SetLevel(Info)
		}
	}
}

func WithLogger(v Logger) Option {
	return func(ls *locatorSettings) {
		ls.logger.SetLogger(v)
	}
}
