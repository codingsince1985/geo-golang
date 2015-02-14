package open_test

import (
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/mapquest/open"
	"strings"
	"testing"
)

const key = "YOUR_KEY"

var geocoder = open.Geocoder(key)

func TestGeocode(t *testing.T) {
	location, err := geocoder.Geocode("Melbourne VIC")
	if err != nil || location.Lat != -37.814218 || location.Lng != 144.963161 {
		t.Error("TestGeocode() failed", err, location)
	}
}

func TestReverseGeocode(t *testing.T) {
	address, err := geocoder.ReverseGeocode(-37.816742, 144.964463)
	if err != nil || !strings.HasSuffix(address, "Melbourne, Victoria, AU") {
		t.Error("TestReverseGeocode() failed", err, address)
	}
}

func TestReverseGeocodeWithNoResult(t *testing.T) {
	_, err := geocoder.ReverseGeocode(-37.816742, 164.964463)
	if err != geo.NoResultError {
		t.Error("TestReverseGeocodeWithNoResult() failed", err)
	}
}
