package google_test

import (
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/google"
	"testing"
)

func TestLocation(t *testing.T) {
	location, err := google.Geocoder.Geocode("Melbourne VIC")
	if err != nil || location.Lat != -37.814107 || location.Lng != 144.96328 {
		t.Error("TestLocation() failed", location)
	}
}

func TestReverseGeocode(t *testing.T) {
	address, err := google.Geocoder.ReverseGeocode(geo.Location{-37.814107, 144.96328})
	if err != nil || address != "338 Bourke Street, Melbourne VIC 3000, Australia" {
		t.Error("TestReverseGeocode() failed", address)
	}
}

func TestReverseGeocodeWithNoResult(t *testing.T) {
	_, err := google.Geocoder.ReverseGeocode(geo.Location{-37.814107, 164.96328})
	if err != geo.NoResultError {
		t.Error("TestReverseGeocodeWithNoResult() failed", err)
	}
}
