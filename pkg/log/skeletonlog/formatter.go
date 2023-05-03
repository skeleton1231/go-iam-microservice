package skeletonlog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// Formatter is an interface for formatting log entries.
type Formatter interface {
	Format(*Entry) ([]byte, error)
}

// TextFormatter is a simple text-based log formatter.
type TextFormatter struct{}

// NewTextFormatter creates a new TextFormatter.
func NewTextFormatter() *TextFormatter {
	return &TextFormatter{}
}

// Format formats a log entry as text.
func (f *TextFormatter) Format(entry *Entry) ([]byte, error) {
	var buf bytes.Buffer

	// Timestamp
	buf.WriteString(entry.Timestamp.Format(time.RFC3339))
	buf.WriteString(" ")

	// Level
	buf.WriteString(entry.Level.String())
	buf.WriteString(" ")

	// Caller
	if entry.Caller != "" {
		buf.WriteString(entry.Caller)
		buf.WriteString(" ")
	}

	// Message
	buf.WriteString(entry.Message)

	// Fields
	if len(entry.Fields) > 0 {
		jsonFields, err := json.Marshal(entry.Fields)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal fields: %v", err)
		}
		buf.WriteString(" ")
		buf.Write(jsonFields)
	}

	buf.WriteString("\n")

	return buf.Bytes(), nil
}
