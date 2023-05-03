package skeletonlog

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

// Logger is the main structure for logging.
type Logger struct {
	config *Config
}

// NewLogger creates a new logger with the given options.
func NewLogger(options ...Option) *Logger {
	config := &Config{
		LogLevel:     INFO,
		Output:       os.Stdout,
		Formatter:    NewTextFormatter(),
		ReportCaller: false,
	}

	for _, option := range options {
		option(config)
	}

	return &Logger{config: config}
}

// Entry represents a log entry with a message and optional fields.
type Entry struct {
	Level     LogLevel
	Message   string
	Fields    map[string]interface{}
	Timestamp time.Time
	Caller    string
}

// logInternal handles the actual logging of the entry.
func (l *Logger) logInternal(level LogLevel, msg string, fields map[string]interface{}) {
	if level < l.config.LogLevel {
		return
	}

	entry := &Entry{
		Level:     level,
		Message:   msg,
		Fields:    fields,
		Timestamp: time.Now(),
	}

	if l.config.ReportCaller {
		pc, file, line, ok := runtime.Caller(3)
		if ok {
			f := runtime.FuncForPC(pc)
			entry.Caller = fmt.Sprintf("%s:%d %s", file, line, f.Name())
		}
	}

	formatted, err := l.config.Formatter.Format(entry)
	if err != nil {
		log.Printf("failed to format log entry: %v", err)
		return
	}

	if _, err := l.config.Output.Write(formatted); err != nil {
		log.Printf("failed to write log entry: %v", err)
	}
}

// Debug logs a debug message.
func (l *Logger) Debug(msg string, fields ...map[string]interface{}) {
	l.logInternal(DEBUG, msg, mergeFields(fields))
}

// Info logs an info message.
func (l *Logger) Info(msg string, fields ...map[string]interface{}) {
	l.logInternal(INFO, msg, mergeFields(fields))
}

// Warn logs a warning message.
func (l *Logger) Warn(msg string, fields ...map[string]interface{}) {
	l.logInternal(WARN, msg, mergeFields(fields))
}

// Error logs an error message.
func (l *Logger) Error(msg string, fields ...map[string]interface{}) {
	l.logInternal(ERROR, msg, mergeFields(fields))
}

// Panic logs a panic message and panics.
func (l *Logger) Panic(msg string, fields ...map[string]interface{}) {
	l.logInternal(PANIC, msg, mergeFields(fields))
	panic(msg)
}

// Fatal logs a fatal message and exits.
func (l *Logger) Fatal(msg string, fields ...map[string]interface{}) {
	l.logInternal(FATAL, msg, mergeFields(fields))
	os.Exit(1)
}

// mergeFields merges multiple fields maps into a single map.
func mergeFields(fieldsSlice []map[string]interface{}) map[string]interface{} {
	if len(fieldsSlice) == 0 {
		return nil
	}

	merged := make(map[string]interface{})
	for _, fields := range fieldsSlice {
		for k, v := range fields {
			merged[k] = v
		}
	}

	return merged
}
