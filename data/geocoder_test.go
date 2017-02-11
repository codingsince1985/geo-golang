package data_test

import (
	"strings"
	"testing"

	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/data"
	"github.com/stretchr/testify/assert"
)

var (
	addressFixture = geo.Address{
		FormattedAddress: "64 Elizabeth Street, Melbourne, Victoria 3000, Australia",
	}
	locationFixture = geo.Location{
		Lat: -37.814107,
		Lng: 144.96328,
	}
	geocoder = data.Geocoder(
		data.AddressToLocation{
			addressFixture: locationFixture,
		},
		data.LocationToAddress{
			locationFixture: addressFixture,
		},
	)
)

func TestGeocode(t *testing.T) {
	location, err := geocoder.Geocode(addressFixture.FormattedAddress)
	assert.NoError(t, err)
	assert.Equal(t, geo.Location{Lat: -37.814107, Lng: 144.96328}, *location)
}

func TestReverseGeocode(t *testing.T) {
	address, err := geocoder.ReverseGeocode(locationFixture.Lat, locationFixture.Lng)
	assert.Nil(t, err)
	assert.NotNil(t, address)
	assert.True(t, strings.Contains(address.FormattedAddress, "Melbourne, Victoria 3000, Australia"))
}

func TestReverseGeocodeWithNoResult(t *testing.T) {
	addr, err := geocoder.ReverseGeocode(1, 2)
	assert.Nil(t, err)
	assert.Nil(t, addr)
}
