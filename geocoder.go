// Package geo is a generic framework to develop geocode/reverse geocode clients
package geo

import (
	"io"
	"log"
)

// Geocoder can look up (lat, long) by address and address by (lat, long)
type Geocoder interface {
	Geocode(address string) (*Location, error)
	ReverseGeocode(lat, lng float64) (*Address, error)
}

// Location is the output of Geocode
type Location struct {
	Lat, Lng float64
}

// Address is returned by ReverseGeocode.
// This is a structured representation of an address, including its flat representation
type Address struct {
	FormattedAddress string
	Street           string
	HouseNumber      string
	Suburb           string
	Postcode         string
	State            string
	StateCode        string
	StateDistrict    string
	County           string
	Country          string
	CountryCode      string
	City             string
}

// Logger is an implementation of StdLogger that geo uses to log its messages.
var Logger StdLogger = log.New(io.Discard, "[Geo]", log.LstdFlags)

// StdLogger is a interface for logging libraries.
type StdLogger interface {
	Printf(string, ...interface{})
}
