package geo

import (
	"testing"
)

func TestGetResponseData(t *testing.T) {
	data := responseData("http://maps.googleapis.com/maps/api/geocode/json?sensor=false&address=Melbourne%20VIC")
	if data == nil {
		t.Error("responseData() failed")
	}
}
