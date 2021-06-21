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
	assert.Equal(t, "Île-de-France", address.State)
	assert.Equal(t, "Paris", address.County)
}

func TestReverseGeocodeWithTypeStreet(t *testing.T) {
	ts := testServer(responseStreet)
	defer ts.Close()
	geocoder := frenchapigouv.GeocoderWithURL(ts.URL + "/")
	address, err := geocoder.ReverseGeocode(50.720114, 3.156717)
	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(address.FormattedAddress, "Rue des Anges,"))
	assert.Equal(t, "Hauts-de-France", address.State)
	assert.Equal(t, "Nord", address.County)
}

func TestReverseGeocodeWithTypeLocality(t *testing.T) {
	ts := testServer(responseLocality)
	defer ts.Close()
	geocoder := frenchapigouv.GeocoderWithURL(ts.URL + "/")
	address, err := geocoder.ReverseGeocode(44.995637, 1.646584)
	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(address.FormattedAddress, "Route de Saint Denis les Martel (Les Quatre-Routes-du-Lot),"))
	assert.Equal(t, "Occitanie", address.State)
	assert.Equal(t, "Lot", address.County)
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
	responseStreet = `{
    	"type": "FeatureCollection",
    	"version": "draft",
    	"features": [{
            "type": "Feature",
            "geometry": {
                "type": "Point",
                "coordinates": [3.156717,50.720114]
            },
            "properties": {
                "label": "Rue des Anges 59200 Tourcoing",
                "score": 0.7319487603305784,
                "id": "59599_0310",
                "name": "Rue des Anges",
                "postcode": "59200",
                "citycode": "59599",
                "x": 711087.23,
                "y": 7069277.25,
                "city": "Tourcoing",
                "context": "59, Nord, Hauts-de-France",
                "type": "street",
                "importance": 0.6878
            }
        }],
    	"attribution": "BAN",
    	"licence": "ETALAB-2.0",
    	"query": "11B Rue des Anges Tourcoing 59200",
    	"limit": 1
	}`
	responseLocality = `{
    	"type": "FeatureCollection",
    	"version": "draft",
    	"features": [{
            "type": "Feature",
            "geometry": {
                "type": "Point",
                "coordinates": [1.646584,44.995637]
            },
            "properties": {
                "label": "Route de Saint Denis les Martel (Les Quatre-Routes-du-Lot) 46110 Le Vignon-en-Quercy",
                "score": 0.8454463636363636,
                "type": "locality",
                "importance": 0.29991,
                "id": "46232_0023",
                "name": "Route de Saint Denis les Martel (Les Quatre-Routes-du-Lot)",
                "postcode": "46110",
                "citycode": "46232",
                "oldcitycode": "46232",
                "x": 593348.58,
                "y": 6433848.43,
                "city": "Le Vignon-en-Quercy",
                "oldcity": "Les Quatre-Routes-du-Lot",
                "context": "46, Lot, Occitanie"
            }
        }],
    	"attribution": "BAN",
    	"licence": "ETALAB-2.0",
    	"query": "Route de Saint-Denis-lès-Martel",
    	"limit": 1
	}`
)
