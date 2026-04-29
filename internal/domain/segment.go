package domain

import "time"

// Segment is a time-partitioned, immutable collection of columns.
// Once written, a segment is never modified — only deleted when expired.
type Segment struct {
	ID      string
	Start   time.Time
	End     time.Time
	Columns []Column
	Index   SegmentIndex
}

// SegmentIndex holds per-column stats for query planning.
type SegmentIndex struct {
	ColumnStats []ColumnStat
}

// ColumnStat enables segment skipping during queries.
type ColumnStat struct {
	Name  string
	Min   any
	Max   any
	Count uint64
	// TODO: bloom filter for string columns
}
