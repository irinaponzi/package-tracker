package _package

import (
	"github.com/google/uuid"
	tc "github.com/irinaponzi/package-tracker/internal/tracking_code"
	"time"
)

type PackageStatus string

const (
	PackageStatusCreated   PackageStatus = "created"
	PackageStatusInTransit PackageStatus = "in_transit"
	PackageStatusDelivered PackageStatus = "delivered"
	PackageStatusLost      PackageStatus = "lost"
	PackageStatusReturned  PackageStatus = "returned"
)

type PackageSize string

const (
	PackageSizeSmall  PackageSize = "small"
	PackageSizeMedium PackageSize = "medium"
	PackageSizeLarge  PackageSize = "large"
)

type Package struct {
	ID           uuid.UUID
	TrackingCode tc.TrackingCode
	Status       PackageStatus
	Size         PackageSize
	WeightKg     float64
	// todo: destination will be a struct in the next iteration
	Destination string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type PackageTracking struct {
	PackageID    uuid.UUID     `json:"package_id"`
	TrackingCode string        `json:"tracking_code"`
	Status       PackageStatus `json:"status"`
}
