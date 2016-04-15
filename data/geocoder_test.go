package data_test

import (
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

var geocoder = data.Geocoder(
	data.AddressToLocation{
		"Melbourne VIC": geo.Location{Lat: -37.814107, Lng: 144.96328},
	},
	data.LocationToAddress{
		geo.Location{Lat: -37.816742, Lng: 144.964463}: "Melbourne VIC 3000, Australia",
	},
)

func TestGeocode(t *testing.T) {
	location, err := geocoder.Geocode("Melbourne VIC")
	assert.NoError(t, err)
	assert.Equal(t, geo.Location{Lat: -37.814107, Lng: 144.96328}, location)
}

func TestReverseGeocode(t *testing.T) {
	address, err := geocoder.ReverseGeocode(-37.816742, 144.964463)
	assert.NoError(t, err)
	assert.Equal(t, "Melbourne VIC 3000, Australia", address)
}

func TestReverseGeocodeWithNoResult(t *testing.T) {
	_, err := geocoder.ReverseGeocode(-37.816742, 164.964463)
	assert.Equal(t, err, geo.ErrNoResult)
}
