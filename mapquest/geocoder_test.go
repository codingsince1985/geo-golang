package mapquest_test

import (
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/mapquest"
	"testing"
)

func TestLocation(t *testing.T) {
	location, err := mapquest.Geocoder.Geocode("Melbourne VIC")
	if err != nil || location.Lat != -37.8142176 || location.Lng != 144.9631608 {
		t.Error("Geocode() failed", location)
	}
}

func TestAddress(t *testing.T) {
	address := mapquest.Geocoder.ReverseGeocode(geo.Location{-37.8142176, 144.9631608})
	if address != "Melbourne's GPO, Postal Lane, Melbourne, City of Melbourne, Greater Melbourne, Victoria, 3000, Australia" {
		t.Error("ReverseGeocode() failed", address)
	}
}
