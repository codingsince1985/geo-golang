package mapquest_test

import (
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/mapquest"
	"strings"
	"testing"
)

var geocoder = mapquest.Geocoder()

func TestGeocode(t *testing.T) {
	location, err := geocoder.Geocode("Melbourne VIC")
	if err != nil || location.Lat != -37.8142176 || location.Lng != 144.9631608 {
		t.Error("TestGeocode() failed", err, location)
	}
}

func TestReverseGeocode(t *testing.T) {
	address, err := geocoder.ReverseGeocode(-37.8142176, 144.9631608)
	if err != nil || !strings.HasSuffix(address, "Melbourne, Victoria, 3000, Australia") {
		t.Error("TestReverseGeocode() failed", err, address)
	}
}

func TestReverseGeocodeWithNoResult(t *testing.T) {
	_, err := geocoder.ReverseGeocode(-37.8142176, 164.9631608)
	if err != geo.NoResultError {
		t.Error("TestReverseGeocodeWithNoResult() failed", err)
	}
}
