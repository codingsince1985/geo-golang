package open_test

import (
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/mapquest/open"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var key = os.Getenv("MAPQUEST_OPEN_KEY")

var geocoder = open.Geocoder(key)

func TestGeocode(t *testing.T) {
	location, err := geocoder.Geocode("Melbourne VIC")
	assert.NoError(t, err)
	assert.Equal(t, geo.Location{Lat: -37.814218, Lng: 144.963161}, location)
}

func TestReverseGeocode(t *testing.T) {
	address, err := geocoder.ReverseGeocode(-37.816742, 144.964463)
	assert.NoError(t, err)
	assert.True(t, strings.HasSuffix(address, "Melbourne, Victoria, AU"))
}

func TestReverseGeocodeWithNoResult(t *testing.T) {
	_, err := geocoder.ReverseGeocode(-37.816742, 164.964463)
	assert.Equal(t, err, geo.ErrNoResult)
}
