package mapzen

import (
	"math"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	geo "github.com/codingsince1985/geo-golang"
)

var key = os.Getenv("MAPZEN_API_KEY")

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

	house := "1109"
	code := "USA"
	state := "Virginia"
	stateCode := "VA"
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

	if addr.HouseNumber != house {
		t.Fatalf("Got: %v\tExpected: %v\n", addr.HouseNumber, house)
	}
	if addr.CountryCode != code {
		t.Fatalf("Got: %v\tExpected: %v\n", addr.CountryCode, code)
	}
	if addr.State != state {
		t.Fatalf("Got: %v\tExpected: %v\n", addr.State, state)
	}
	if addr.StateCode != stateCode {
		t.Fatalf("Got: %v\tExpected: %v\n", addr.StateCode, stateCode)
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
    "geocoding": {
        "version": "0.2",
        "attribution": "https://search.mapzen.com/v1/attribution",
        "query": {
            "text": "1109 N Highland St, Arlington VA",
            "parsed_text": {
                "number": "1109",
                "street": "n highland st",
                "city": "arlington",
                "state": "va"
            },
            "size": 1,
            "private": false,
            "lang": {
                "name": "English",
                "iso6391": "en",
                "iso6393": "eng",
                "defaulted": true
            },
            "querySize": 20
        },
        "engine": {
            "name": "Pelias",
            "author": "Mapzen",
            "version": "1.0"
        },
        "timestamp": 1499034434562
    },
    "type": "FeatureCollection",
    "features": [
        {
            "type": "Feature",
            "geometry": {
                "type": "Point",
                "coordinates": [
                    -77.094733,
                    38.886665
                ]
            },
            "properties": {
                "id": "us/va/statewide:7120c09d09e25349",
                "gid": "openaddresses:address:us/va/statewide:7120c09d09e25349",
                "layer": "address",
                "source": "openaddresses",
                "source_id": "us/va/statewide:7120c09d09e25349",
                "name": "1109 N Highland St",
                "housenumber": "1109",
                "street": "N Highland St",
                "postalcode": "22201",
                "confidence": 1,
                "match_type": "exact",
                "accuracy": "point",
                "country": "United States",
                "country_gid": "whosonfirst:country:85633793",
                "country_a": "USA",
                "region": "Virginia",
                "region_gid": "whosonfirst:region:85688747",
                "region_a": "VA",
                "county": "Arlington County",
                "county_gid": "whosonfirst:county:102085953",
                "locality": "Arlington",
                "locality_gid": "whosonfirst:locality:101729469",
                "neighbourhood": "Courthouse",
                "neighbourhood_gid": "whosonfirst:neighbourhood:85811147",
                "label": "1109 N Highland St, Arlington, VA, USA"
            }
        }
    ],
    "bbox": [
        -77.094733,
        38.886665,
        -77.094733,
        38.886665
    ]
}
	`

	reverseResp = `
{
    "geocoding": {
        "version": "0.2",
        "attribution": "https://search.mapzen.com/v1/attribution",
        "query": {
            "size": 1,
            "private": false,
            "point.lat": 38.886665,
            "point.lon": -77.094733,
            "boundary.circle.radius": 1,
            "boundary.circle.lat": 38.886665,
            "boundary.circle.lon": -77.094733,
            "lang": {
                "name": "English",
                "iso6391": "en",
                "iso6393": "eng",
                "defaulted": true
            },
            "querySize": 20
        },
        "engine": {
            "name": "Pelias",
            "author": "Mapzen",
            "version": "1.0"
        },
        "timestamp": 1499034728798
    },
    "type": "FeatureCollection",
    "features": [
        {
            "type": "Feature",
            "geometry": {
                "type": "Point",
                "coordinates": [
                    -77.094733,
                    38.886665
                ]
            },
            "properties": {
                "id": "us/va/statewide:7120c09d09e25349",
                "gid": "openaddresses:address:us/va/statewide:7120c09d09e25349",
                "layer": "address",
                "source": "openaddresses",
                "source_id": "us/va/statewide:7120c09d09e25349",
                "name": "1109 N Highland St",
                "housenumber": "1109",
                "street": "N Highland St",
                "postalcode": "22201",
                "confidence": 1,
                "distance": 0,
                "accuracy": "point",
                "country": "United States",
                "country_gid": "whosonfirst:country:85633793",
                "country_a": "USA",
                "region": "Virginia",
                "region_gid": "whosonfirst:region:85688747",
                "region_a": "VA",
                "county": "Arlington County",
                "county_gid": "whosonfirst:county:102085953",
                "locality": "Arlington",
                "locality_gid": "whosonfirst:locality:101729469",
                "neighbourhood": "Courthouse",
                "neighbourhood_gid": "whosonfirst:neighbourhood:85811147",
                "label": "1109 N Highland St, Arlington, VA, USA"
            }
        }
    ],
    "bbox": [
        -77.094733,
        38.886665,
        -77.094733,
        38.886665
    ]
}`
)
