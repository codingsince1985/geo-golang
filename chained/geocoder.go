package chained

import (
	"github.com/codingsince1985/geo-golang"
)

type chainedGeocoder struct{ Geocoders []geo.Geocoder }

// Geocoder creates a chain of Geocoders to lookup address and fallback on
func Geocoder(geocoders ...geo.Geocoder) geo.Geocoder { return chainedGeocoder{Geocoders: geocoders} }

// Geocode returns location for address
func (c chainedGeocoder) Geocode(address string) (*geo.Location, error) {
	// Geocode address by each geocoder until we get a real location response
	for i := range c.Geocoders {
		if l, err := c.Geocoders[i].Geocode(address); err == nil && l != nil {
			return l, nil
		}
		// skip error and try the next geocoder
		continue
	}
	// No geocoders found a result
	return nil, nil
}

// ReverseGeocode returns address for location
func (c chainedGeocoder) ReverseGeocode(lat, lng float64) (*geo.Address, error) {
	// Geocode address by each geocoder until we get a real location response
	for i := range c.Geocoders {
		if addr, err := c.Geocoders[i].ReverseGeocode(lat, lng); err == nil && addr != nil {
			return addr, nil
		}
		// skip error and try the next geocoder
		continue
	}
	// No geocoders found a result
	return nil, nil
}
