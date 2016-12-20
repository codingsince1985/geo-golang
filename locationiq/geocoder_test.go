package locationiq

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGeocodeYieldsResult(t *testing.T) {
	ts := testServer(responseForGeocode)
	defer ts.Close()

	gc := Geocoder("foobar", 18, ts.URL+"/")
	l, err := gc.Geocode("Seidlstraße 26, 80335 München")

	expLat := 48.1453641
	expLon := 11.5582083

	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	if expLat != l.Lat {
		t.Errorf("Expected latitude: %s, got %s", expLat, l.Lat)
	}

	if expLon != l.Lng {
		t.Errorf("Expected longitude %s, got %s", expLon, l.Lng)
	}
}

func TestGeocodeYieldsNoResult(t *testing.T) {
	ts := testServer("[]")
	defer ts.Close()

	gc := Geocoder("foobar", 18, ts.URL+"/")
	l, err := gc.Geocode("Seidlstraße 26, 80335 München")

	if err == nil {
		t.Error("Got nil error")
	}
	if l.Lat != 0 {
		t.Errorf("Expected latitude: %d, got: %f", 0, l.Lat)
	}
	if l.Lng != 0 {
		t.Errorf("Expected longitude: %d, got: %f", 0, l.Lat)
	}
}

func TestReverseGeocodeYieldsResult(t *testing.T) {
	ts := testServer(responseForReverse)
	defer ts.Close()

	gc := Geocoder("foobar", 18, ts.URL+"/")
	addr, err := gc.ReverseGeocode(48.1453641, 11.5582083)

	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
	if !strings.HasPrefix(addr, "26, Seidlstraße") {
		t.Errorf("Expected address string starting with %s, got string: %s", "26, Seidlstraße", addr)
	}
}

func TestReverseGeocodeYieldsNoResult(t *testing.T) {
	ts := testServer(errorResponse)
	defer ts.Close()

	gc := Geocoder("foobar", 18, ts.URL+"/")
	addr, err := gc.ReverseGeocode(48.1453641, 11.5582083)

	if err == nil {
		t.Error("Expected error, got nil")
	}
	if addr != "" {
		t.Errorf("Expected empty string as address, got: %s", addr)
	}
}

func testServer(response string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte(response))
	}))
}

const (
	responseForGeocode = `[
  {
    "place_id": "25798174",
    "licence": "Data © OpenStreetMap contributors, ODbL 1.0. http://www.openstreetmap.org/copyright",
    "osm_type": "node",
    "osm_id": "2475749822",
    "boundingbox": [
      "48.1453141",
      "48.1454141",
      "11.5581583",
      "11.5582583"
    ],
    "lat": "48.1453641",
    "lon": "11.5582083",
    "display_name": "26, Seidlstraße, Bezirksteil Augustenstraße, Maxvorstadt, Munich, Upper Bavaria, Free State of Bavaria, 80335, Germany",
    "class": "place",
    "type": "house",
    "importance": 0.111
  }
]`
	responseForReverse = `{
  "place_id": "25798174",
  "licence": "Data © OpenStreetMap contributors, ODbL 1.0. http://www.openstreetmap.org/copyright",
  "osm_type": "node",
  "osm_id": "2475749822",
  "lat": "48.1453641",
  "lon": "11.5582083",
  "display_name": "26, Seidlstraße, Bezirksteil Augustenstraße, Maxvorstadt, Munich, Upper Bavaria, Free State of Bavaria, 80335, Germany",
  "address": {
    "house_number": "26",
    "road": "Seidlstraße",
    "suburb": "Maxvorstadt",
    "city": "Munich",
    "state_district": "Upper Bavaria",
    "state": "Free State of Bavaria",
    "postcode": "80335",
    "country": "Germany",
    "country_code": "de"
  },
  "boundingbox": [
    "48.1452641",
    "48.1454641",
    "11.5581083",
    "11.5583083"
  ]
}`
	errorResponse = `{
  "error": "Unable to geocode"
}`
)
