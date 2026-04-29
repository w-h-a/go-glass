// Package encoding contains column encoders and decoders.
// Pure transforms: typed values -> bytes (encode), bytes -> typed values (decode).
// Dictionary for strings, delta for timestamps, RLE for booleans, raw for numbers.
package encoding

// Encoder encodes a slice of typed values into bytes.
type Encoder interface {
	Encode(values []any) ([]byte, error)
}

// Decoder decodes bytes back into typed values.
type Decoder interface {
	Decode(data []byte) ([]any, error)
}
