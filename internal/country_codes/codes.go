package country_codes

import "fmt"

var CountryCodes = map[string]string{
	"Argentina":      "AR",
	"Brazil":         "BR",
	"United States":  "US",
	"Mexico":         "MX",
	"Spain":          "ES",
	"Germany":        "DE",
	"France":         "FR",
	"Italy":          "IT",
	"Canada":         "CA",
	"Chile":          "CL",
	"Colombia":       "CO",
	"Peru":           "PE",
	"Uruguay":        "UY",
	"Paraguay":       "PY",
	"Bolivia":        "BO",
	"Venezuela":      "VE",
	"United Kingdom": "GB",
	"Australia":      "AU",
	"New Zealand":    "NZ",
	"Japan":          "JP",
	"China":          "CN",
	"South Korea":    "KR",
	"India":          "IN",
	"South Africa":   "ZA",
}

func GetCountryCode(countryName string) (*string, error) {
	code, exists := CountryCodes[countryName]
	if !exists {
		return nil, fmt.Errorf("invalid country name: %s", countryName)
	}
	return &code, nil
}
