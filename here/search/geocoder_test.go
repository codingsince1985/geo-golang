package search_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/here/search"
)

var apiKey = os.Getenv("HERE_API_KEY")

func TestGeocode(t *testing.T) {
	ts := testServer(response1)
	defer ts.Close()

	geocoder := search.Geocoder(apiKey, ts.URL+"/")
	location, err := geocoder.Geocode("60 Collins St, Melbourne VIC 3000")
	require.NoError(t, err, "Geocode error")
	require.NotNil(t, location, "Geocode location")
	assert.Equal(t, geo.Location{Lat: -37.81375, Lng: 144.97176}, *location)
}

func TestReverseGeocode(t *testing.T) {
	ts := testServer(response2)
	defer ts.Close()

	geocoder := search.Geocoder(apiKey, ts.URL+"/")
	address, err := geocoder.ReverseGeocode(-37.81375, 144.97176)
	require.NoError(t, err, "ReverseGeocode address")
	require.NotNil(t, address, "ReverseGeocode address")
	assert.True(t, strings.HasPrefix(address.FormattedAddress, "56-64 Collins St"))
}

func TestReverseGeocodeWithNoResult(t *testing.T) {
	ts := testServer(response3)
	defer ts.Close()

	geocoder := search.Geocoder(apiKey, ts.URL+"/")
	addr, _ := geocoder.ReverseGeocode(-37.81375, 164.97176)
	assert.Nil(t, addr)
}

func testServer(response string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte(response))
	}))
}

const (
	response1 = `{
  "items": [
    {
      "title": "60 Collins St, Melbourne VIC 3000, Australia",
      "id": "here:af:streetsection:ek.gZtN5HHOffWV1FK-xUB:CggIBCDu357RAhABGgI2MChk",
      "resultType": "houseNumber",
      "houseNumberType": "PA",
      "address": {
        "label": "60 Collins St, Melbourne VIC 3000, Australia",
        "countryCode": "AUS",
        "countryName": "Australia",
        "stateCode": "VIC",
        "state": "Victoria",
        "city": "Melbourne",
        "street": "Collins St",
        "postalCode": "3000",
        "houseNumber": "60"
      },
      "position": {
        "lat": -37.81375,
        "lng": 144.97176
      },
      "access": [
        {
          "lat": -37.81393,
          "lng": 144.97184
        }
      ],
      "mapView": {
        "west": 144.97062,
        "south": -37.81465,
        "east": 144.9729,
        "north": -37.81285
      },
      "scoring": {
        "queryScore": 1,
        "fieldScore": {
          "state": 1,
          "city": 1,
          "streets": [
            1
          ],
          "houseNumber": 1,
          "postalCode": 1
        }
      }
    }
  ]
}`
	response2 = `{
  "items": [
    {
      "title": "56-64 Collins St, Melbourne VIC 3000, Australia",
      "id": "here:af:streetsection:ek.gZtN5HHOffWV1FK-xUB:CggIBCChp5agARABGgU1Ni02NA",
      "resultType": "houseNumber",
      "houseNumberType": "PA",
      "address": {
        "label": "56-64 Collins St, Melbourne VIC 3000, Australia",
        "countryCode": "AUS",
        "countryName": "Australia",
        "stateCode": "VIC",
        "state": "Victoria",
        "city": "Melbourne",
        "street": "Collins St",
        "postalCode": "3000",
        "houseNumber": "56-64"
      },
      "position": {
        "lat": -37.81375,
        "lng": 144.97176
      },
      "access": [
        {
          "lat": -37.81393,
          "lng": 144.97184
        }
      ],
      "distance": 0,
      "mapView": {
        "west": 144.95405,
        "south": -37.81909,
        "east": 144.97391,
        "north": -37.81323
      }
    }
  ]
}`
	response3 = `{
  "items": []
}`
)
