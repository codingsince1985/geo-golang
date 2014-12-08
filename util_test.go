package geo_test

import (
	"fmt"
	"github.com/codingsince1985/geo-golang"
	"testing"
)

func TestGetResponseData(t *testing.T) {
	data := geo.ResponseData("http://maps.googleapis.com/maps/api/geocode/json?sensor=false&address=Melbourne%20VIC")
	if data == nil {
		t.Error("ResponseData() failed")
	}
	fmt.Println(string(data))
}
