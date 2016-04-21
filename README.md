GeoService in Go
==
[![GoDoc](https://godoc.org/github.com/codingsince1985/geo-golang?status.svg)](https://godoc.org/github.com/codingsince1985/geo-golang) [![Build Status](https://travis-ci.org/codingsince1985/geo-golang.svg?branch=master)](https://travis-ci.org/codingsince1985/geo-golang) [![codecov.io](https://codecov.io/github/codingsince1985/geo-golang/coverage.svg?branch=master)](https://codecov.io/github/codingsince1985/geo-golang?branch=master) [![Go Report Card](https://goreportcard.com/badge/codingsince1985/geo-golang)](https://goreportcard.com/report/codingsince1985/geo-golang)

![codecov.io](https://codecov.io/github/codingsince1985/geo-golang/branch.svg?branch=master)

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
)

const addr = "Melbourne VIC"
const lat, lng = -37.8167416, 144.964463

func main() {
	fmt.Println("Google Geocoding API")
	try(google.Geocoder(os.Getenv("GOOGLE_API_KEY")))

	fmt.Println("Mapquest Nominatim")
	try(nominatim.Geocoder(os.Getenv("MAPQUEST_NOMINATUM_KEY")))

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

	// Chained geocoder will fallback to subsequent geocoders
	fmt.Println("ChainedAPI[OpenStreetmap -> Google]")
	try(chained.Geocoder(
		openstreetmap.Geocoder(),
		google.Geocoder(os.Getenv("GOOGLE_API_KEY")),
	))
}

func try(geocoder geo.Geocoder) {
	location, _ := geocoder.Geocode(addr)
	fmt.Printf("%s location is %v\n", addr, location)
	address, _ := geocoder.ReverseGeocode(lat, lng)
	fmt.Printf("Address of (%f,%f) is %s\n\n", lat, lng, address)
}
```
###Result
```
Google Geocoding API
Melbourne VIC location is {-37.814107 144.96328}
Address of (-37.816742,144.964463) is 66 Elizabeth St, Melbourne VIC 3000, Australia

Mapquest Nominatim
Melbourne VIC location is {-37.8142176 144.9631608}
Address of (-37.816742,144.964463) is Bankwest, Elizabeth Street, Melbourne, City of Melbourne, Greater Melbourne, Victoria, 3000, Australia

Mapquest Open streetmaps
Melbourne VIC location is {-37.814218 144.963161}
Address of (-37.816742,144.964463) is Elizabeth Street, Melbourne, Victoria, AU

OpenCage Data
Melbourne VIC location is {-37.8142175 144.9631608}
Address of (-37.816742,144.964463) is Bankwest, Elizabeth Street, Melbourne VIC 3000, Australia

HERE API
Melbourne VIC location is {-37.81753 144.96715}
Address of (-37.816742,144.964463) is 40 Elizabeth St, Melbourne VIC 3000, Australia

Bing Geocoding API
Melbourne VIC location is {-37.82429885864258 144.97799682617188}
Address of (-37.816742,144.964463) is 46 Elizabeth St, Melbourne, VIC 3000

Mapbox API
Melbourne VIC location is {-37.8142 144.9632}
Address of (-37.816742,144.964463) is Bankwest ATM, 43-55 Elizabeth St, 3000 Melbourne, Australia

OpenStreetMap
Melbourne VIC location is {-37.8142175 144.9631608}
Address of (-37.816742,144.964463) is Bankwest, Elizabeth Street, Melbourne, City of Melbourne, Greater Melbourne, Victoria, 3000, Australia

ChainedAPI[OpenStreetmap -> Google]
Melbourne VIC location is {-37.8142175 144.9631608}
Address of (-37.816742,144.964463) is Bankwest, Elizabeth Street, Melbourne, City of Melbourne, Greater Melbourne, Victoria, 3000, Australia
```
License
==
geo-golang is distributed under the terms of the MIT license. See LICENSE for details.
