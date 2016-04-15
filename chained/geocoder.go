package chained

import (
	"github.com/codingsince1985/geo-golang"
)

type chainedGeocoder struct {
	Geocoders []geo.Geocoder
}

// Geocoder creates a chain of Geocoders to lookup address and fallback on
func Geocoder(geocoders ...geo.Geocoder) geo.Geocoder {
	return chainedGeocoder{Geocoders: geocoders}
}

// Geocode returns location for address
func (c chainedGeocoder) Geocode(address string) (geo.Location, error) {
	// Geocode address by each geocoder until we get a real location response
	for i := range c.Geocoders {
		l, err := c.Geocoders[i].Geocode(address)
		if err != nil {
			// skip error and try the next geocoder
			continue
		}
		return l, nil
	}
	// No geocoders found a result
	return geo.Location{}, geo.ErrNoResult
}

// ReverseGeocode returns address for location
func (c chainedGeocoder) ReverseGeocode(lat, lng float64) (string, error) {
	// Geocode address by each geocoder until we get a real location response
	for i := range c.Geocoders {
		addr, err := c.Geocoders[i].ReverseGeocode(lat, lng)
		if err != nil {
			// skip error and try the next geocoder
			continue
		}
		return addr, nil
	}
	// No geocoders found a result
	return "", geo.ErrNoResult
}
