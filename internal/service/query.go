package service

// QueryEngine manages the read path:
//   1. Parse query
//   2. Plan: select segments by time range, project columns
//   3. Parallel scan with filter pushdown
//   4. Hash aggregation + approximate accumulators
//   5. Merge across segments
//   6. Return result
type QueryEngine struct {
	// ports injected at construction
}
