package mapbox_test

import (
	"fmt"
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/mapbox"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var token = os.Getenv("MAPBOX_API_KEY")

func TestGeocode(t *testing.T) {
	ts := testServer(response1)
	defer ts.Close()

	geocoder := mapbox.Geocoder(token, ts.URL+"/")
	location, err := geocoder.Geocode("60 Collins St, Melbourne VIC 3000")
	assert.NoError(t, err)
	assert.Equal(t, geo.Location{Lat: -37.813754, Lng: 144.971756}, location)
}

func TestReverseGeocode(t *testing.T) {
	ts := testServer(response2)
	defer ts.Close()

	geocoder := mapbox.Geocoder(token, ts.URL+"/")
	address, err := geocoder.ReverseGeocode(-37.813754, 144.971756)
	assert.NoError(t, err)
	fmt.Println(address)
	assert.True(t, strings.Index(address, "60 Collins St") >= 0)
}

func TestReverseGeocodeWithNoResult(t *testing.T) {
	ts := testServer(response3)
	defer ts.Close()

	geocoder := mapbox.Geocoder(token, ts.URL+"/")
	_, err := geocoder.ReverseGeocode(-37.813754, 164.971756)
	assert.Equal(t, err, geo.ErrNoResult)
}

func testServer(response string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte(response))
	}))
}

const (
	response1 = `{
   "type":"FeatureCollection",
   "query":[
      "60",
      "collins",
      "st",
      "melbourne",
      "vic",
      "3000"
   ],
   "features":[
      {
         "id":"address.10543198797834830",
         "type":"Feature",
         "text":"COLLINS STREET",
         "place_name":"60 COLLINS STREET, Melbourne, Victoria 3000, Australia",
         "relevance":0.822,
         "properties":{

         },
         "center":[
            144.971756,
            -37.813754
         ],
         "geometry":{
            "type":"Point",
            "coordinates":[
               144.971756,
               -37.813754
            ]
         },
         "address":"60",
         "context":[
            {
               "id":"locality.5321754973111320",
               "text":"Melbourne"
            },
            {
               "id":"place.1007",
               "text":"Melbourne"
            },
            {
               "id":"postcode.898",
               "text":"3000"
            },
            {
               "id":"region.4247793028417240",
               "text":"Victoria",
               "short_code":"AU-VIC",
               "wikidata":"Q36687"
            },
            {
               "id":"country.8513529047581710",
               "text":"Australia",
               "short_code":"au",
               "wikidata":"Q408"
            }
         ]
      },
      {
         "id":"address.2932253700834830",
         "type":"Feature",
         "text":"COLLINS STREET",
         "place_name":"60 COLLINS STREET, Thornbury, Victoria 3071, Australia",
         "relevance":0.746,
         "properties":{

         },
         "center":[
            145.006522,
            -37.75522
         ],
         "geometry":{
            "type":"Point",
            "coordinates":[
               145.006522,
               -37.75522
            ]
         },
         "address":"60",
         "context":[
            {
               "id":"locality.4441790122623020",
               "text":"Thornbury"
            },
            {
               "id":"place.1007",
               "text":"Melbourne"
            },
            {
               "id":"postcode.963",
               "text":"3071"
            },
            {
               "id":"region.4247793028417240",
               "text":"Victoria",
               "short_code":"AU-VIC",
               "wikidata":"Q36687"
            },
            {
               "id":"country.8513529047581710",
               "text":"Australia",
               "short_code":"au",
               "wikidata":"Q408"
            }
         ]
      },
      {
         "id":"address.9774289738834830",
         "type":"Feature",
         "text":"COLLINS STREET",
         "place_name":"60 COLLINS STREET, Mentone, Victoria 3194, Australia",
         "relevance":0.746,
         "properties":{

         },
         "center":[
            145.060207,
            -37.979776
         ],
         "geometry":{
            "type":"Point",
            "coordinates":[
               145.060207,
               -37.979776
            ]
         },
         "address":"60",
         "context":[
            {
               "id":"locality.10225647547503770",
               "text":"Mentone"
            },
            {
               "id":"place.1007",
               "text":"Melbourne"
            },
            {
               "id":"postcode.1116",
               "text":"3194"
            },
            {
               "id":"region.4247793028417240",
               "text":"Victoria",
               "short_code":"AU-VIC",
               "wikidata":"Q36687"
            },
            {
               "id":"country.8513529047581710",
               "text":"Australia",
               "short_code":"au",
               "wikidata":"Q408"
            }
         ]
      },
      {
         "id":"address.4503958544834830",
         "type":"Feature",
         "text":"COLLINS STREET",
         "place_name":"60 COLLINS STREET, St Albans, Victoria 3021, Australia",
         "relevance":0.746,
         "properties":{

         },
         "center":[
            144.803097,
            -37.738092
         ],
         "geometry":{
            "type":"Point",
            "coordinates":[
               144.803097,
               -37.738092
            ]
         },
         "address":"60",
         "context":[
            {
               "id":"locality.14089594205891930",
               "text":"St Albans"
            },
            {
               "id":"place.1007",
               "text":"Melbourne"
            },
            {
               "id":"postcode.915",
               "text":"3021"
            },
            {
               "id":"region.4247793028417240",
               "text":"Victoria",
               "short_code":"AU-VIC",
               "wikidata":"Q36687"
            },
            {
               "id":"country.8513529047581710",
               "text":"Australia",
               "short_code":"au",
               "wikidata":"Q408"
            }
         ]
      },
      {
         "id":"address.18053312966457090",
         "type":"Feature",
         "text":"Collins Street",
         "place_name":"60 Collins Street, Thornbury, Victoria 3071, Australia",
         "relevance":0.696,
         "properties":{

         },
         "center":[
            145.006223,
            -37.755052
         ],
         "geometry":{
            "type":"Point",
            "coordinates":[
               145.006223,
               -37.755052
            ],
            "interpolated":true
         },
         "address":"60",
         "context":[
            {
               "id":"locality.4441790122623020",
               "text":"Thornbury"
            },
            {
               "id":"place.1007",
               "text":"Melbourne"
            },
            {
               "id":"postcode.963",
               "text":"3071"
            },
            {
               "id":"region.4247793028417240",
               "text":"Victoria",
               "short_code":"AU-VIC",
               "wikidata":"Q36687"
            },
            {
               "id":"country.8513529047581710",
               "text":"Australia",
               "short_code":"au",
               "wikidata":"Q408"
            }
         ]
      }
   ],
   "attribution":"NOTICE: © 2016 Mapbox and its suppliers. All rights reserved. Use of this data is subject to the Mapbox Terms of Service (https://www.mapbox.com/about/maps/). This response and the information it contains may not be retained."
}`
	response2 = `{
   "type":"FeatureCollection",
   "query":[
      144.971756,
      -37.813754
   ],
   "features":[
      {
         "id":"poi.14351227508208000",
         "type":"Feature",
         "text":"Nak Cosmo",
         "place_name":"Nak Cosmo, 60 Collins St, 3000 Melbourne, Australia",
         "relevance":1,
         "properties":{
            "tel":"(03) 9654 6587",
            "address":"60 Collins St",
            "category":null
         },
         "center":[
            144.971756,
            -37.813754
         ],
         "geometry":{
            "type":"Point",
            "coordinates":[
               144.971756,
               -37.813754
            ]
         },
         "context":[
            {
               "id":"locality.5321754973111320",
               "text":"Melbourne"
            },
            {
               "id":"place.1007",
               "text":"Melbourne"
            },
            {
               "id":"postcode.898",
               "text":"3000"
            },
            {
               "id":"region.4247793028417240",
               "text":"Victoria",
               "short_code":"AU-VIC",
               "wikidata":"Q36687"
            },
            {
               "id":"country.8513529047581710",
               "text":"Australia",
               "wikidata":"Q408",
               "short_code":"au"
            }
         ]
      },
      {
         "id":"locality.5321754973111320",
         "type":"Feature",
         "text":"Melbourne",
         "place_name":"Melbourne, Melbourne, Victoria, Australia",
         "relevance":1,
         "properties":{

         },
         "bbox":[
            144.95143505,
            -37.855527,
            144.98909711,
            -37.799446
         ],
         "center":[
            144.9632,
            -37.8142
         ],
         "geometry":{
            "type":"Point",
            "coordinates":[
               144.9632,
               -37.8142
            ]
         },
         "context":[
            {
               "id":"place.1007",
               "text":"Melbourne"
            },
            {
               "id":"postcode.898",
               "text":"3000"
            },
            {
               "id":"region.4247793028417240",
               "text":"Victoria",
               "short_code":"AU-VIC",
               "wikidata":"Q36687"
            },
            {
               "id":"country.8513529047581710",
               "text":"Australia",
               "wikidata":"Q408",
               "short_code":"au"
            }
         ]
      },
      {
         "id":"place.1007",
         "type":"Feature",
         "text":"Melbourne",
         "place_name":"Melbourne, Victoria, Australia",
         "relevance":1,
         "properties":{

         },
         "bbox":[
            144.593741856,
            -38.433859306,
            145.512528832,
            -37.5112737225
         ],
         "center":[
            144.9632,
            -37.8142
         ],
         "geometry":{
            "type":"Point",
            "coordinates":[
               144.9632,
               -37.8142
            ]
         },
         "context":[
            {
               "id":"postcode.898",
               "text":"3000"
            },
            {
               "id":"region.4247793028417240",
               "text":"Victoria",
               "short_code":"AU-VIC",
               "wikidata":"Q36687"
            },
            {
               "id":"country.8513529047581710",
               "text":"Australia",
               "wikidata":"Q408",
               "short_code":"au"
            }
         ]
      },
      {
         "id":"postcode.898",
         "type":"Feature",
         "text":"3000",
         "place_name":"3000, Victoria, Australia",
         "relevance":1,
         "properties":{

         },
         "bbox":[
            144.9514229,
            -37.8297445,
            144.9890618,
            -37.7994973
         ],
         "center":[
            144.96317,
            -37.814264
         ],
         "geometry":{
            "type":"Point",
            "coordinates":[
               144.96317,
               -37.814264
            ]
         },
         "context":[
            {
               "id":"region.4247793028417240",
               "text":"Victoria",
               "short_code":"AU-VIC",
               "wikidata":"Q36687"
            },
            {
               "id":"country.8513529047581710",
               "text":"Australia",
               "wikidata":"Q408",
               "short_code":"au"
            }
         ]
      },
      {
         "id":"region.4247793028417240",
         "type":"Feature",
         "text":"Victoria",
         "place_name":"Victoria, Australia",
         "relevance":1,
         "properties":{
            "short_code":"AU-VIC",
            "wikidata":"Q36687"
         },
         "bbox":[
            140.948118489,
            -39.258117163,
            150.055423988,
            -33.9806475865
         ],
         "center":[
            144.549796,
            -36.559729
         ],
         "geometry":{
            "type":"Point",
            "coordinates":[
               144.549796,
               -36.559729
            ]
         },
         "context":[
            {
               "id":"country.8513529047581710",
               "text":"Australia",
               "wikidata":"Q408",
               "short_code":"au"
            }
         ]
      },
      {
         "id":"country.8513529047581710",
         "type":"Feature",
         "text":"Australia",
         "place_name":"Australia",
         "relevance":1,
         "properties":{
            "wikidata":"Q408",
            "short_code":"au"
         },
         "bbox":[
            112.821339418004,
            -43.84028825,
            159.209167131,
            -9.04271917700223
         ],
         "center":[
            134.489563,
            -25.734968
         ],
         "geometry":{
            "type":"Point",
            "coordinates":[
               134.489563,
               -25.734968
            ]
         }
      }
   ],
   "attribution":"NOTICE: © 2016 Mapbox and its suppliers. All rights reserved. Use of this data is subject to the Mapbox Terms of Service (https://www.mapbox.com/about/maps/). This response and the information it contains may not be retained."
}`
	response3 = `{
   "type":"FeatureCollection",
   "query":[
      164.971756,
      -37.813754
   ],
   "features":[

   ],
   "attribution":"NOTICE: © 2016 Mapbox and its suppliers. All rights reserved. Use of this data is subject to the Mapbox Terms of Service (https://www.mapbox.com/about/maps/). This response and the information it contains may not be retained."
}`
)
