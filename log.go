package selo

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warn
)

type Logger interface {
	Info(format string, v ...any)
	Warn(format string, v ...any)
	Debug(format string, v ...any)
	SetLevel(level LogLevel)
}

type proxyLogger struct {
	real  Logger
	level LogLevel
}

func (p *proxyLogger) SetLogger(l Logger) {
	p.real = l
}

func (p *proxyLogger) Info(format string, v ...any) {
	if p.real == nil {
		return
	}

	if Info >= p.level {
		p.real.Info(format, v)
	}
}

func (p *proxyLogger) Warn(format string, v ...any) {
	if p.real == nil {
		return
	}

	if Warn >= p.level {
		p.real.Warn(format, v)
	}
}

func (p *proxyLogger) Debug(format string, v ...any) {
	if p.real == nil {
		return
	}

	if Debug >= p.level {
		p.real.Debug(format, v)
	}
}

func (p *proxyLogger) SetLevel(level LogLevel) {
	p.level = level
}
