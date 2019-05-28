package frenchapigouv_test

import (
	"strings"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/frenchapigouv"
	"github.com/stretchr/testify/assert"
)

func TestGeocode(t *testing.T) {
	ts := testServer(response1)
	defer ts.Close()

	geocoder := frenchapigouv.GeocoderWithURL(ts.URL + "/")
	location, err := geocoder.Geocode("Champ de Mars, 5 Avenue Anatole France, 75007 Paris")
	assert.Nil(t, err)
	assert.Equal(t, geo.Location{Lat: 48.859831, Lng: 2.328123}, *location)
}

func TestGeocodeWithNoResult(t *testing.T) {
	ts := testServer(response2)
	defer ts.Close()

	geocoder := frenchapigouv.GeocoderWithURL(ts.URL + "/")
	location, err := geocoder.Geocode("nowhere")
	assert.Nil(t, err)
	assert.Nil(t, location)
}

func TestReverseGeocode(t *testing.T) {
	ts := testServer(response1)
	defer ts.Close()

	geocoder := frenchapigouv.GeocoderWithURL(ts.URL + "/")
	address, err := geocoder.ReverseGeocode(48.859831, 2.328123)
	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(address.FormattedAddress, "5, Quai Anatole France,"))
}

func TestReverseGeocodeWithNoResult(t *testing.T) {
	ts := testServer(response2)
	defer ts.Close()

	geocoder := frenchapigouv.GeocoderWithURL(ts.URL + "/")
	addr, err := geocoder.ReverseGeocode(0, 0.34)
	assert.Nil(t, addr)
	assert.Nil(t, err)
}

func TestReverseGeocodeWithNoResultByDefaultPoints(t *testing.T) {
	ts := testServer(response3)
	defer ts.Close()

	geocoder := frenchapigouv.GeocoderWithURL(ts.URL + "/")
	addr, err := geocoder.ReverseGeocode(0, 0)
	assert.Nil(t, addr)
	assert.Nil(t, err)
}

func testServer(response string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte(response))
	}))
}

const (
	response1 = `[
		{
			 "type": "FeatureCollection",
			 "version": "draft",
			 "features": [{
				 "type": "Feature",
				 "geometry": {
					 "type": "Point",
					 "coordinates": [2.328123, 48.859831]
				 },
				 "properties": {
					 "label": "5 Quai Anatole France 75007 Paris",
					 "score": 0.4882581475128644,
					 "housenumber": "5",
					 "citycode": "75107",
					 "context": "75, Paris, Île-de-France",
					 "postcode": "75007",
					 "name": "5 Quai Anatole France",
					 "id": "ADRNIVX_0000000270768224",
					 "y": 6862409.3,
					 "importance": 0.2765,
					 "type": "housenumber",
					 "city": "Paris",
					 "x": 650705.5,
					 "street": "Quai Anatole France"
				 }
			 }],
			 "attribution": "BAN",
			 "licence": "ODbL 1.0",
			 "query": "Champ de Mars, 5 Avenue Anatole France, 75007 Paris",
			 "limit": 10
		}
	]`
	response2 = `{
		"type": "FeatureCollection",
		"version": "draft",
		"features": [],
		"attribution": "BAN",
		"licence": "ODbL 1.0",
		"limit": 1
	}`
	response3 = `{
		"type": "FeatureCollection",
		"version": "draft",
		"features": [{
			"type": "Feature",
			"geometry": {
				"type": "Point",
				"coordinates": [0.0, 0.0]
			},
			"properties": {
				"label": "baninfo",
				"score": 1.0,
				"name": "baninfo",
				"id": "db",
				"type": "info",
				"context": "20181125T223127",
				"distance": 0
			}
		}],
		"attribution": "BAN",
		"licence": "ODbL 1.0",
		"limit": 1
	}`
)
