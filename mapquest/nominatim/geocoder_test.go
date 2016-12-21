package nominatim_test

import (
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/mapquest/nominatim"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var key = os.Getenv("MAPQUEST_NOMINATUM_KEY")

func TestGeocode(t *testing.T) {
	ts := testServer(response1)
	defer ts.Close()

	geocoder := nominatim.Geocoder(key, ts.URL+"/")
	location, err := geocoder.Geocode("60 Collins St, Melbourne VIC 3000")
	assert.NoError(t, err)
	assert.Equal(t, geo.Location{Lat: -37.8137433689794, Lng: 144.971745104488}, location)
}

func TestReverseGeocode(t *testing.T) {
	ts := testServer(response2)
	defer ts.Close()

	geocoder := nominatim.Geocoder(key, ts.URL+"/")
	address, err := geocoder.ReverseGeocode(-37.8137433689794, 144.971745104488)
	assert.NoError(t, err)
	assert.True(t, strings.HasPrefix(address, "Reserve Bank of Australia"))
}

func TestReverseGeocodeWithNoResult(t *testing.T) {
	ts := testServer(response3)
	defer ts.Close()

	geocoder := nominatim.Geocoder(key, ts.URL+"/")
	_, err := geocoder.ReverseGeocode(-37.8137433689794, 164.971745104488)
	assert.Equal(t, err, geo.ErrNoResult)
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
)
