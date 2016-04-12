package openstreetmap_test

import (
	"strings"
	"testing"

	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/openstreetmap"
)

var geocoder = openstreetmap.Geocoder()

func TestGeocode(t *testing.T) {
	if location, err := geocoder.Geocode("Melbourne VIC"); err != nil || location.Lat != -37.8142175 || location.Lng != 144.9631608 {
		t.Error("TestGeocode() failed", err, location)
	}
}

func TestReverseGeocode(t *testing.T) {
	if address, err := geocoder.ReverseGeocode(-37.816742, 144.964463); err != nil || !strings.HasSuffix(address, "Melbourne, Victoria, 3000, Australia") {
		t.Error("TestReverseGeocode() failed", err, address)
	}
}

func TestReverseGeocodeWithNoResult(t *testing.T) {
	if _, err := geocoder.ReverseGeocode(-37.816742, 164.964463); err != geo.ErrNoResult {
		t.Error("TestReverseGeocodeWithNoResult() failed", err)
	}
}
