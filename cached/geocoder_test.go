package cached_test

import (
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/cached"
	"github.com/codingsince1985/geo-golang/data"
	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"

	"strings"
	"testing"
	"time"
)

var geoCache = cache.New(5*time.Minute, 30*time.Second)

// geocoder is chained with one data geocoder with address -> location data
// the other has location -> address data
// this will exercise the chained fallback handling
var geocoder = cached.Geocoder(
	data.Geocoder(
		data.AddressToLocation{
			"Melbourne VIC": geo.Location{Lat: -37.814107, Lng: 144.96328},
		},
		data.LocationToAddress{
			geo.Location{Lat: -37.816742, Lng: 144.964463}: "Melbourne VIC 3000, Australia",
		},
	),
	geoCache,
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

func TestCachedGeocode(t *testing.T) {
	mock1 := data.Geocoder(
		data.AddressToLocation{
			"Austin,TX": geo.Location{Lat: 1, Lng: 2},
		},
		data.LocationToAddress{},
	)

	c := cached.Geocoder(mock1, geoCache)

	l, err := c.Geocode("Austin,TX")
	assert.NoError(t, err)
	assert.Equal(t, geo.Location{Lat: 1, Lng: 2}, l)

	// Should be cached
	// TODO: write a mock Cache impl to test cache is being used
	l, err = c.Geocode("Austin,TX")
	assert.NoError(t, err)
	assert.Equal(t, geo.Location{Lat: 1, Lng: 2}, l)

	_, err = c.Geocode("NOWHERE,TX")
	assert.Equal(t, geo.ErrNoResult, err)
}
