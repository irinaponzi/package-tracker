package tracking

import (
	"fmt"
	"github.com/irinaponzi/package-tracker/internal/country_codes"
	"time"
)

// ITrackingService is an interface that represents the tracking code service
type ITrackingService interface {
	Create(amount int, country string) error
}

// TrackingService is a struct that represents a tracking code service
type TrackingService struct {
	rp ITrackingRepository
}

// NewTrackingService creates a new TrackingService instance
func NewTrackingService(rp ITrackingRepository) *TrackingService {
	return &TrackingService{
		rp: rp,
	}
}

// Create creates a new tracking code
func (s *TrackingService) Create(amount int, country string) error {
	var trackingCodes []Tracking
	date := time.Now()
	// todo: send date formated

	// country code // todo: refactor this
	countryCode, err := country_codes.GetCountryCode(country)
	if err != nil {
		return fmt.Errorf("invalid country code: %w", err)
	}

	// get the last number for sequential numbers generation
	lastNumber, err := s.rp.GetLastSequence(country, date)
	if err != nil {
		return fmt.Errorf("get last number: %w", err)
	}

	// generate sequential numbers
	sequences := s.GenerateSequentialNumbers(lastNumber, amount)
	// generate tracking code

	// create tracking codes
	for _, seq := range sequences {
		newTrackingCode := NewTracking(*countryCode, seq)
		trackingCodes = append(trackingCodes, *newTrackingCode)
	}

	err = s.rp.SaveTrackingCode(trackingCodes)
	if err != nil {
		return fmt.Errorf("save tracking codes: %w", err)
	}
	return nil
}

// GenerateSequentialNumbers generates a slice of sequential numbers
func (s *TrackingService) GenerateSequentialNumbers(lastNumber int, amount int) []int {
	numbers := make([]int, amount)

	if lastNumber != 0 {
		lastNumber++
	}

	for i := range amount {
		numbers[i] = lastNumber + i
	}
	return numbers
}
