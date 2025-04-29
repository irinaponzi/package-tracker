package tracking

import (
	"fmt"
	"github.com/irinaponzi/package-tracker/internal/country_codes"
	"time"
)

// maxTrackingCodesPerDayByCountry is a constant that defines the maximum number
// of tracking codes that can be generated per day for each country
const (
	maxTrackingCodesPerDayByCountry = 1_000_000
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
	date := time.Now().UTC().Truncate(24 * time.Hour)

	// country code
	countryCode, err := country_codes.GetCountryCode(country)
	if err != nil {
		return err
	}

	// check stock for the given country and date
	count, err := s.rp.CountByDateAndCountry(*countryCode, date)
	if err != nil {
		return err
	}
	available := maxTrackingCodesPerDayByCountry - count
	if available < amount {
		// todo create a custom error for this
		return fmt.Errorf("not enough tracking codes available for %s on %s: requested %d, available %d",
			country, date.Format("2006-01-02"), amount, available)
	}

	// get the last number for sequential numbers generation
	lastNumber, err := s.rp.GetLastSequence(*countryCode, date)
	if err != nil {
		return err
	}

	// generate sequential numbers
	sequences := s.generateSequentialNumbers(lastNumber, amount)

	// create tracking codes
	for _, seq := range sequences {
		newTrackingCode := NewTracking(*countryCode, date, seq)
		trackingCodes = append(trackingCodes, *newTrackingCode)
	}

	err = s.rp.SaveTrackingCode(trackingCodes)
	if err != nil {
		return err
	}
	return nil
}

// GenerateSequentialNumbers generates a slice of sequential numbers
func (s *TrackingService) generateSequentialNumbers(lastNumber int, amount int) []int {
	numbers := make([]int, amount)

	if lastNumber != 0 {
		lastNumber++
	}

	for i := range amount {
		numbers[i] = lastNumber + i
	}
	return numbers
}
