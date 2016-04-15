package openstreetmap_test

import (
	"strings"
	"testing"

	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/openstreetmap"
	"github.com/stretchr/testify/assert"
)

var geocoder = openstreetmap.Geocoder()

func TestGeocode(t *testing.T) {
	location, err := geocoder.Geocode("Melbourne VIC")
	assert.NoError(t, err)
	assert.Equal(t, geo.Location{Lat: -37.8142175, Lng: 144.9631608}, location)
}

func TestReverseGeocode(t *testing.T) {
	address, err := geocoder.ReverseGeocode(-37.816742, 144.964463)
	assert.NoError(t, err)
	assert.True(t, strings.HasSuffix(address, "Melbourne, Victoria, 3000, Australia"))
}

func TestReverseGeocodeWithNoResult(t *testing.T) {
	_, err := geocoder.ReverseGeocode(-37.816742, 164.964463)
	assert.Equal(t, err, geo.ErrNoResult)
}
