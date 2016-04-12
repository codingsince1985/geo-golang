package mapbox_test

import (
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/mapbox"
	"os"
	"strings"
	"testing"
)

var token = os.Getenv("MAPBOX_API_KEY")

var geocoder = mapbox.Geocoder(token)

func TestGeocode(t *testing.T) {
	if location, err := geocoder.Geocode("Melbourne VIC"); err != nil || location.Lat != -37.8142 || location.Lng != 144.9632 {
		t.Error("TestGeocode() failed", err, location)
	}
}

func TestReverseGeocode(t *testing.T) {
	if address, err := geocoder.ReverseGeocode(-37.8142, 144.9632); err != nil || !strings.HasSuffix(address, "Melbourne, Victoria 3000, Australia") {
		t.Error("TestReverseGeocode() failed", err, address)
	}
}

func TestReverseGeocodeWithNoResult(t *testing.T) {
	if _, err := geocoder.ReverseGeocode(-37.8142, 164.9632); err != geo.ErrNoResult {
		t.Error("TestReverseGeocodeWithNoResult() failed", err)
	}
}
