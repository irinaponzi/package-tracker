package _package

// PackageDto is a Data Transfer Object for the Package entity
type PackageDto struct {
	TrackingCode string        `json:"tracking_code"`
	Status       PackageStatus `json:"status"`
	Size         PackageSize   `json:"size"`
	WeightKg     float64       `json:"weight_kg"`
	Destination  string        `json:"destination"`
}

// NewPackageDto creates a new PackageDto instance
func NewPackageDto(trackingCode string, status PackageStatus, size PackageSize, weightKg float64, destination string) *PackageDto {
	return &PackageDto{
		TrackingCode: trackingCode,
		Status:       status,
		Size:         size,
		WeightKg:     weightKg,
		Destination:  destination,
	}
}

// MapPackageToDto maps a Package entity to a PackageDto
func MapPackageToDto(p *Package) *PackageDto {
	return &PackageDto{
		TrackingCode: p.TrackingCode,
		Status:       p.Status,
		Size:         p.Size,
		WeightKg:     p.WeightKg,
		Destination:  p.Destination,
	}
}
