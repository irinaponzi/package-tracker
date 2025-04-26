package tracking_code

import (
	"github.com/google/uuid"
	"time"
)

type TrackingStatus string

const (
	TrackingGenerated TrackingStatus = "generated"
	TrackingAssigned  TrackingStatus = "assigned"
	TrackingCancelled TrackingStatus = "cancelled"
)

type TrackingCode struct {
	ID        uuid.UUID
	Country   string
	Date      time.Time
	Sequence  int
	Code      string
	Status    TrackingStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}
