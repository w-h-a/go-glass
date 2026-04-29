// Package service contains the controller layer.
// Thin glue: receives events, buffers, encodes, writes segments.
package service

// Ingester manages the write path:
//   1. Receive events via OTLP
//   2. Buffer in memory, batched by time
//   3. Flush: encode columns, write immutable segment
//   4. Update schema registry
type Ingester struct {
	// ports injected at construction
}
