package geocod

import (
	"math"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	geo "github.com/codingsince1985/geo-golang"
)

var key = os.Getenv("GECOD_API_KEY")

func TestGeocode(t *testing.T) {
	ts := testServer(geocodeResp)
	defer ts.Close()

	address := "1109 N Highland St, Arlington VA"
	geocoder := Geocoder(key, ts.URL)
	loc, err := geocoder.Geocode(address)
	if err != nil {
		t.Fatal(err)
	}

	if loc == nil {
		t.Fatalf("Address: %s - Not Found\n", address)
	}
	expected := geo.Location{Lng: -77.094733, Lat: 38.886665}
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

	code := "US"
	state := "DC"
	lat := 38.886665
	lng := -77.094733
	geocoder := Geocoder(key, ts.URL)
	addr, err := geocoder.ReverseGeocode(lat, lng)
	if err != nil {
		t.Fatal(err)
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

	geocodeResp = `
{
  "input": {
    "address_components": {
      "number": "1109",
      "predirectional": "N",
      "street": "Highland",
      "suffix": "St",
      "formatted_street": "N Highland St",
      "city": "Arlington",
      "state": "VA",
      "zip": "22201",
      "country": "US"
    },
    "formatted_address": "1109 N Highland St, Arlington, VA 22201"
  },
  "results": [
    {
      "address_components": {
        "number": "1109",
        "predirectional": "N",
        "street": "Highland",
        "suffix": "St",
        "formatted_street": "N Highland St",
        "city": "Arlington",
        "county": "Arlington County",
        "state": "VA",
        "zip": "22201",
        "country": "US"
      },
      "formatted_address": "1109 N Highland St, Arlington, VA 22201",
      "location": {
        "lat": 38.886665,
        "lng": -77.094733
      },
      "accuracy": 1,
      "accuracy_type": "rooftop",
      "source": "Virginia GIS Clearinghouse"
    }
  ]
}`

	reverseResp = `
{
  "results": [
    {
      "address_components": {
        "number": "500",
        "street": "H",
        "suffix": "St",
        "postdirectional": "NE",
        "formatted_street": "H St NE",
        "city": "Washington",
        "county": "District of Columbia",
        "state": "DC",
        "zip": "20002",
        "country": "US"
      },
      "formatted_address": "500 H St NE, Washington, DC 20002",
      "location": {
        "lat": 38.900203,
        "lng": -76.999507
      },
      "accuracy": 1,
      "accuracy_type": "nearest_street",
      "source": "TIGER/Line® dataset from the US Census Bureau"
    },
    {
      "address_components": {
        "number": "800",
        "street": "5th",
        "suffix": "St",
        "postdirectional": "NE",
        "formatted_street": "5th St NE",
        "city": "Washington",
        "county": "District of Columbia",
        "state": "DC",
        "zip": "20002",
        "country": "US"
      },
      "formatted_address": "800 5th St NE, Washington, DC 20002",
      "location": {
        "lat": 38.900203,
        "lng": -76.999507
      },
      "accuracy": 0.18,
      "accuracy_type": "nearest_street",
      "source": "TIGER/Line® dataset from the US Census Bureau"
    },
    {
      "address_components": {
        "number": "474",
        "street": "H",
        "suffix": "St",
        "postdirectional": "NE",
        "formatted_street": "H St NE",
        "city": "Washington",
        "county": "District of Columbia",
        "state": "DC",
        "zip": "20002",
        "country": "US"
      },
      "formatted_address": "474 H St NE, Washington, DC 20002",
      "location": {
        "lat": 38.900205,
        "lng": -76.99994
      },
      "accuracy": 0.18,
      "accuracy_type": "nearest_street",
      "source": "TIGER/Line® dataset from the US Census Bureau"
    }
  ]
}`
)
