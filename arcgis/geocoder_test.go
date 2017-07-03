package arcgis

import (
	"math"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	geo "github.com/codingsince1985/geo-golang"
)

var token = os.Getenv("ARCGIS_TOKEN")

func TestGeocode(t *testing.T) {
	ts := testServer(geocodeResp)
	defer ts.Close()

	address := "380 New York, Redlands, CA 92373, USA"
	geocoder := Geocoder(token, ts.URL)
	loc, err := geocoder.Geocode(address)
	if err != nil {
		t.Fatal(err)
	}

	if loc == nil {
		t.Fatalf("Address: %s - Not Found\n", address)
	}
	expected := geo.Location{Lng: -117.1956703176181, Lat: 34.056488119308924}
	if math.Abs(loc.Lng-expected.Lng) > eps {
		t.Fatalf("Got: %v\tExpected: %v\n", loc, expected)
	}
	if math.Abs(loc.Lat-expected.Lat) > eps {
		t.Fatalf("Got: %v\tExpected: %v\n", loc, expected)
	}
}

func TestReverseGeocode(t *testing.T) {
	ts := testServer(reverseResp)
	defer ts.Close()

	code := "USA"
	state := "California"
	lat := 34.056488119308924
	lng := -117.1956703176181
	geocoder := Geocoder(token, ts.URL)
	addr, err := geocoder.ReverseGeocode(lat, lng)
	if err != nil {
		t.Error(err)
	}

	if addr == nil {
		t.Fatalf("Location: lat:%f, lng:%f - Not Found\n", lat, lng)
	}

	if addr.CountryCode != code {
		t.Fatalf("Got: %v\tExpected: %v\n", addr.CountryCode, code)
	}
	if addr.State != state {
		t.Fatalf("Got: %v\tExpected: %v\n", addr.State, state)
	}
}

func testServer(response string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte(response))
	}))
}

const (
	eps = 1.0e-5

	geocodeResp = `{
 "spatialReference": {
  "wkid": 4326,
  "latestWkid": 4326
 },
 "candidates": [
  {
   "address": "380 New York St, Redlands, California, 92373",
   "location": {
    "x": -117.1956703176181,
    "y": 34.056488119308924
   },
   "score": 100,
   "attributes": {

   },
   "extent": {
    "xmin": -117.1963135,
    "ymin": 34.055108000000011,
    "xmax": -117.19431349999999,
    "ymax": 34.057108000000007
   }
  }
 ]
}`

	reverseResp = `{
 "address": {
  "Match_addr": "117 Norwood St, Redlands, California, 92373",
  "LongLabel": "117 Norwood St, Redlands, CA, 92373, USA",
  "ShortLabel": "117 Norwood St",
  "Addr_type": "PointAddress",
  "Type": "",
  "PlaceName": "",
  "AddNum": "117",
  "Address": "117 Norwood St",
  "Block": "",
  "Sector": "",
  "Neighborhood": "South Redlands",
  "District": "",
  "City": "Redlands",
  "MetroArea": "Inland Empire",
  "Subregion": "San Bernardino",
  "Region": "California",
  "Territory": "",
  "Postal": "92373",
  "PostalExt": "",
  "CountryCode": "USA"
 },
 "location": {
  "x": -117.19052869813629,
  "y": 34.049723195135194,
  "spatialReference": {
   "wkid": 4326,
   "latestWkid": 4326
  }
 }
}`
)
