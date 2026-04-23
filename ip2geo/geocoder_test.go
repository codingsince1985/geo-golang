package ip2geo_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/codingsince1985/geo-golang/ip2geo"
	"github.com/stretchr/testify/assert"
)

func TestGeocode(t *testing.T) {
	ts := testServer(response1)
	defer ts.Close()

	geocoder := ip2geo.Geocoder("test-key", ts.URL)
	location, err := geocoder.Geocode("134.201.250.155")
	assert.NoError(t, err)
	assert.NotNil(t, location)
	assert.InDelta(t, 34.0544, location.Lat, 0.01)
	assert.InDelta(t, -118.244, location.Lng, 0.01)
}

func TestGeocodeNoCity(t *testing.T) {
	ts := testServer(responseNoCity)
	defer ts.Close()

	geocoder := ip2geo.Geocoder("test-key", ts.URL)
	location, err := geocoder.Geocode("8.8.8.8")
	assert.NoError(t, err)
	assert.NotNil(t, location)
	assert.InDelta(t, 37.751, location.Lat, 0.01)
	assert.InDelta(t, -97.822, location.Lng, 0.01)
}

func TestGeocodeError(t *testing.T) {
	ts := testServer(responseError)
	defer ts.Close()

	geocoder := ip2geo.Geocoder("bad-key", ts.URL)
	location, err := geocoder.Geocode("8.8.8.8")
	assert.Error(t, err)
	assert.Nil(t, location)
}

func TestReverseGeocode(t *testing.T) {
	geocoder := ip2geo.Geocoder("test-key")
	address, err := geocoder.ReverseGeocode(34.0, -118.0)
	assert.Error(t, err)
	assert.Nil(t, address)
}

func TestGeocodeVerifiesHeader(t *testing.T) {
	var receivedKey string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedKey = r.Header.Get("X-Api-Key")
		w.Write([]byte(response1))
	}))
	defer ts.Close()

	geocoder := ip2geo.Geocoder("my-secret-key", ts.URL)
	geocoder.Geocode("134.201.250.155")
	assert.Equal(t, "my-secret-key", receivedKey)
}

func testServer(response string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(response))
	}))
}

var response1 string
var responseNoCity string
var responseError string

func init() {
	r1, _ := json.Marshal(map[string]interface{}{
		"success": true,
		"code":    200,
		"data": map[string]interface{}{
			"ip":   "134.201.250.155",
			"type": "ipv4",
			"continent": map[string]interface{}{
				"name": "North America",
				"code": "na",
				"country": map[string]interface{}{
					"name": "United States",
					"code": "us",
					"subdivision": map[string]interface{}{
						"name": "California",
						"code": "CA",
					},
					"city": map[string]interface{}{
						"name":        "Los Angeles",
						"latitude":    34.0544,
						"longitude":   -118.244,
						"postal_code": "90060",
						"timezone": map[string]interface{}{
							"name": "America/Los_Angeles",
						},
					},
				},
			},
			"asn": map[string]interface{}{
				"number": 25876,
				"name":   "Los Angeles Department Of Water & Power",
			},
		},
	})
	response1 = string(r1)

	r2, _ := json.Marshal(map[string]interface{}{
		"success": true,
		"code":    200,
		"data": map[string]interface{}{
			"ip":   "8.8.8.8",
			"type": "ipv4",
			"continent": map[string]interface{}{
				"name": "North America",
				"code": "na",
				"country": map[string]interface{}{
					"name": "United States",
					"code": "us",
					"city": map[string]interface{}{
						"latitude":  37.751,
						"longitude": -97.822,
					},
				},
			},
		},
	})
	responseNoCity = string(r2)

	r3, _ := json.Marshal(map[string]interface{}{
		"success": false,
		"code":    401,
		"message": "Invalid API key",
	})
	responseError = string(r3)
}
