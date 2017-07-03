package tomtom

import (
	"math"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	geo "github.com/codingsince1985/geo-golang"
)

var key = os.Getenv("TOMTOM_API_KEY")

func TestGeocode(t *testing.T) {
	ts := testServer(geocodeResp)
	defer ts.Close()

	address := "1109 N Highland St, Arlington VA, US"
	geocoder := Geocoder(key, ts.URL)
	loc, err := geocoder.Geocode(address)
	if err != nil {
		t.Fatal(err)
	}

	if loc == nil {
		t.Fatalf("Address: %s - Not Found\n", address)
	}
	expected := geo.Location{Lng: -77.09464, Lat: 38.88669}
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
	lat := 38.88669
	lng := -77.09464
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
    "summary": {
        "query": "1109 N Highland St, Arlington, VA 22201",
        "queryType": "NON_NEAR",
        "queryTime": 158,
        "numResults": 1,
        "offset": 0,
        "totalResults": 1,
        "fuzzyLevel": 1
    },
    "results": [
        {
            "type": "Point Address",
            "id": "US/PAD/p0/26924656",
            "score": 11.51,
            "address": {
                "streetNumber": "1109",
                "streetName": "N Highland St",
                "municipalitySubdivision": "Arlington, Clarendon Courthouse",
                "municipality": "Arlington",
                "countrySecondarySubdivision": "Arlington",
                "countryTertiarySubdivision": "Arlington",
                "countrySubdivision": "VA",
                "postalCode": "22201",
                "extendedPostalCode": "222012890",
                "countryCode": "US",
                "country": "United States Of America",
                "countryCodeISO3": "USA",
                "freeformAddress": "1109 N Highland St, Arlington, VA 222012890",
                "countrySubdivisionName": "Virginia"
            },
            "position": {
                "lat": 38.88669,
                "lon": -77.09464
            },
            "viewport": {
                "topLeftPoint": {
                    "lat": 38.88759,
                    "lon": -77.0958
                },
                "btmRightPoint": {
                    "lat": 38.88579,
                    "lon": -77.09348
                }
            },
            "entryPoints": [
                {
                    "type": "main",
                    "position": {
                        "lat": 38.88667,
                        "lon": -77.09488
                    }
                }
            ]
        }
    ]
}`

	reverseResp = `
{
    "summary": {
        "queryTime": 504,
        "numResults": 1
    },
    "addresses": [
        {
            "address": {
                "buildingNumber": "414",
                "streetNumber": "414",
                "routeNumbers": [],
                "street": "Seward Sq SE",
                "streetName": "Seward Sq SE",
                "streetNameAndNumber": "414 Seward Sq SE",
                "countryCode": "US",
                "countrySubdivision": "DC",
                "countrySecondarySubdivision": "District of Columbia",
                "countryTertiarySubdivision": "Washington",
                "municipality": "Washington",
                "postalCode": "20003",
                "country": "United States Of America",
                "countryCodeISO3": "USA",
                "freeformAddress": "414 Seward Sq SE, Washington, DC 20003",
                "countrySubdivisionName": "District of Columbia"
            },
            "position": "38.886540,-77.000001"
        }
    ]
}`
)
