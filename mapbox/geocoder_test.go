package mapbox_test

import (
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/mapbox"
	"strings"
	"testing"
)

const token = "YOUR_ACCESS_TOKEN"

var geocoder = mapbox.Geocoder(token)

func TestGeocode(t *testing.T) {
	location, err := geocoder.Geocode("Melbourne VIC")
	if err != nil || location.Lat != -37.8142 || location.Lng != 144.9632 {
		t.Error("TestGeocode() failed", err, location)
	}
}

func TestReverseGeocode(t *testing.T) {
	address, err := geocoder.ReverseGeocode(-37.8142, 144.9632)
	if err != nil || !strings.HasSuffix(address, "Melbourne, Victoria 3000, Australia") {
		t.Error("TestReverseGeocode() failed", err, address)
	}
}

func TestReverseGeocodeWithNoResult(t *testing.T) {
	_, err := geocoder.ReverseGeocode(-37.8142, 164.9632)
	if err != geo.ErrNoResult {
		t.Error("TestReverseGeocodeWithNoResult() failed", err)
	}
}
