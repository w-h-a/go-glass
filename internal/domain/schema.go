package domain

// Schema is the discovered set of columns with types and cardinality.
// Schema grows monotonically — new columns appear, columns never disappear.
// No DDL. Schema is discovered from data at flush time.
type Schema struct {
	Columns map[string]ColumnMeta
}

// ColumnMeta tracks a discovered column's type and cardinality estimate.
type ColumnMeta struct {
	Name        string
	Type        ColumnType
	Cardinality uint64 // estimated distinct values
}

// Discover infers schema from a batch of events.
func Discover(events []Event) Schema {
	return Schema{} // TODO
}
