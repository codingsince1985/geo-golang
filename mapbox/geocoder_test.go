package mapbox_test

import (
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/mapbox"
	_ "github.com/orchestrate-io/dvr"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var token = os.Getenv("MAPBOX_API_KEY")

var geocoder = mapbox.Geocoder(token)

func TestGeocode(t *testing.T) {
	location, err := geocoder.Geocode("Melbourne VIC")
	assert.NoError(t, err)
	assert.Equal(t, geo.Location{Lat: -37.8142, Lng: 144.9632}, location)
}

func TestReverseGeocode(t *testing.T) {
	address, err := geocoder.ReverseGeocode(-37.8142, 144.9632)
	assert.NoError(t, err)
	assert.True(t, strings.HasSuffix(address, "Melbourne, Victoria 3000, Australia"))
}

func TestReverseGeocodeWithNoResult(t *testing.T) {
	_, err := geocoder.ReverseGeocode(-37.8142, 164.9632)
	assert.Equal(t, err, geo.ErrNoResult)
}
