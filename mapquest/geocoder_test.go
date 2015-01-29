package mapquest_test

import (
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/mapquest"
	"testing"
)

func TestLocation(t *testing.T) {
	location, err := mapquest.Geocoder.Geocode("Melbourne VIC")
	if err != nil || location.Lat != -37.8142176 || location.Lng != 144.9631608 {
		t.Error("TestLocation() failed", location)
	}
}

func TestReverseGeocode(t *testing.T) {
	address, err := mapquest.Geocoder.ReverseGeocode(geo.Location{-37.8142176, 144.9631608})
	if err != nil || address != "Melbourne's GPO, Postal Lane, Melbourne, City of Melbourne, Greater Melbourne, Victoria, 3000, Australia" {
		t.Error("TestReverseGeocode() failed", address)
	}
}

func TestReverseGeocodeWithNoResult(t *testing.T) {
	_, err := mapquest.Geocoder.ReverseGeocode(geo.Location{-37.8142176, 164.9631608})
	if err != geo.NoResultError {
		t.Error("TestReverseGeocodeWithNoResult() failed", err)
	}
}
