// Package domain contains the pure functional core of the engine.
// No I/O, no infrastructure imports.
package domain

import "time"

// Event is an arbitrarily wide structured event — the only primitive.
// Every span, log, or metric is an Event with different fields.
type Event struct {
	Timestamp time.Time
	TraceID   string
	SpanID    string
	Fields    map[string]any
}
