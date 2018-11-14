package opencage_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/codingsince1985/geo-golang/opencage"
	"github.com/stretchr/testify/assert"
)

var key = os.Getenv("OPENCAGE_API_KEY")

// locDelta is the lat long precision significant digits we care about
// because the opencage API isn't the same exact lat,long as other systems
var locDelta = 0.01

func TestGeocode(t *testing.T) {
	ts := testServer(response1)
	defer ts.Close()

	geocoder := opencage.Geocoder(key, ts.URL+"/")
	location, err := geocoder.Geocode("60 Collins St, Melbourne VIC 3000")
	assert.Nil(t, err)
	assert.InDelta(t, -37.8154176, location.Lat, locDelta)
	assert.InDelta(t, 144.9665563, location.Lng, locDelta)
}

func TestReverseGeocode(t *testing.T) {
	ts := testServer(response2)
	defer ts.Close()

	geocoder := opencage.Geocoder(key, ts.URL+"/")
	address, err := geocoder.ReverseGeocode(-37.8154176, 144.9665563)
	assert.NoError(t, err)
	assert.True(t, strings.Index(address.FormattedAddress, "Collins St") > 0)
}

func TestReverseGeocodeWithNoResult(t *testing.T) {
	ts := testServer(response3)
	defer ts.Close()

	geocoder := opencage.Geocoder(key, ts.URL+"/")
	address, err := geocoder.ReverseGeocode(-37.8154176, 164.9665563)
	assert.Nil(t, err)
	assert.Nil(t, address)
}

func TestReverseGeocodeUsingSuburbAsLocality(t *testing.T) {
	ts := testServer(responseMissingLocality)
	defer ts.Close()

	geocoder := opencage.Geocoder(key, ts.URL+"/")
	address, err := geocoder.ReverseGeocode(-37.8154176, 164.9665563)
	assert.Nil(t, err)
	assert.NotNil(t, address)
	assert.Equal(t, "Lütten Klein", address.City)
}

func testServer(response string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte(response))
	}))
}

const (
	response1 = `{
   "documentation":"https://geocoder.opencagedata.com/api",
   "licenses":[
      {
         "name":"CC-BY-SA",
         "url":"http://creativecommons.org/licenses/by-sa/3.0/"
      },
      {
         "name":"ODbL",
         "url":"http://opendatacommons.org/licenses/odbl/summary/"
      }
   ],
   "rate":{
      "limit":2500,
      "remaining":2383,
      "reset":1463184000
   },
   "results":[
      {
         "annotations":{
            "DMS":{
               "lat":"37\u00b0 48' 59.45148'' S",
               "lng":"144\u00b0 57' 47.23560'' E"
            },
            "MGRS":"55HCU2071312588",
            "Maidenhead":"QF22le54na",
            "Mercator":{
               "x":16137220.814,
               "y":-4527336.343
            },
            "OSM":{
               "edit_url":"https://www.openstreetmap.org/edit?way=215320282#map=17/-37.81651/144.96312",
               "url":"https://www.openstreetmap.org/?mlat=-37.81651&mlon=144.96312#map=17/-37.81651/144.96312"
            },
            "callingcode":61,
            "geohash":"r1r0fewyvdc3r425n4nx",
            "sun":{
               "rise":{
                  "apparent":1463173980,
                  "astronomical":1463168580,
                  "civil":1463172300,
                  "nautical":1463170440
               },
               "set":{
                  "apparent":1463124120,
                  "astronomical":1463129580,
                  "civil":1463125800,
                  "nautical":1463127720
               }
            },
            "timezone":{
               "name":"Australia/Melbourne",
               "now_in_dst":0,
               "offset_sec":36000,
               "offset_string":1000,
               "short_name":"AEST"
            },
            "what3words":{
               "words":"themes.fees.tolls"
            }
         },
         "bounds":{
            "northeast":{
               "lat":-37.8162553,
               "lng":144.9640149
            },
            "southwest":{
               "lat":-37.8169249,
               "lng":144.9617036
            }
         },
         "components":{
            "_type":"road",
            "city":"Melbourne",
            "country":"Australia",
            "country_code":"au",
            "county":"City of Melbourne",
            "postcode":"3000",
            "region":"Greater Melbourne",
            "road":"Collins Street",
            "state":"Victoria",
            "suburb":"Melbourne"
         },
         "confidence":10,
         "formatted":"Collins Street, Melbourne VIC 3000, Australia",
         "geometry":{
            "lat":-37.8165143,
            "lng":144.963121
         }
      },
      {
         "annotations":{
            "DMS":{
               "lat":"37\u00b0 48' 50.40000'' S",
               "lng":"144\u00b0 57' 47.88000'' E"
            },
            "MGRS":"55HCU2072312867",
            "Maidenhead":"QF22le54op",
            "Mercator":{
               "x":16137240.74,
               "y":-4526983.531
            },
            "OSM":{
               "url":"https://www.openstreetmap.org/?mlat=-37.81400&mlon=144.96330#map=17/-37.81400/144.96330"
            },
            "geohash":"r1r0fspj3xqzwneegpsn",
            "sun":{
               "rise":{
                  "apparent":1463173980,
                  "astronomical":1463168580,
                  "civil":1463172300,
                  "nautical":1463170440
               },
               "set":{
                  "apparent":1463124120,
                  "astronomical":1463129580,
                  "civil":1463125800,
                  "nautical":1463127720
               }
            },
            "timezone":{
               "name":"Australia/Melbourne",
               "now_in_dst":0,
               "offset_sec":36000,
               "offset_string":1000,
               "short_name":"AEST"
            },
            "what3words":{
               "words":"chat.seat.fortunate"
            }
         },
         "components":{
            "_type":"postcode",
            "country":"Australia",
            "postcode":"3000"
         },
         "confidence":10,
         "formatted":"3000, Australia",
         "geometry":{
            "lat":-37.814,
            "lng":144.9633
         }
      }
   ],
   "status":{
      "code":200,
      "message":"OK"
   },
   "stay_informed":{
      "blog":"http://blog.opencagedata.com",
      "twitter":"https://twitter.com/opencagedata"
   },
   "thanks":"For using an OpenCage Data API",
   "timestamp":{
      "created_http":"Fri, 13 May 2016 10:39:20 GMT",
      "created_unix":1463135960
   },
   "total_results":2
}`
	response2 = `{
   "documentation":"https://geocoder.opencagedata.com/api",
   "licenses":[
      {
         "name":"CC-BY-SA",
         "url":"http://creativecommons.org/licenses/by-sa/3.0/"
      },
      {
         "name":"ODbL",
         "url":"http://opendatacommons.org/licenses/odbl/summary/"
      }
   ],
   "rate":{
      "limit":2500,
      "remaining":2382,
      "reset":1463184000
   },
   "results":[
      {
         "annotations":{
            "DMS":{
               "lat":"37\u00b0 48' 55.11266'' S",
               "lng":"144\u00b0 57' 59.36892'' E"
            },
            "MGRS":"55HCU2100712728",
            "Maidenhead":"QF22le54xh",
            "Mercator":{
               "x":16137596.001,
               "y":-4527167.222
            },
            "OSM":{
               "edit_url":"https://www.openstreetmap.org/edit?way=87337974#map=17/-37.81531/144.96649",
               "url":"https://www.openstreetmap.org/?mlat=-37.81531&mlon=144.96649#map=17/-37.81531/144.96649"
            },
            "callingcode":61,
            "geohash":"r1r0fgcmu56wg15ek296",
            "sun":{
               "rise":{
                  "apparent":1463173980,
                  "astronomical":1463168580,
                  "civil":1463172300,
                  "nautical":1463170440
               },
               "set":{
                  "apparent":1463124120,
                  "astronomical":1463129580,
                  "civil":1463125800,
                  "nautical":1463127720
               }
            },
            "timezone":{
               "name":"Australia/Melbourne",
               "now_in_dst":0,
               "offset_sec":36000,
               "offset_string":1000,
               "short_name":"AEST"
            },
            "what3words":{
               "words":"clocks.beard.slices"
            }
         },
         "components":{
            "_type":"building",
            "city":"Melbourne",
            "country":"Australia",
            "country_code":"au",
            "county":"City of Melbourne",
            "house_number":"230",
            "postcode":"3000",
            "region":"Greater Melbourne",
            "road":"Collins Street",
            "state":"Victoria",
            "suburb":"Melbourne"
         },
         "confidence":10,
         "formatted":"230 Collins Street, Melbourne VIC 3000, Australia",
         "geometry":{
            "lat":-37.8153091,
            "lng":144.9664914
         }
      }
   ],
   "status":{
      "code":200,
      "message":"OK"
   },
   "stay_informed":{
      "blog":"http://blog.opencagedata.com",
      "twitter":"https://twitter.com/opencagedata"
   },
   "thanks":"For using an OpenCage Data API",
   "timestamp":{
      "created_http":"Fri, 13 May 2016 10:43:19 GMT",
      "created_unix":1463136199
   },
   "total_results":1
}`
	response3 = `{
   "documentation":"https://geocoder.opencagedata.com/api",
   "licenses":[
      {
         "name":"CC-BY-SA",
         "url":"http://creativecommons.org/licenses/by-sa/3.0/"
      },
      {
         "name":"ODbL",
         "url":"http://opendatacommons.org/licenses/odbl/summary/"
      }
   ],
   "rate":{
      "limit":2500,
      "remaining":2381,
      "reset":1463184000
   },
   "results":[

   ],
   "status":{
      "code":200,
      "message":"OK"
   },
   "stay_informed":{
      "blog":"http://blog.opencagedata.com",
      "twitter":"https://twitter.com/opencagedata"
   },
   "thanks":"For using an OpenCage Data API",
   "timestamp":{
      "created_http":"Fri, 13 May 2016 10:34:05 GMT",
      "created_unix":1463135645
   },
   "total_results":0
}`
	responseMissingLocality = `{
    "documentation": "https://opencagedata.com/api",
    "licenses": [
        {
            "name": "CC-BY-SA",
            "url": "https://creativecommons.org/licenses/by-sa/3.0/"
        },
        {
            "name": "ODbL",
            "url": "https://opendatacommons.org/licenses/odbl/summary/"
        }
    ],
    "results": [
        {
            "annotations": {
                "DMS": {
                    "lat": "54° 8' 30.29676'' N",
                    "lng": "12° 3' 0.72252'' E"
                },
                "MGRS": "33UUA0732603314",
                "Maidenhead": "JO64ad64aa",
                "Mercator": {
                    "x": 1341422.206,
                    "y": 7162391.731
                },
                "OSM": {
                    "edit_url": "https://www.openstreetmap.org/edit?way=89311416#map=17/54.14175/12.05020",
                    "url": "https://www.openstreetmap.org/?mlat=54.14175&mlon=12.05020#map=17/54.14175/12.05020"
                },
                "callingcode": 49,
                "currency": {
                    "alternate_symbols": [],
                    "decimal_mark": ",",
                    "html_entity": "&#x20AC;",
                    "iso_code": "EUR",
                    "iso_numeric": 978,
                    "name": "Euro",
                    "smallest_denomination": 1,
                    "subunit": "Cent",
                    "subunit_to_unit": 100,
                    "symbol": "€",
                    "symbol_first": 1,
                    "thousands_separator": "."
                },
                "flag": "\ud83c\udde9\ud83c\uddea🇩🇪",
                "geohash": "u38s40nww18cs2tdhe9q",
                "qibla": 136.28,
                "sun": {
                    "rise": {
                        "apparent": 1542177540,
                        "astronomical": 1542170040,
                        "civil": 1542175140,
                        "nautical": 1542172560
                    },
                    "set": {
                        "apparent": 1542208380,
                        "astronomical": 1542215880,
                        "civil": 1542210780,
                        "nautical": 1542213360
                    }
                },
                "timezone": {
                    "name": "Europe/Berlin",
                    "now_in_dst": 0,
                    "offset_sec": 3600,
                    "offset_string": 100,
                    "short_name": "CET"
                },
                "what3words": {
                    "words": "staked.nimbly.scatter"
                }
            },
            "bounds": {
                "northeast": {
                    "lat": 54.1418491,
                    "lng": 12.0503007
                },
                "southwest": {
                    "lat": 54.1416491,
                    "lng": 12.0501007
                }
            },
            "components": {
                "ISO_3166-1_alpha-2": "DE",
                "_type": "building",
                "city_district": "Ortsbeirat 5 : Lütten Klein",
                "country": "Germany",
                "country_code": "de",
                "county": "Rostock",
                "house_number": "30",
                "political_union": "European Union",
                "postcode": "18107",
                "road": "Turkuer Straße",
                "state": "Mecklenburg-Vorpommern",
                "suburb": "Lütten Klein"
            },
            "confidence": 10,
            "formatted": "Turkuer Straße 30, 18107 Rostock, Germany",
            "geometry": {
                "lat": 54.1417491,
                "lng": 12.0502007
            }
        }
    ],
    "status": {
        "code": 200,
        "message": "OK"
    },
    "stay_informed": {
        "blog": "https://blog.opencagedata.com",
        "twitter": "https://twitter.com/opencagedata"
    },
    "thanks": "For using an OpenCage Data API",
    "timestamp": {
        "created_http": "Wed, 14 Nov 2018 11:55:11 GMT",
        "created_unix": 1542196511
    },
    "total_results": 1
}`
)
