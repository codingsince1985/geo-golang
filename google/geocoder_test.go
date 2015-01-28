package google_test

import (
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/google"
	"testing"
)

func TestLocation(t *testing.T) {
	location, err := google.Geocoder.Geocode("Melbourne VIC")
	if err != nil || location.Lat != -37.814107 || location.Lng != 144.96328 {
		t.Error("Geocode() failed", location)
	}
}

func TestAddress(t *testing.T) {
	address, err := google.Geocoder.ReverseGeocode(geo.Location{-37.814107, 144.96328})
	if err != nil || address != "338 Bourke Street, Melbourne VIC 3000, Australia" {
		t.Error("ReverseGeocode() failed", address)
	}
}
