GeoService in Go
==
[![GoDoc](https://godoc.org/github.com/codingsince1985/geo-golang?status.svg)](https://godoc.org/github.com/codingsince1985/geo-golang)

A geocoding service developed in Go's way, idiomatic and elegant, not just in golang.

This product is designed to open to any Geocoding service. Based on it,
+ [Google Maps](https://developers.google.com/maps/documentation/geocoding/)
+ MapQuest
 - [Nominatim Search](http://open.mapquestapi.com/nominatim/)
 - [Open Geocoding](http://open.mapquestapi.com/geocoding/)
+ [OpenCage](http://geocoder.opencagedata.com/api.html)
+ [HERE](https://developer.here.com/rest-apis/documentation/geocoder)
+ [Bing](https://msdn.microsoft.com/en-us/library/ff701715.aspx)

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
)

const addr = "Melbourne VIC"
const lat, lng = -37.8167416, 144.964463

func main() {
	// Goole Maps
	try(google.Geocoder())

	// MapQuest Nominatim
	try(nominatim.Geocoder())

	// MapQuest Open
	try(open.Geocoder("MAPQUEST_KEY"))

	// OpenCage Data
	try(opencage.Geocoder("OPENCAGE_KEY"))

	// HERE
	try(here.Geocoder("HERE_APP_ID", "HERE_APP_CODE", RADIUS))

	// Bing
	try(bing.Geocoder("BING_KEY"))
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
// Goole Maps
Melbourne VIC location is {-37.814107 144.96328}
Address of (-37.816742,144.964463) is 66 Elizabeth Street, Melbourne VIC 3000, Australia

// MapQuest Nominatim
Melbourne VIC location is {-37.8142176 144.9631608}
Address of (-37.816742,144.964463) is Bankwest, Elizabeth Street, Melbourne, City of Melbourne, Greater Melbourne, Victoria, 3000, Australia

// MapQuest Open
Melbourne VIC location is {-37.814218 144.963161}
Address of (-37.816742,144.964463) is Elizabeth Street, Melbourne, Victoria, AU

// OpenCage Data
Melbourne VIC location is {-37.8142176 144.9631608}
Address of (-37.816742,144.964463) is Bankwest, Elizabeth Street, Melbourne VIC 3000, Australia

// HERE
Melbourne VIC location is {-37.81753 144.96715}
Address of (-37.816742,144.964463) is 40-44 Elizabeth St, Melbourne VIC 3000, Australia

// Bing
Melbourne VIC location is {-37.82429885864258 144.97799682617188}
Address of (-37.816742,144.964463) is 46 Elizabeth St, Melbourne, VIC 3000
```
License
==
geo-goang is distributed under the terms of the MIT license. See LICENSE for details.