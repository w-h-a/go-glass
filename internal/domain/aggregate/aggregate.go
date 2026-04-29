// Package aggregate contains approximate accumulators.
// t-digest for percentiles, HyperLogLog for count-distinct.
// Merge is commutative, associative — same properties as CRDTs.
// Pure functions: no I/O, no side effects.
package aggregate

// Accumulator computes an approximate aggregate that can merge across segments.
type Accumulator interface {
	Add(value float64)
	Merge(other Accumulator) Accumulator
	Result() float64
}
