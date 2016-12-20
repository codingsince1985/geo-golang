package bing_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/bing"
	"github.com/stretchr/testify/assert"
)

var key = os.Getenv("BING_API_KEY")

func TestGeocode(t *testing.T) {
	ts := testServer(response1)
	defer ts.Close()

	geocoder := bing.Geocoder(key, ts.URL+"/")
	location, err := geocoder.Geocode("60 Collins St, Melbourne VIC")

	assert.NoError(t, err)
	assert.Equal(t, geo.Location{Lat: -37.81375, Lng: 144.97176}, *location)
}

func TestReverseGeocode(t *testing.T) {
	ts := testServer(response2)
	defer ts.Close()

	geocoder := bing.Geocoder(key, ts.URL+"/")
	address, err := geocoder.ReverseGeocode(-37.81375, 144.97176)

	assert.NoError(t, err)
	assert.True(t, strings.Index(address.FormattedAddress, "Collins St") > 0)
}

func TestReverseGeocodeWithNoResult(t *testing.T) {
	ts := testServer(response3)
	defer ts.Close()

	geocoder := bing.Geocoder(key, ts.URL+"/")
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
	response1 = `{
   "authenticationResultCode":"ValidCredentials",
   "brandLogoUri":"http:\/\/dev.virtualearth.net\/Branding\/logo_powered_by.png",
   "copyright":"Copyright © 2016 Microsoft and its suppliers. All rights reserved. This API cannot be accessed and the content and any results may not be used, reproduced or transmitted in any manner without express written permission from Microsoft Corporation.",
   "resourceSets":[
      {
         "estimatedTotal":4,
         "resources":[
            {
               "__type":"Location:http:\/\/schemas.microsoft.com\/search\/local\/ws\/rest\/v1",
               "bbox":[
                  -37.8148742,
                  144.970337,
                  -37.8126258,
                  144.973183
               ],
               "name":"60 Collins St, Melbourne VIC 3000, Australia",
               "point":{
                  "type":"Point",
                  "coordinates":[
                     -37.81375,
                     144.97176
                  ]
               },
               "address":{
                  "addressLine":"60 Collins St",
                  "adminDistrict":"VIC",
                  "countryRegion":"Australia",
                  "formattedAddress":"60 Collins St, Melbourne VIC 3000, Australia",
                  "locality":"Melbourne",
                  "postalCode":"3000"
               },
               "confidence":"Medium",
               "entityType":"Address",
               "geocodePoints":[
                  {
                     "type":"Point",
                     "coordinates":[
                        -37.81375,
                        144.97176
                     ],
                     "calculationMethod":"Rooftop",
                     "usageTypes":[
                        "Display"
                     ]
                  },
                  {
                     "type":"Point",
                     "coordinates":[
                        -37.81393,
                        144.97185
                     ],
                     "calculationMethod":"Rooftop",
                     "usageTypes":[
                        "Route"
                     ]
                  }
               ],
               "matchCodes":[
                  "Ambiguous",
                  "Good"
               ]
            },
            {
               "__type":"Location:http:\/\/schemas.microsoft.com\/search\/local\/ws\/rest\/v1",
               "bbox":[
                  -37.7563242,
                  145.0049782,
                  -37.7540758,
                  145.0078218
               ],
               "name":"60 Collins St, Thornbury VIC 3071, Australia",
               "point":{
                  "type":"Point",
                  "coordinates":[
                     -37.7552,
                     145.0064
                  ]
               },
               "address":{
                  "addressLine":"60 Collins St",
                  "adminDistrict":"VIC",
                  "countryRegion":"Australia",
                  "formattedAddress":"60 Collins St, Thornbury VIC 3071, Australia",
                  "locality":"Thornbury",
                  "postalCode":"3071"
               },
               "confidence":"Medium",
               "entityType":"Address",
               "geocodePoints":[
                  {
                     "type":"Point",
                     "coordinates":[
                        -37.7552,
                        145.0064
                     ],
                     "calculationMethod":"Rooftop",
                     "usageTypes":[
                        "Display"
                     ]
                  },
                  {
                     "type":"Point",
                     "coordinates":[
                        -37.75505,
                        145.00642
                     ],
                     "calculationMethod":"Rooftop",
                     "usageTypes":[
                        "Route"
                     ]
                  }
               ],
               "matchCodes":[
                  "Ambiguous",
                  "Good"
               ]
            },
            {
               "__type":"Location:http:\/\/schemas.microsoft.com\/search\/local\/ws\/rest\/v1",
               "bbox":[
                  -37.9809042,
                  145.0587838,
                  -37.9786558,
                  145.0616362
               ],
               "name":"60 Collins St, Mentone VIC 3194, Australia",
               "point":{
                  "type":"Point",
                  "coordinates":[
                     -37.97978,
                     145.06021
                  ]
               },
               "address":{
                  "addressLine":"60 Collins St",
                  "adminDistrict":"VIC",
                  "countryRegion":"Australia",
                  "formattedAddress":"60 Collins St, Mentone VIC 3194, Australia",
                  "locality":"Mentone",
                  "postalCode":"3194"
               },
               "confidence":"Medium",
               "entityType":"Address",
               "geocodePoints":[
                  {
                     "type":"Point",
                     "coordinates":[
                        -37.97978,
                        145.06021
                     ],
                     "calculationMethod":"Rooftop",
                     "usageTypes":[
                        "Display"
                     ]
                  },
                  {
                     "type":"Point",
                     "coordinates":[
                        -37.97961,
                        145.06024
                     ],
                     "calculationMethod":"Rooftop",
                     "usageTypes":[
                        "Route"
                     ]
                  }
               ],
               "matchCodes":[
                  "Ambiguous",
                  "Good"
               ]
            },
            {
               "__type":"Location:http:\/\/schemas.microsoft.com\/search\/local\/ws\/rest\/v1",
               "bbox":[
                  -37.7392142,
                  144.8016785,
                  -37.7369658,
                  144.8045215
               ],
               "name":"60 Collins St, St Albans VIC 3021, Australia",
               "point":{
                  "type":"Point",
                  "coordinates":[
                     -37.73809,
                     144.8031
                  ]
               },
               "address":{
                  "addressLine":"60 Collins St",
                  "adminDistrict":"VIC",
                  "countryRegion":"Australia",
                  "formattedAddress":"60 Collins St, St Albans VIC 3021, Australia",
                  "locality":"St Albans",
                  "postalCode":"3021"
               },
               "confidence":"Medium",
               "entityType":"Address",
               "geocodePoints":[
                  {
                     "type":"Point",
                     "coordinates":[
                        -37.73809,
                        144.8031
                     ],
                     "calculationMethod":"Rooftop",
                     "usageTypes":[
                        "Display"
                     ]
                  },
                  {
                     "type":"Point",
                     "coordinates":[
                        -37.73807,
                        144.80292
                     ],
                     "calculationMethod":"Rooftop",
                     "usageTypes":[
                        "Route"
                     ]
                  }
               ],
               "matchCodes":[
                  "Ambiguous",
                  "Good"
               ]
            }
         ]
      }
   ],
   "statusCode":200,
   "statusDescription":"OK",
   "traceId":"f0ff69bb189c4fdba96a05a0ad86658c|HK20271556|02.00.164.2600|HK2SCH010280621, HK2SCH010281221, HK2SCH010281326, HK2SCH010260633, HK2SCH010310422, i-0186bca5.ap-southeast-1b"
}`
	response2 = `{
   "authenticationResultCode":"ValidCredentials",
   "brandLogoUri":"http:\/\/dev.virtualearth.net\/Branding\/logo_powered_by.png",
   "copyright":"Copyright © 2016 Microsoft and its suppliers. All rights reserved. This API cannot be accessed and the content and any results may not be used, reproduced or transmitted in any manner without express written permission from Microsoft Corporation.",
   "resourceSets":[
      {
         "estimatedTotal":2,
         "resources":[
            {
               "__type":"Location:http:\/\/schemas.microsoft.com\/search\/local\/ws\/rest\/v1",
               "bbox":[
                  -37.817612717570675,
                  144.96524105171127,
                  -37.809887282429322,
                  144.9782789482887
               ],
               "name":"58 Collins St, Melbourne, VIC 3000",
               "point":{
                  "type":"Point",
                  "coordinates":[
                     -37.81375,
                     144.97176
                  ]
               },
               "address":{
                  "addressLine":"58 Collins St",
                  "adminDistrict":"VIC",
                  "countryRegion":"Australia",
                  "formattedAddress":"58 Collins St, Melbourne, VIC 3000",
                  "locality":"Melbourne",
                  "postalCode":"3000"
               },
               "confidence":"Medium",
               "entityType":"Address",
               "geocodePoints":[
                  {
                     "type":"Point",
                     "coordinates":[
                        -37.81375,
                        144.97176
                     ],
                     "calculationMethod":"Interpolation",
                     "usageTypes":[
                        "Display",
                        "Route"
                     ]
                  }
               ],
               "matchCodes":[
                  "Ambiguous",
                  "Good"
               ]
            },
            {
               "__type":"Location:http:\/\/schemas.microsoft.com\/search\/local\/ws\/rest\/v1",
               "bbox":[
                  -37.817612717570675,
                  144.96524105171127,
                  -37.809887282429322,
                  144.9782789482887
               ],
               "name":"Collins St, Melbourne, VIC 3000",
               "point":{
                  "type":"Point",
                  "coordinates":[
                     -37.81375,
                     144.97176
                  ]
               },
               "address":{
                  "addressLine":"Collins St",
                  "adminDistrict":"VIC",
                  "countryRegion":"Australia",
                  "formattedAddress":"Collins St, Melbourne, VIC 3000",
                  "locality":"Melbourne",
                  "postalCode":"3000"
               },
               "confidence":"Medium",
               "entityType":"Address",
               "geocodePoints":[
                  {
                     "type":"Point",
                     "coordinates":[
                        -37.81375,
                        144.97176
                     ],
                     "calculationMethod":"Interpolation",
                     "usageTypes":[
                        "Display",
                        "Route"
                     ]
                  }
               ],
               "matchCodes":[
                  "Ambiguous",
                  "Good"
               ]
            }
         ]
      }
   ],
   "statusCode":200,
   "statusDescription":"OK",
   "traceId":"2c056005fa6643a6ac6bb8a9bb2a3a5c|HK20271556|02.00.164.2600|HK2SCH010280621, HK2SCH010281619"
}`
	response3 = `{
   "authenticationResultCode":"ValidCredentials",
   "brandLogoUri":"http:\/\/dev.virtualearth.net\/Branding\/logo_powered_by.png",
   "copyright":"Copyright © 2016 Microsoft and its suppliers. All rights reserved. This API cannot be accessed and the content and any results may not be used, reproduced or transmitted in any manner without express written permission from Microsoft Corporation.",
   "resourceSets":[
      {
         "estimatedTotal":0,
         "resources":[

         ]
      }
   ],
   "statusCode":200,
   "statusDescription":"OK",
   "traceId":"101c3af9b0984cd0937b9e4dde2910a6|HK20271556|02.00.164.2600|HK2SCH010280621, HK2SCH010281619"
}`
)
