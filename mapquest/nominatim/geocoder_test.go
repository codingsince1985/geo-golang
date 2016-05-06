package nominatim_test

import (
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/mapquest/nominatim"
	_ "github.com/orchestrate-io/dvr"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var key = os.Getenv("MAPQUEST_NOMINATUM_KEY")

var geocoder = nominatim.Geocoder(key)

func TestGeocode(t *testing.T) {
	location, err := geocoder.Geocode("Melbourne VIC")
	assert.NoError(t, err)
	assert.Equal(t, geo.Location{Lat: -37.8142176, Lng: 144.9631608}, location)
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
