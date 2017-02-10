GeoService in Go
==
[![GoDoc](https://godoc.org/github.com/codingsince1985/geo-golang?status.svg)](https://godoc.org/github.com/codingsince1985/geo-golang) [![Build Status](https://travis-ci.org/codingsince1985/geo-golang.svg?branch=master)](https://travis-ci.org/codingsince1985/geo-golang)
[![codecov](https://codecov.io/gh/codingsince1985/geo-golang/branch/master/graph/badge.svg)](https://codecov.io/gh/codingsince1985/geo-golang)
[![Go Report Card](https://goreportcard.com/badge/codingsince1985/geo-golang)](https://goreportcard.com/report/codingsince1985/geo-golang)

## Code Coverage
[![codecov.io](https://codecov.io/gh/codingsince1985/geo-golang/branch/master/graphs/icicle.svg)](https://codecov.io/gh/codingsince1985/geo-golang)

A geocoding service developed in Go's way, idiomatic and elegant, not just in golang.

This product is designed to open to any Geocoding service. Based on it,
+ [Google Maps](https://developers.google.com/maps/documentation/geocoding/)
+ MapQuest
 - [Nominatim Search](http://open.mapquestapi.com/nominatim/)
 - [Open Geocoding](http://open.mapquestapi.com/geocoding/)
+ [OpenCage](http://geocoder.opencagedata.com/api.html)
+ [HERE](https://developer.here.com/rest-apis/documentation/geocoder)
+ [Bing](https://msdn.microsoft.com/en-us/library/ff701715.aspx)
+ [Mapbox](https://www.mapbox.com/developers/api/geocoding/)
+ [OpenStreetMap](https://wiki.openstreetmap.org/wiki/Nominatim)
+ [LocationIQ](http://locationiq.org/)

clients are implemented in ~50 LoC each.

It allows you to switch from one service to another by changing only 1 line, or enjoy all the free quota (requests/sec, day, month...) from them at the same time. Just like this.

```go
package main

import (
	"fmt"
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/bing"
	"github.com/codingsince1985/geo-golang/google"
	"github.com/codingsince1985/geo-golang/here"
	"github.com/codingsince1985/geo-golang/mapquest/nominatim"
	"github.com/codingsince1985/geo-golang/mapquest/open"
	"github.com/codingsince1985/geo-golang/opencage"
	"github.com/codingsince1985/geo-golang/mapbox"
	"github.com/codingsince1985/geo-golang/openstreetmap"
	"github.com/codingsince1985/geo-golang/chained"
	"github.com/codingsince1985/geo-golang/locationiq"
)

const (
	addr     = "Melbourne VIC"
	lat, lng = -37.813611, 144.963056
	RADIUS   = 50
	ZOOM     = 18
)

func main() {
	ExampleGeocoder()
}

func ExampleGeocoder() {
	fmt.Println("Google Geocoding API")
	try(google.Geocoder(os.Getenv("GOOGLE_API_KEY")))

	fmt.Println("Mapquest Nominatim")
	try(nominatim.Geocoder(os.Getenv("MAPQUEST_NOMINATIM_KEY")))

	fmt.Println("Mapquest Open streetmaps")
	try(open.Geocoder(os.Getenv("MAPQUEST_OPEN_KEY")))

	fmt.Println("OpenCage Data")
	try(opencage.Geocoder(os.Getenv("OPENCAGE_API_KEY")))

	fmt.Println("HERE API")
	try(here.Geocoder(os.Getenv("HERE_APP_ID"), os.Getenv("HERE_APP_CODE"), RADIUS))

	fmt.Println("Bing Geocoding API")
	try(bing.Geocoder(os.Getenv("BING_API_KEY")))

	fmt.Println("Mapbox API")
	try(mapbox.Geocoder(os.Getenv("MAPBOX_API_KEY")))

	fmt.Println("OpenStreetMap")
	try(openstreetmap.Geocoder())

	fmt.Println("LocationIQ")
	try(locationiq.Geocoder(os.Getenv("LOCATIONIQ_API_KEY"), ZOOM))

	// Chained geocoder will fallback to subsequent geocoders
	fmt.Println("ChainedAPI[OpenStreetmap -> Google]")
	try(chained.Geocoder(
		openstreetmap.Geocoder(),
		google.Geocoder(os.Getenv("GOOGLE_API_KEY")),
	))
}

func try(geocoder geo.Geocoder) {
	location, _ := geocoder.Geocode(addr)
	if location != nil {
		fmt.Printf("%s location is (%.6f, %.6f)\n", addr, location.Lat, location.Lng)
	} else {
		fmt.Println("got <nil> location")
	}
	address, _ := geocoder.ReverseGeocode(lat, lng)
	if address != nil {
		fmt.Printf("Address of (%.6f,%.6f) is %s\n", lat, lng, address.FormattedAddress)
		fmt.Printf("Detailed address: %#v\n", address)
	} else {
		fmt.Println("got <nil> address")
	}
	fmt.Println("\n")
}
```
###Result
```
Google Geocoding API
Melbourne VIC location is (-37.813611, 144.963056)
Address of (-37.813611,144.963056) is 197 Elizabeth St, Melbourne VIC 3000, Australia
Detailed address: &geo.Address{FormattedAddress:"197 Elizabeth St, Melbourne VIC 3000, Australia", Street:"Elizabeth Street", HouseNumber:"197", Suburb:"", Postcode:"3000", State:"Victoria", StateDistrict:"Melbourne City", County:"", Country:"Australia", CountryCode:"AU", City:"Melbourne"}

Mapquest Nominatim
Melbourne VIC location is (-37.814218, 144.963161)
Address of (-37.813611,144.963056) is Melbourne's GPO, Postal Lane, Melbourne, City of Melbourne, Greater Melbourne, Victoria, 3000, Australia
Detailed address: &geo.Address{FormattedAddress:"Melbourne's GPO, Postal Lane, Melbourne, City of Melbourne, Greater Melbourne, Victoria, 3000, Australia", Street:"Postal Lane", HouseNumber:"", Suburb:"Melbourne", Postcode:"3000", State:"Victoria", StateDistrict:"", County:"City of Melbourne", Country:"Australia", CountryCode:"AU", City:"Melbourne"}

Mapquest Open streetmaps
Melbourne VIC location is (-37.814218, 144.963161)
Address of (-37.813611,144.963056) is Elizabeth Street, 3000, Melbourne, Victoria, AU
Detailed address: &geo.Address{FormattedAddress:"Elizabeth Street, 3000, Melbourne, Victoria, AU", Street:"Elizabeth Street", HouseNumber:"", Suburb:"", Postcode:"3000", State:"Victoria", StateDistrict:"", County:"", Country:"", CountryCode:"AU", City:"Melbourne"}

OpenCage Data
Melbourne VIC location is (-37.814217, 144.963161)
Address of (-37.813611,144.963056) is Melbourne's GPO, Postal Lane, Melbourne VIC 3000, Australia
Detailed address: &geo.Address{FormattedAddress:"Melbourne's GPO, Postal Lane, Melbourne VIC 3000, Australia", Street:"Postal Lane", HouseNumber:"", Suburb:"Melbourne (3000)", Postcode:"3000", State:"Victoria", StateDistrict:"", County:"City of Melbourne", Country:"Australia", CountryCode:"AU", City:"Melbourne"}

HERE API
Melbourne VIC location is (-37.817530, 144.967150)
Address of (-37.813611,144.963056) is 197 Elizabeth St, Melbourne VIC 3000, Australia
Detailed address: &geo.Address{FormattedAddress:"197 Elizabeth St, Melbourne VIC 3000, Australia", Street:"Elizabeth St", HouseNumber:"197", Suburb:"", Postcode:"3000", State:"Victoria", StateDistrict:"", County:"", Country:"Australia", CountryCode:"AUS", City:"Melbourne"}

Bing Geocoding API
Melbourne VIC location is (-37.824299, 144.977997)
Address of (-37.813611,144.963056) is Elizabeth St, Melbourne, VIC 3000
Detailed address: &geo.Address{FormattedAddress:"Elizabeth St, Melbourne, VIC 3000", Street:"Elizabeth St", HouseNumber:"", Suburb:"", Postcode:"3000", State:"", StateDistrict:"", County:"", Country:"Australia", CountryCode:"", City:"Melbourne"}

Mapbox API
Melbourne VIC location is (-37.814200, 144.963200)
Address of (-37.813611,144.963056) is Elwood Park Playground, Melbourne, Victoria 3000, Australia
Detailed address: &geo.Address{FormattedAddress:"Elwood Park Playground, Melbourne, Victoria 3000, Australia", Street:"Elwood Park Playground", HouseNumber:"", Suburb:"", Postcode:"3000", State:"Victoria", StateDistrict:"", County:"", Country:"Australia", CountryCode:"AU", City:"Melbourne"}

OpenStreetMap
Melbourne VIC location is (-37.814217, 144.963161)
Address of (-37.813611,144.963056) is Melbourne's GPO, Postal Lane, Chinatown, Melbourne, City of Melbourne, Greater Melbourne, Victoria, 3000, Australia
Detailed address: &geo.Address{FormattedAddress:"Melbourne's GPO, Postal Lane, Chinatown, Melbourne, City of Melbourne, Greater Melbourne, Victoria, 3000, Australia", Street:"Postal Lane", HouseNumber:"", Suburb:"Melbourne", Postcode:"3000", State:"Victoria", StateDistrict:"", County:"", Country:"Australia", CountryCode:"AU", City:"Melbourne"}

LocationIQ
Melbourne VIC location is (-37.814217, 144.963161)
Address of (-37.813611,144.963056) is Melbourne's GPO, Postal Lane, Chinatown, Melbourne, City of Melbourne, Greater Melbourne, Victoria, 3000, Australia
Detailed address: &geo.Address{FormattedAddress:"Melbourne's GPO, Postal Lane, Chinatown, Melbourne, City of Melbourne, Greater Melbourne, Victoria, 3000, Australia", Street:"Postal Lane", HouseNumber:"", Suburb:"Melbourne", Postcode:"3000", State:"Victoria", StateDistrict:"", County:"", Country:"Australia", CountryCode:"AU", City:"Melbourne"}

ChainedAPI[OpenStreetmap -> Google]
Melbourne VIC location is (-37.814217, 144.963161)
Address of (-37.813611,144.963056) is Melbourne's GPO, Postal Lane, Chinatown, Melbourne, City of Melbourne, Greater Melbourne, Victoria, 3000, Australia
Detailed address: &geo.Address{FormattedAddress:"Melbourne's GPO, Postal Lane, Chinatown, Melbourne, City of Melbourne, Greater Melbourne, Victoria, 3000, Australia", Street:"Postal Lane", HouseNumber:"", Suburb:"Melbourne", Postcode:"3000", State:"Victoria", StateDistrict:"", County:"", Country:"Australia", CountryCode:"AU", City:"Melbourne"}
```
License
==
geo-golang is distributed under the terms of the MIT license. See LICENSE for details.
