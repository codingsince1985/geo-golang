package nominatim_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/mapquest/nominatim"
	"github.com/stretchr/testify/assert"
)

var key = os.Getenv("MAPQUEST_NOMINATUM_KEY")

func TestGeocode(t *testing.T) {
	ts := testServer(response1)
	defer ts.Close()

	geocoder := nominatim.Geocoder(key, ts.URL+"/")
	location, err := geocoder.Geocode("60 Collins St, Melbourne VIC 3000")
	assert.NoError(t, err)
	assert.Equal(t, geo.Location{Lat: -37.8137433689794, Lng: 144.971745104488}, *location)
}

func TestReverseGeocode(t *testing.T) {
	ts := testServer(response2)
	defer ts.Close()

	geocoder := nominatim.Geocoder(key, ts.URL+"/")
	addr, err := geocoder.ReverseGeocode(-37.8137433689794, 144.971745104488)
	assert.NoError(t, err)
	assert.True(t, strings.HasPrefix(addr.FormattedAddress, "Reserve Bank of Australia"))

	ts2 := testServer(response4)
	defer ts2.Close()

	geocoder = nominatim.Geocoder(key, ts2.URL+"/")
	addr, err = geocoder.ReverseGeocode(43.0280986, -78.8136961)
	assert.NoError(t, err)
	assert.Equal(t, "Audubon Industrial Park", addr.City)
}

func TestReverseGeocodeWithNoResult(t *testing.T) {
	ts := testServer(response3)
	defer ts.Close()

	geocoder := nominatim.Geocoder(key, ts.URL+"/")
	//geocoder := nominatim.Geocoder(key)
	addr, err := geocoder.ReverseGeocode(-37.8137433689794, 164.971745104488)
	assert.NotNil(t, err)
	assert.Nil(t, addr)
}

func testServer(response string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte(response))
	}))
}

const (
	response1 = `[
   {
      "place_id":"2685280165",
      "licence":"Data \u00a9 OpenStreetMap contributors, ODbL 1.0. http:\/\/www.openstreetmap.org\/copyright",
      "osm_type":"node",
      "osm_id":"1038589059",
      "boundingbox":[
         "-37.8137433689794",
         "-37.8137433689794",
         "144.971745104488",
         "144.971745104488"
      ],
      "lat":"-37.8137433689794",
      "lon":"144.971745104488",
      "display_name":"60, Collins Street, Melbourne, City of Melbourne, Greater Melbourne, Victoria, 3000, Australia",
      "class":"place",
      "type":"house",
      "importance":0.411
   }
]`
	response2 = `{
   "place_id":"70709718",
   "licence":"Data \u00a9 OpenStreetMap contributors, ODbL 1.0. http:\/\/www.openstreetmap.org\/copyright",
   "osm_type":"way",
   "osm_id":"52554876",
   "lat":"-37.81362105",
   "lon":"144.971609569165",
   "display_name":"Reserve Bank of Australia, Exhibition Street, Melbourne, City of Melbourne, Greater Melbourne, Victoria, 3000, Australia",
   "address":{
      "building":"Reserve Bank of Australia",
      "road":"Exhibition Street",
      "suburb":"Melbourne",
      "city":"Melbourne",
      "county":"City of Melbourne",
      "region":"Greater Melbourne",
      "state":"Victoria",
      "postcode":"3000",
      "country":"Australia",
      "country_code":"au"
   }
}`
	response3 = `{
   "error":"Unable to geocode"
}`
	response4 = `{
    "place_id": 259174465,
    "licence": "Data \u00a9 OpenStreetMap contributors, ODbL 1.0. https:\/\/osm.org\/copyright",
    "osm_type": "way",
    "osm_id": 672030029,
    "lat": "43.0275119527027",
    "lon": "-78.81372135810811",
    "display_name": "188, Commerce Drive, Audubon Industrial Park, Amherst Town, Erie County, New York, 14228, United States of America",
    "address": {
        "house_number": "188",
        "road": "Commerce Drive",
        "hamlet": "Audubon Industrial Park",
        "county": "Erie County",
        "state": "New York",
        "postcode": "14228",
        "country": "United States of America",
        "country_code": "us"
    },
    "boundingbox": [
        "43.027411952703",
        "43.027611952703",
        "-78.813821358108",
        "-78.813621358108"
    ]
}`
)
