package google_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/google"
	"github.com/stretchr/testify/assert"
)

var token = os.Getenv("GOOGLE_API_KEY")

func TestGeocode(t *testing.T) {
	ts := testServer(response1)
	defer ts.Close()

	geocoder := google.Geocoder(token, ts.URL+"/")
	location, err := geocoder.Geocode("60 Collins St, Melbourne VIC 3000")
	assert.NoError(t, err)
	assert.Equal(t, geo.Location{Lat: -37.8137683, Lng: 144.9718448}, *location)
}

func TestReverseGeocode(t *testing.T) {
	ts := testServer(response2)
	defer ts.Close()

	geocoder := google.Geocoder(token, ts.URL+"/")
	address, err := geocoder.ReverseGeocode(-37.8137683, 144.9718448)
	assert.NoError(t, err)
	assert.True(t, strings.HasPrefix(address.FormattedAddress, "60 Collins St"))
	assert.True(t, strings.HasPrefix(address.Street, "Collins St"))
	assert.Equal(t, address.StateCode, "VIC")
}

func TestReverseGeocodeWithNoResult(t *testing.T) {
	ts := testServer(response3)
	defer ts.Close()

	geocoder := google.Geocoder(token, ts.URL+"/")
	addr, err := geocoder.ReverseGeocode(-37.8137683, 164.9718448)
	assert.Nil(t, err)
	assert.Nil(t, addr)
}

func testServer(response string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte(response))
	}))
}

const (
	response1 = `{
   "results" : [
      {
         "address_components" : [
            {
               "long_name" : "60",
               "short_name" : "60",
               "types" : [ "street_number" ]
            },
            {
               "long_name" : "Collins Street",
               "short_name" : "Collins St",
               "types" : [ "route" ]
            },
            {
               "long_name" : "Melbourne",
               "short_name" : "Melbourne",
               "types" : [ "locality", "political" ]
            },
            {
               "long_name" : "Victoria",
               "short_name" : "VIC",
               "types" : [ "administrative_area_level_1", "political" ]
            },
            {
               "long_name" : "Australia",
               "short_name" : "AU",
               "types" : [ "country", "political" ]
            },
            {
               "long_name" : "3000",
               "short_name" : "3000",
               "types" : [ "postal_code" ]
            }
         ],
         "formatted_address" : "60 Collins St, Melbourne VIC 3000, Australia",
         "geometry" : {
            "location" : {
               "lat" : -37.8137683,
               "lng" : 144.9718448
            },
            "location_type" : "ROOFTOP",
            "viewport" : {
               "northeast" : {
                  "lat" : -37.8124193197085,
                  "lng" : 144.9731937802915
               },
               "southwest" : {
                  "lat" : -37.8151172802915,
                  "lng" : 144.9704958197085
               }
            }
         },
         "place_id" : "ChIJq9MCeshC1moRB2Z1U9DOwK0",
         "types" : [ "street_address" ]
      }
   ],
   "status" : "OK"
}`
	response2 = `{
   "results" : [
      {
         "address_components" : [
            {
               "long_name" : "60",
               "short_name" : "60",
               "types" : [ "street_number" ]
            },
            {
               "long_name" : "Collins Street",
               "short_name" : "Collins St",
               "types" : [ "route" ]
            },
            {
               "long_name" : "Melbourne",
               "short_name" : "Melbourne",
               "types" : [ "locality", "political" ]
            },
            {
               "long_name" : "Victoria",
               "short_name" : "VIC",
               "types" : [ "administrative_area_level_1", "political" ]
            },
            {
               "long_name" : "Australia",
               "short_name" : "AU",
               "types" : [ "country", "political" ]
            },
            {
               "long_name" : "3000",
               "short_name" : "3000",
               "types" : [ "postal_code" ]
            }
         ],
         "formatted_address" : "60 Collins St, Melbourne VIC 3000, Australia",
         "geometry" : {
            "location" : {
               "lat" : -37.8137683,
               "lng" : 144.9718448
            },
            "location_type" : "ROOFTOP",
            "viewport" : {
               "northeast" : {
                  "lat" : -37.8124193197085,
                  "lng" : 144.9731937802915
               },
               "southwest" : {
                  "lat" : -37.8151172802915,
                  "lng" : 144.9704958197085
               }
            }
         },
         "place_id" : "ChIJq9MCeshC1moRB2Z1U9DOwK0",
         "types" : [ "street_address" ]
      },
      {
         "address_components" : [
            {
               "long_name" : "Melbourne",
               "short_name" : "Melbourne",
               "types" : [ "locality", "political" ]
            },
            {
               "long_name" : "Victoria",
               "short_name" : "VIC",
               "types" : [ "administrative_area_level_1", "political" ]
            },
            {
               "long_name" : "Australia",
               "short_name" : "AU",
               "types" : [ "country", "political" ]
            }
         ],
         "formatted_address" : "Melbourne VIC, Australia",
         "geometry" : {
            "bounds" : {
               "northeast" : {
                  "lat" : -37.7994893,
                  "lng" : 144.9890618
               },
               "southwest" : {
                  "lat" : -37.8546255,
                  "lng" : 144.9514222
               }
            },
            "location" : {
               "lat" : -37.8142155,
               "lng" : 144.9632307
            },
            "location_type" : "APPROXIMATE",
            "viewport" : {
               "northeast" : {
                  "lat" : -37.7994893,
                  "lng" : 144.9890618
               },
               "southwest" : {
                  "lat" : -37.8546255,
                  "lng" : 144.9514222
               }
            }
         },
         "place_id" : "ChIJgf0RD69C1moR4OeMIXVWBAU",
         "types" : [ "locality", "political" ]
      },
      {
         "address_components" : [
            {
               "long_name" : "Melbourne",
               "short_name" : "Melbourne",
               "types" : [ "colloquial_area", "locality", "political" ]
            },
            {
               "long_name" : "Victoria",
               "short_name" : "VIC",
               "types" : [ "administrative_area_level_1", "political" ]
            },
            {
               "long_name" : "Australia",
               "short_name" : "AU",
               "types" : [ "country", "political" ]
            }
         ],
         "formatted_address" : "Melbourne VIC, Australia",
         "geometry" : {
            "bounds" : {
               "northeast" : {
                  "lat" : -37.4598457,
                  "lng" : 145.76474
               },
               "southwest" : {
                  "lat" : -38.2607199,
                  "lng" : 144.3944921
               }
            },
            "location" : {
               "lat" : -37.814107,
               "lng" : 144.96328
            },
            "location_type" : "APPROXIMATE",
            "viewport" : {
               "northeast" : {
                  "lat" : -37.4598457,
                  "lng" : 145.76474
               },
               "southwest" : {
                  "lat" : -38.2607199,
                  "lng" : 144.3944921
               }
            }
         },
         "place_id" : "ChIJ90260rVG1moRkM2MIXVWBAQ",
         "types" : [ "colloquial_area", "locality", "political" ]
      },
      {
         "address_components" : [
            {
               "long_name" : "3000",
               "short_name" : "3000",
               "types" : [ "postal_code" ]
            },
            {
               "long_name" : "Melbourne",
               "short_name" : "Melbourne",
               "types" : [ "locality", "political" ]
            },
            {
               "long_name" : "Victoria",
               "short_name" : "VIC",
               "types" : [ "administrative_area_level_1", "political" ]
            },
            {
               "long_name" : "Australia",
               "short_name" : "AU",
               "types" : [ "country", "political" ]
            }
         ],
         "formatted_address" : "Melbourne VIC 3000, Australia",
         "geometry" : {
            "bounds" : {
               "northeast" : {
                  "lat" : -37.7994893,
                  "lng" : 144.9890618
               },
               "southwest" : {
                  "lat" : -37.8300559,
                  "lng" : 144.9514222
               }
            },
            "location" : {
               "lat" : -37.8152065,
               "lng" : 144.963937
            },
            "location_type" : "APPROXIMATE",
            "viewport" : {
               "northeast" : {
                  "lat" : -37.7994893,
                  "lng" : 144.9890618
               },
               "southwest" : {
                  "lat" : -37.8300559,
                  "lng" : 144.9514222
               }
            }
         },
         "place_id" : "ChIJm7IcwrhC1moREDUuRnhWBBw",
         "types" : [ "postal_code" ]
      },
      {
         "address_components" : [
            {
               "long_name" : "CBD & South Melbourne",
               "short_name" : "CBD & South Melbourne",
               "types" : [ "political" ]
            },
            {
               "long_name" : "Victoria",
               "short_name" : "VIC",
               "types" : [ "administrative_area_level_1", "political" ]
            },
            {
               "long_name" : "Australia",
               "short_name" : "AU",
               "types" : [ "country", "political" ]
            }
         ],
         "formatted_address" : "CBD & South Melbourne, VIC, Australia",
         "geometry" : {
            "bounds" : {
               "northeast" : {
                  "lat" : -37.7729624,
                  "lng" : 145.0155638
               },
               "southwest" : {
                  "lat" : -37.8574994,
                  "lng" : 144.8984946
               }
            },
            "location" : {
               "lat" : -37.8362164,
               "lng" : 144.9501708
            },
            "location_type" : "APPROXIMATE",
            "viewport" : {
               "northeast" : {
                  "lat" : -37.7729624,
                  "lng" : 145.0155638
               },
               "southwest" : {
                  "lat" : -37.8574994,
                  "lng" : 144.8984946
               }
            }
         },
         "place_id" : "ChIJORuuCkxd1moRNMrml7yk-C8",
         "types" : [ "political" ]
      },
      {
         "address_components" : [
            {
               "long_name" : "Victoria",
               "short_name" : "VIC",
               "types" : [ "administrative_area_level_1", "political" ]
            },
            {
               "long_name" : "Australia",
               "short_name" : "AU",
               "types" : [ "country", "political" ]
            }
         ],
         "formatted_address" : "Victoria, Australia",
         "geometry" : {
            "bounds" : {
               "northeast" : {
                  "lat" : -33.9810507,
                  "lng" : 149.9764882
               },
               "southwest" : {
                  "lat" : -39.2247306,
                  "lng" : 140.9624773
               }
            },
            "location" : {
               "lat" : -37.4713077,
               "lng" : 144.7851531
            },
            "location_type" : "APPROXIMATE",
            "viewport" : {
               "northeast" : {
                  "lat" : -33.9810507,
                  "lng" : 149.9764882
               },
               "southwest" : {
                  "lat" : -39.2247306,
                  "lng" : 140.9624773
               }
            }
         },
         "place_id" : "ChIJT5UYfksx1GoRNJWCvuL8Tlo",
         "types" : [ "administrative_area_level_1", "political" ]
      },
      {
         "address_components" : [
            {
               "long_name" : "Australia",
               "short_name" : "AU",
               "types" : [ "country", "political" ]
            }
         ],
         "formatted_address" : "Australia",
         "geometry" : {
            "bounds" : {
               "northeast" : {
                  "lat" : -9.2198215,
                  "lng" : 159.255497
               },
               "southwest" : {
                  "lat" : -54.7772185,
                  "lng" : 112.9214336
               }
            },
            "location" : {
               "lat" : -25.274398,
               "lng" : 133.775136
            },
            "location_type" : "APPROXIMATE",
            "viewport" : {
               "northeast" : {
                  "lat" : -0.6911343999999999,
                  "lng" : 166.7429167
               },
               "southwest" : {
                  "lat" : -51.66332320000001,
                  "lng" : 100.0911072
               }
            }
         },
         "place_id" : "ChIJ38WHZwf9KysRUhNblaFnglM",
         "types" : [ "country", "political" ]
      }
   ],
   "status" : "OK"
}`
	response3 = `{
   "results" : [],
   "status" : "ZERO_RESULTS"
}`
)
