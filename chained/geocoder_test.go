package chained_test

import (
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/chained"
	"github.com/codingsince1985/geo-golang/data"
	"github.com/stretchr/testify/assert"

	"strings"
	"testing"
)

// geocoder is chained with one data geocoder with address -> location data
// the other has location -> address data
// this will exercise the chained fallback handling
var geocoder = chained.Geocoder(
	data.Geocoder(
		data.AddressToLocation{
			"Melbourne VIC": geo.Location{Lat: -37.814107, Lng: 144.96328},
		},
		data.LocationToAddress{},
	),

	data.Geocoder(
		data.AddressToLocation{},
		data.LocationToAddress{
			geo.Location{Lat: -37.816742, Lng: 144.964463}: "Melbourne VIC 3000, Australia",
		},
	),
)

func TestGeocode(t *testing.T) {
	location, err := geocoder.Geocode("Melbourne VIC")
	assert.NoError(t, err)
	assert.Equal(t, geo.Location{Lat: -37.814107, Lng: 144.96328}, location)
}

func TestReverseGeocode(t *testing.T) {
	address, err := geocoder.ReverseGeocode(-37.816742, 144.964463)
	assert.NoError(t, err)
	assert.True(t, strings.HasSuffix(address, "Melbourne VIC 3000, Australia"))
}

func TestReverseGeocodeWithNoResult(t *testing.T) {
	_, err := geocoder.ReverseGeocode(-37.816742, 164.964463)
	assert.Equal(t, err, geo.ErrNoResult)
}

func TestChainedGeocode(t *testing.T) {
	mock1 := data.Geocoder(
		data.AddressToLocation{
			"Austin,TX": geo.Location{Lat: 1, Lng: 2},
		},
		data.LocationToAddress{},
	)

	mock2 := data.Geocoder(
		data.AddressToLocation{
			"Dallas,TX": geo.Location{Lat: 3, Lng: 4},
		},
		data.LocationToAddress{},
	)

	c := chained.Geocoder(mock1, mock2)

	l, err := c.Geocode("Austin,TX")
	assert.NoError(t, err)
	assert.Equal(t, geo.Location{Lat: 1, Lng: 2}, l)

	l, err = c.Geocode("Dallas,TX")
	assert.NoError(t, err)
	assert.Equal(t, geo.Location{Lat: 3, Lng: 4}, l)

	_, err = c.Geocode("NOWHERE,TX")
	assert.Equal(t, geo.ErrNoResult, err)
}
