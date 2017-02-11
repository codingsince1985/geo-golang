package open_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/mapquest/open"
	"github.com/stretchr/testify/assert"
)

var key = os.Getenv("MAPQUEST_OPEN_KEY")

func TestGeocode(t *testing.T) {
	ts := testServer(response1)
	defer ts.Close()

	geocoder := open.Geocoder(key, ts.URL+"/")
	location, err := geocoder.Geocode("60 Collins St, Melbourne VIC 3000")
	assert.NoError(t, err)
	assert.Equal(t, geo.Location{Lat: -37.813743, Lng: 144.971745}, *location)
}

func TestReverseGeocode(t *testing.T) {
	ts := testServer(response2)
	defer ts.Close()

	geocoder := open.Geocoder(key, ts.URL+"/")
	address, err := geocoder.ReverseGeocode(-37.813743, 144.971745)
	assert.NoError(t, err)
	assert.True(t, strings.HasPrefix(address.FormattedAddress, "Exhibition Street"))
}

func TestReverseGeocodeWithNoResult(t *testing.T) {
	ts := testServer(response3)
	defer ts.Close()

	geocoder := open.Geocoder(key, ts.URL+"/")
	//geocoder := open.Geocoder(key)
	addr, err := geocoder.ReverseGeocode(-37.813743, 164.971745)
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
   "info":{
      "statuscode":0,
      "copyright":{
         "text":"\u00A9 2016 MapQuest, Inc.",
         "imageUrl":"http://api.mqcdn.com/res/mqlogo.gif",
         "imageAltText":"\u00A9 2016 MapQuest, Inc."
      },
      "messages":[

      ]
   },
   "options":{
      "maxResults":-1,
      "thumbMaps":true,
      "ignoreLatLngInput":false
   },
   "results":[
      {
         "providedLocation":{
            "location":"60 Collins St, Melbourne VIC 3000"
         },
         "locations":[
            {
               "street":"60 Collins Street",
               "adminArea6":"Melbourne",
               "adminArea6Type":"Neighborhood",
               "adminArea5":"Melbourne",
               "adminArea5Type":"City",
               "adminArea4":"City of Melbourne",
               "adminArea4Type":"County",
               "adminArea3":"Victoria",
               "adminArea3Type":"State",
               "adminArea1":"AU",
               "adminArea1Type":"Country",
               "postalCode":"3000",
               "geocodeQualityCode":"P1XAA",
               "geocodeQuality":"POINT",
               "dragPoint":false,
               "sideOfStreet":"N",
               "linkId":"0",
               "unknownInput":"",
               "type":"s",
               "latLng":{
                  "lat":-37.813743,
                  "lng":144.971745
               },
               "displayLatLng":{
                  "lat":-37.813743,
                  "lng":144.971745
               },
               "mapUrl":"http://open.mapquestapi.com/staticmap/v4/getmap?key=Fmjtd|luu8216t21,rn=o5-94255w&type=map&size=225,160&pois=purple-1,-37.8137433689794,144.971745104488,0,0,|&center=-37.8137433689794,144.971745104488&zoom=15&rand=-2113669715"
            }
         ]
      }
   ]
}`
	response2 = `{
   "info":{
      "statuscode":0,
      "copyright":{
         "text":"\u00A9 2016 MapQuest, Inc.",
         "imageUrl":"http://api.mqcdn.com/res/mqlogo.gif",
         "imageAltText":"\u00A9 2016 MapQuest, Inc."
      },
      "messages":[

      ]
   },
   "options":{
      "maxResults":1,
      "thumbMaps":true,
      "ignoreLatLngInput":false
   },
   "results":[
      {
         "providedLocation":{
            "latLng":{
               "lat":-37.813743,
               "lng":144.971745
            }
         },
         "locations":[
            {
               "street":"Exhibition Street",
               "adminArea6":"Melbourne",
               "adminArea6Type":"Neighborhood",
               "adminArea5":"Melbourne",
               "adminArea5Type":"City",
               "adminArea4":"City of Melbourne",
               "adminArea4Type":"County",
               "adminArea3":"Victoria",
               "adminArea3Type":"State",
               "adminArea1":"AU",
               "adminArea1Type":"Country",
               "postalCode":"3000",
               "geocodeQualityCode":"L1AAA",
               "geocodeQuality":"ADDRESS",
               "dragPoint":false,
               "sideOfStreet":"N",
               "linkId":"0",
               "unknownInput":"",
               "type":"s",
               "latLng":{
                  "lat":-37.813621,
                  "lng":144.97161
               },
               "displayLatLng":{
                  "lat":-37.813621,
                  "lng":144.97161
               },
               "mapUrl":"http://open.mapquestapi.com/staticmap/v4/getmap?key=Fmjtd|luu8216t21,rn=o5-94255w&type=map&size=225,160&pois=purple-1,-37.81362105,144.971609569165,0,0,|&center=-37.81362105,144.971609569165&zoom=15&rand=2088657841"
            }
         ]
      }
   ]
}`
	response3 = `{
   "info":{
      "statuscode":0,
      "copyright":{
         "text":"\u00A9 2016 MapQuest, Inc.",
         "imageUrl":"http://api.mqcdn.com/res/mqlogo.gif",
         "imageAltText":"\u00A9 2016 MapQuest, Inc."
      },
      "messages":[

      ]
   },
   "options":{
      "maxResults":1,
      "thumbMaps":true,
      "ignoreLatLngInput":false
   },
   "results":[
      {
         "providedLocation":{
            "latLng":{
               "lat":-37.813743,
               "lng":164.971745
            }
         },
         "locations":[
            {
               "street":"",
               "adminArea6":"",
               "adminArea6Type":"Neighborhood",
               "adminArea5":"",
               "adminArea5Type":"City",
               "adminArea4":"",
               "adminArea4Type":"County",
               "adminArea3":"",
               "adminArea3Type":"State",
               "adminArea1":"",
               "adminArea1Type":"Country",
               "postalCode":"",
               "geocodeQualityCode":"XXXXX",
               "geocodeQuality":"UNKNOWN",
               "dragPoint":false,
               "sideOfStreet":"N",
               "linkId":"0",
               "unknownInput":"",
               "type":"s",
               "latLng":{
                  "lat":-37.813743,
                  "lng":164.971745
               },
               "displayLatLng":{
                  "lat":-37.813743,
                  "lng":164.971745
               },
               "mapUrl":"http://open.mapquestapi.com/staticmap/v4/getmap?key=Fmjtd|luu8216t21,rn=o5-94255w&type=map&size=225,160&pois=purple-1,-37.813743,164.971745,0,0,|&center=-37.813743,164.971745&zoom=15&rand=342866536"
            }
         ]
      }
   ]
}`
)
