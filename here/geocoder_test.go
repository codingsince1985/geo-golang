package here_test

import (
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/here"
	"strings"
	"testing"
)

const app_id = "YOUR_APP_ID"
const app_code = "YOUR_APP_CODE"

var geocoder = here.Geocoder(app_id, app_code, 100)

func TestGeocode(t *testing.T) {
	location, err := geocoder.Geocode("Melbourne VIC")
	if err != nil || location.Lat != -37.81753 || location.Lng != 144.96715 {
		t.Error("TestGeocode() failed", err, location)
	}
}

func TestReverseGeocode(t *testing.T) {
	address, err := geocoder.ReverseGeocode(-37.81753, 144.96715)
	if err != nil || !strings.HasSuffix(address, "VIC 3000, Australia") {
		t.Error("TestReverseGeocode() failed", err, address)
	}
}

func TestReverseGeocodeWithNoResult(t *testing.T) {
	_, err := geocoder.ReverseGeocode(-37.81753, 164.96715)
	if err != geo.NoResultError {
		t.Error("TestReverseGeocodeWithNoResult() failed", err)
	}
}
