package logger

import (
	"encoding/json"
	"fmt"
	"time"
)

// LogEntry representa una entrada de log estructurada para IAs
type LogEntry struct {
	Timestamp string                 `json:"timestamp"`
	Level     string                 `json:"level"`
	Message   string                 `json:"message"`
	TraceID   string                 `json:"trace_id,omitempty"`
	Context   map[string]interface{} `json:"context,omitempty"`
}

func Log(level, msg, traceID string, ctx map[string]interface{}) {
	entry := LogEntry{
		Timestamp: time.Now().Format(time.RFC3339),
		Level:     level,
		Message:   msg,
		TraceID:   traceID,
		Context:   ctx,
	}

	b, _ := json.Marshal(entry)
	fmt.Println(string(b))
}

func Info(msg string, ctx map[string]interface{}) {
	Log("INFO", msg, "", ctx)
}

func Error(msg string, traceID string, ctx map[string]interface{}) {
	Log("ERROR", msg, traceID, ctx)
}
