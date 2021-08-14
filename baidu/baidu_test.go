package baidu_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/baidu"
	"github.com/stretchr/testify/assert"
)

var key = os.Getenv("BAIDU_APP_KEY")

func TestGeocode(t *testing.T) {
	ts := testServer(response1)
	defer ts.Close()

	geocoder := baidu.Geocoder(key, "en", "bd09ll", ts.URL+"/")
	location, err := geocoder.Geocode("60 Collins St, Melbourne VIC")

	assert.NoError(t, err)
	assert.Equal(t, geo.Location{Lat: 40.05703033345938, Lng: 116.3084202915042}, *location)
}

func TestReverseGeocode(t *testing.T) {
	ts := testServer(response2)
	defer ts.Close()

	geocoder := baidu.Geocoder(key, "en", "bd09ll", ts.URL+"/")
	address, err := geocoder.ReverseGeocode(40.03333340036988, 116.29999999999993)

	assert.NoError(t, err)
	assert.True(t, strings.HasPrefix(address.FormattedAddress, "43号 农大南路, Haidian, Beijing, China"))
	assert.True(t, address.City == "Beijing")
}

func TestReverseGeocodeWithNoResult(t *testing.T) {
	ts := testServer(response3)
	defer ts.Close()

	geocoder := baidu.Geocoder(key, "en", "bd09ll", ts.URL+"/")
	addr, err := geocoder.ReverseGeocode(-37.81375, 164.97176)

	assert.NoError(t, err)
	assert.Nil(t, addr)
}

func testServer(response string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte(response))
	}))
}

const (
	response1 = `{"status":0,"result":{"location":{"lng":116.3084202915042,"lat":40.05703033345938},"precise":1,"confidence":80,"comprehension":100,"level":"门址"}}`
	response2 = `{"status":0,"result":{"location":{"lng":116.29999999999993,"lat":40.03333340036988},"formatted_address":"43号 农大南路, Haidian, Beijing, China","business":"马连洼,上地","addressComponent":{"country":"China","country_code":0,"country_code_iso":"CHN","country_code_iso2":"CN","province":"Beijing","city":"Beijing","city_level":2,"district":"Haidian","town":"","town_code":"","adcode":"110108","street":"农大南路","street_number":"43号","direction":"附近","distance":"26"},"pois":[],"roads":[],"poiRegions":[],"sematic_description":"","cityCode":131}}`
	response3 = `{"status":0,"result":{"location":{"lng":164.97175999999986,"lat":-37.81375002268602},"formatted_address":"","business":"","addressComponent":{"country":"","country_code":-1,"country_code_iso":"","country_code_iso2":"","province":"","city":"","city_level":2,"district":"","town":"","town_code":"","adcode":"0","street":"","street_number":"","direction":"","distance":""},"pois":[],"roads":[],"poiRegions":[],"sematic_description":"","cityCode":0}}`
)
