package amap_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/amap"
	"github.com/stretchr/testify/assert"
)

var key = os.Getenv("AMAP_APP_KEY")

func TestGeocode(t *testing.T) {
	ts := testServer(response1)
	defer ts.Close()

	geocoder := amap.Geocoder(key, 1000, ts.URL+"/")
	location, err := geocoder.Geocode("北京市海淀区清河街道西二旗西路领秀新硅谷")

	assert.NoError(t, err)
	assert.NotNil(t, location)
	if location != nil {
		assert.Equal(t, geo.Location{Lat: 40.055106, Lng: 116.309866}, *location)
	}
}

func TestReverseGeocode(t *testing.T) {
	ts := testServer(response2)
	defer ts.Close()

	geocoder := amap.Geocoder(key, 1000, ts.URL+"/")
	address, err := geocoder.ReverseGeocode(116.3084202915042, 116.3084202915042)

	assert.NoError(t, err)
	assert.Equal(t, address.FormattedAddress, "北京市海淀区上地街道树村郊野公园")
}

func TestReverseGeocodeWithNoResult(t *testing.T) {
	ts := testServer(response3)
	defer ts.Close()

	geocoder := amap.Geocoder(key, 1000, ts.URL+"/")
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
	response1 = `<?xml version="1.0" encoding="UTF-8"?><response><status>1</status><info>OK</info><infocode>10000</infocode><count>1</count><geocodes type="list"><geocode><formatted_address>北京市海淀区清河街道西二旗西路领秀新硅谷</formatted_address><country>中国</country><province>北京市</province><citycode>010</citycode><city>北京市</city><district>海淀区</district><township></township><neighborhood><name></name><type></type></neighborhood><building><name></name><type></type></building><adcode>110108</adcode><street>西二旗西路</street><number></number><location>116.309866,40.055106</location><level>住宅区</level></geocode></geocodes></response>`
	response2 = `<?xml version="1.0" encoding="UTF-8"?><response><status>1</status><info>OK</info><infocode>10000</infocode><regeocode><formatted_address>北京市海淀区上地街道树村郊野公园</formatted_address><addressComponent><country>中国</country><province>北京市</province><city></city><citycode>010</citycode><district>海淀区</district><adcode>110108</adcode><township>上地街道</township><towncode>110108022000</towncode><neighborhood><name></name><type></type></neighborhood><building><name>树村郊野公园</name><type>风景名胜;公园广场;公园</type></building><streetNumber><street>马连洼北路</street><number>29号</number><location>116.299587,40.034620</location><direction>北</direction><distance>147.306</distance></streetNumber><businessAreas type="list"><businessArea><location>116.303276,40.035542</location><name>上地</name><id>110108</id></businessArea><businessArea><location>116.256057,40.054273</location><name>西北   </name><id>110108</id></businessArea><businessArea><location>116.281156,40.028654</location><name>马连洼</name><id>110108</id></businessArea></businessAreas></addressComponent></regeocode></response>`
	response3 = `<?xml version="1.0" encoding="UTF-8"?><response><status>1</status><info>OK</info><infocode>10000</infocode><regeocode><formatted_address></formatted_address><addressComponent><country></country><province></province><city></city><citycode></citycode><district></district><adcode></adcode><township></township><towncode></towncode></addressComponent><pois type="list"/><roads type="list"/><roadinters type="list"/><aois type="list"/></regeocode></response>`
)
