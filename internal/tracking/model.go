package tracking

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type TrackingStatus string

const (
	TrackingGenerated TrackingStatus = "generated"
	TrackingAssigned  TrackingStatus = "assigned"
	TrackingCancelled TrackingStatus = "cancelled"
)

type Tracking struct {
	ID          uuid.UUID
	CountryCode string
	Date        time.Time
	Sequence    int
	Code        string
	Status      TrackingStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTracking(countryCode string, date time.Time, sequence int) *Tracking {
	return &Tracking{
		ID:          uuid.New(),
		CountryCode: countryCode,
		Date:        date,
		Sequence:    sequence,
		Code:        generateTrackingCode(countryCode, date, sequence),
	}
}

func generateTrackingCode(countryCode string, date time.Time, sequence int) string {
	dateStr := date.Format("20060102")
	code := fmt.Sprintf("%s-%s-%06d", countryCode, dateStr, sequence)

	return code
}
