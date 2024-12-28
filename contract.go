package slogadapter

// Contract to provide reporting capabilities
// about an instance state
type EnabledLogger interface {
	Enabled() bool
}

// Logger implementation using [fmt.Print] as
// message renderer
type PrintLogger interface {
	Print(msg ...interface{})
}

// Logger implementation using [fmt.Printf] as
// message renderer
type PrintFormatLogger interface {
	Printf(format string, args ...interface{})
}

// Logger implementation using [fmt.Print] as
// message renderer
type PrintLineLogger interface {
	Println(msg ...interface{})
}

// Logger implementation using [fmt.Print] as
// message renderer
type LeveledLogger interface {
	Error(msg ...interface{})
	Info(msg ...interface{})
	Debug(msg ...interface{})
	Warn(msg ...interface{})
}

// Logger implementation using [fmt.Printf] as
// message renderer
type LeveledFormatLogger interface {
	Errorf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
}

// Logger implementation using [fmt.Print] as
// message renderer
type LeveledLineLogger interface {
	Errorln(msg ...interface{})
	Infoln(msg ...interface{})
	Debugln(msg ...interface{})
	Warnln(msg ...interface{})
}
