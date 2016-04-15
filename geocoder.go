// Package geo is a generic framework to develop geocode/reverse geocode clients
package geo

// Geocoder can look up (lat, long) by address and address by (lat, long)
type Geocoder interface {
	Geocode(address string) (Location, error)
	ReverseGeocode(lat, lng float64) (string, error)
}
