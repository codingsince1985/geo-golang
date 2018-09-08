package pickpoint_test

import (
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/pickpoint"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var key = os.Getenv("PICKPOINT_API_KEY")

func TestGeocode(t *testing.T) {
	ts := testServer(response1)
	defer ts.Close()

	geocoder := pickpoint.Geocoder(key, ts.URL+"/")
	location, err := geocoder.Geocode("60 Collins St, Melbourne VIC 3000")
	assert.Nil(t, err)
	assert.Equal(t, geo.Location{Lat: -37.8157915, Lng: 144.9656171}, *location)
}

func TestReverseGeocode(t *testing.T) {
	ts := testServer(response2)
	defer ts.Close()

	geocoder := pickpoint.Geocoder(key, ts.URL+"/")
	address, err := geocoder.ReverseGeocode(-37.8157915, 144.9656171)
	assert.Nil(t, err)
	assert.True(t, strings.Index(address.FormattedAddress, "Collins St") > 0)
}

func TestReverseGeocodeWithNoResult(t *testing.T) {
	ts := testServer(response3)
	defer ts.Close()

	geocoder := pickpoint.Geocoder(key, ts.URL+"/")
	addr, err := geocoder.ReverseGeocode(-37.8157915, 164.9656171)
	assert.Nil(t, addr)
	assert.NotNil(t, err)
}

func testServer(response string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte(response))
	}))
}

const (
	response1 = `[
   {
      "place_id":"133372311",
      "licence":"Data © OpenStreetMap contributors, ODbL 1.0. http:\/\/www.openstreetmap.org\/copyright",
      "osm_type":"way",
      "osm_id":"316166613",
      "boundingbox":[
         "-37.8162553",
         "-37.815533",
         "144.9640149",
         "144.9665099"
      ],
      "lat":"-37.8157915",
      "lon":"144.9656171",
      "display_name":"Collins Street, Melbourne, City of Melbourne, Greater Melbourne, Victoria, 3000, Australia",
      "class":"highway",
      "type":"tertiary",
      "importance":0.51
   }
]`
	response2 = `{
   "place_id":"5122082",
   "licence":"Data © OpenStreetMap contributors, ODbL 1.0. http:\/\/www.openstreetmap.org\/copyright",
   "osm_type":"node",
   "osm_id":"594206614",
   "lat":"-37.8158091",
   "lon":"144.9656492",
   "display_name":"Telstra, Collins Street, Melbourne, City of Melbourne, Greater Melbourne, Victoria, 3000, Australia",
   "address":{
      "telephone":"Telstra",
      "road":"Collins Street",
      "suburb":"Melbourne",
      "city":"Melbourne",
      "county":"City of Melbourne",
      "region":"Greater Melbourne",
      "state":"Victoria",
      "postcode":"3000",
      "country":"Australia",
      "country_code":"au"
   },
   "boundingbox":[
      "-37.8159091",
      "-37.8157091",
      "144.9655492",
      "144.9657492"
   ]
}`
	response3 = `{
   "error":"Unable to geocode"
}`
)
