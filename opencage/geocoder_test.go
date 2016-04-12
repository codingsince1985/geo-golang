package opencage_test

import (
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/opencage"
	"math"
	"os"
	"strings"
	"testing"
)

var key = os.Getenv("OPENCAGE_API_KEY")

var geocoder = opencage.Geocoder(key)

// InDelta asserts that the two numerals are within delta of each other.
// From: https://github.com/stretchr/testify/blob/master/assert/assertions.go#L723
//
// 	 assert.InDelta(t, math.Pi, (22 / 7.0), 0.01)
//
// Returns whether the assertion was successful (true) or not (false).
func InDelta(expected, actual float64, delta float64) bool {

	if math.IsNaN(actual) {
		return false
	}

	if math.IsNaN(expected) {
		return false
	}

	dt := actual - expected
	if dt < -delta || dt > delta {
		return false
	}

	return true
}

// locDelta is the lat long precision significant digits we care about
// because the opencage API isn't the same exact lat,long as other systems
var locDelta = 0.00001

func TestGeocode(t *testing.T) {
	location, err := geocoder.Geocode("Melbourne VIC")
	if err != nil || !InDelta(-37.8142176, location.Lat, locDelta) || !InDelta(144.9631608, location.Lng, locDelta) {
		t.Error("TestGeocode() failed", err, location)
	}
}

func TestReverseGeocode(t *testing.T) {
	if address, err := geocoder.ReverseGeocode(-37.816742, 144.964463); err != nil || !strings.HasSuffix(address, "Melbourne VIC 3000, Australia") {
		t.Error("TestReverseGeocode() failed", err, address)
	}
}

func TestReverseGeocodeWithNoResult(t *testing.T) {
	if _, err := geocoder.ReverseGeocode(-37.816742, 164.964463); err != geo.ErrNoResult {
		t.Error("TestReverseGeocodeWithNoResult() failed", err)
	}
}
