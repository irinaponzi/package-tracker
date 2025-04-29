package tracking

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

// BatchSize is the size of the batch for inserting tracking codes
const (
	BatchSize = 1000
)

// ITrackingRepository is an interface that represents the tracking code repository
type ITrackingRepository interface {
	GetLastSequence(countryCode string, date time.Time) (int, error)
	SaveTrackingCode(tks []Tracking) error
}

// TrackingRepository is a struct that represents a tracking code repository
type TrackingRepository struct {
	db *sql.DB
}

// NewTrackingRepository creates a new TrackingRepository instance
func NewTrackingRepository(db *sql.DB) *TrackingRepository {
	return &TrackingRepository{
		db: db,
	}
}

// GetLastSequence retrieves the last sequence number for a given country and date
func (r *TrackingRepository) GetLastSequence(countryCode string, date time.Time) (int, error) {
	var last sql.NullInt64
	query := `
		SELECT MAX(sequence)
		FROM tracking
		WHERE country = ? AND date = ?
	`
	err := r.db.QueryRow(query, countryCode, date).Scan(&last)
	if err != nil {
		return 0, fmt.Errorf("sql: error getting last sequence: %w", err)
	}

	// if there is no record, it starts from 0
	if !last.Valid {
		return 0, nil
	}
	return int(last.Int64), nil
}

// SaveTrackingCode saves a list of tracking codes to the database by inserting them in batches of BatchSize
// It calls the saveChunk method to save each batch of tracking codes
func (r *TrackingRepository) SaveTrackingCode(tks []Tracking) error {
	for i := 0; i < len(tks); i += BatchSize {
		end := i + BatchSize
		if end > len(tks) {
			end = len(tks)
		}

		batch := tks[i:end]

		err := r.saveChunk(batch)
		if err != nil {
			return err
		}
	}
	return nil
}

// saveChunk saves a chunk of tracking codes to the database
// It uses a transaction to ensure that all tracking codes are saved successfully
// if any error occurs, the transaction is rolled back
// if all tracking codes are saved successfully, the transaction is committed
func (r *TrackingRepository) saveChunk(tks []Tracking) error {
	// start transaction
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("sql: error starting transaction: %w", err)
	}

	// build query
	query := `
		INSERT INTO tracking (id, country, date, sequence, code)
		VALUES 
	`

	var args []interface{}
	var valueStrings []string

	for _, t := range tks {
		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?)")

		args = append(args,
			t.ID,
			t.CountryCode,
			t.Date,
			t.Sequence,
			t.Code,
		)
	}

	query += strings.Join(valueStrings, ", ")

	// execute inside the transaction
	_, err = tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("sql: error executing chunk insert: %w", err)
	}

	// commit if all ok
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("sql: error committing transaction: %w", err)
	}

	return nil
}
