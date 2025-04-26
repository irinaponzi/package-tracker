package batch

import (
	"database/sql"
)

// IBatchRepository is an interface that represents the batch repository
type IBatchRepository interface {
}

// BatchRepository is a struct that represents a batch repository
type BatchRepository struct {
	db *sql.DB
}

// NewBatchRepository creates a new BatchRepository instance
func NewBatchRepository(db *sql.DB) *BatchRepository {
	return &BatchRepository{
		db: db,
	}
}
