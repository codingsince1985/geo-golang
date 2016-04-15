package opencage_test

import (
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/opencage"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var key = os.Getenv("OPENCAGE_API_KEY")

var geocoder = opencage.Geocoder(key)

// locDelta is the lat long precision significant digits we care about
// because the opencage API isn't the same exact lat,long as other systems
var locDelta = 0.00001

func TestGeocode(t *testing.T) {
	location, err := geocoder.Geocode("Melbourne VIC")
	assert.NoError(t, err)
	assert.InDelta(t, -37.8142176, location.Lat, locDelta)
	assert.InDelta(t, 144.9631608, location.Lng, locDelta)
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
