package domain

import "time"

// Query is the observability query model: time range + filters + breakdowns + calculations.
// Not SQL. Narrow by design.
type Query struct {
	Start        time.Time
	End          time.Time
	Filters      []Filter
	Breakdowns   []string       // GROUP BY fields
	Calculations []Calculation
}

// Filter restricts events by field value.
type Filter struct {
	Field string
	Op    FilterOp
	Value any
}

// FilterOp is a filter comparison operator.
type FilterOp int

const (
	OpEq FilterOp = iota
	OpNeq
	OpGt
	OpLt
	OpGte
	OpLte
	OpExists
	OpContains
)

// Calculation is an aggregation to compute over matching events.
type Calculation struct {
	Op    CalcOp
	Field string // empty for COUNT
}

// CalcOp is an aggregation operation.
type CalcOp int

const (
	CalcCount CalcOp = iota
	CalcSum
	CalcAvg
	CalcMin
	CalcMax
	CalcP50
	CalcP95
	CalcP99
	CalcCountDistinct
)

// Result is the output of a query — time-bucketed, grouped aggregations.
type Result struct {
	Groups []Group
}

// Group is one row of query output.
type Group struct {
	Keys   map[string]any // breakdown field values
	Values map[string]any // calculation results
}
