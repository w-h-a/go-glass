// Package storage is the port interface for segment persistence.
// Segments are immutable once written. Read and write, never update.
package storage

import (
	"context"
	"time"

	"github.com/w-h-a/go-glass/internal/domain"
)

// Storage persists and retrieves immutable segments.
type Storage interface {
	Write(ctx context.Context, segment domain.Segment) error
	Read(ctx context.Context, id string) (domain.Segment, error)
	List(ctx context.Context, start, end time.Time) ([]string, error) // segment IDs in time range
	Delete(ctx context.Context, id string) error                      // for expiration
}
