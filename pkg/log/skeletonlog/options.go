package skeletonlog

import "io"

// LogLevel is an enumeration for log levels.
type LogLevel int

// Enumeration of different log levels.
const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	PANIC
	FATAL
)

// String returns a string representation of the LogLevel.
func (l LogLevel) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case PANIC:
		return "PANIC"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// Option is a function type that modifies a logger configuration.
type Option func(*Config)

// Config is the logger configuration struct.
type Config struct {
	LogLevel     LogLevel
	Output       io.Writer
	Formatter    Formatter
	ReportCaller bool
}

// WithLogLevel is an Option that sets the log level.
func WithLogLevel(logLevel LogLevel) Option {
	return func(c *Config) {
		c.LogLevel = logLevel
	}
}

// WithOutput is an Option that sets the output destination.
func WithOutput(output io.Writer) Option {
	return func(c *Config) {
		c.Output = output
	}
}

// WithFormatter is an Option that sets the log formatter.
func WithFormatter(formatter Formatter) Option {
	return func(c *Config) {
		c.Formatter = formatter
	}
}

// WithReportCaller is an Option that enables or disables the reporting of the caller.
func WithReportCaller(reportCaller bool) Option {
	return func(c *Config) {
		c.ReportCaller = reportCaller
	}
}
