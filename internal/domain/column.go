package domain

// ColumnType determines encoding strategy.
type ColumnType int

const (
	ColumnString ColumnType = iota // dictionary encoding
	ColumnInt                      // raw encoding
	ColumnFloat                    // raw encoding
	ColumnBool                     // RLE encoding
	ColumnTime                     // delta encoding
)

// Column is a typed series of values from a single field across events.
type Column struct {
	Name     string
	Type     ColumnType
	Encoded  []byte // encoded column data
	Count    uint64
	Min, Max any // for segment index
}
