GeoService in Go
=

A geocoding service developed in Go's way, idiomatic and elegant, not just in golang.

This product is designed to open to any Geocoding service. Based on it,
+ [Google Maps](https://developers.google.com/maps/documentation/geocoding/)
+ MapQuest
 - [Nominatim](http://open.mapquestapi.com/nominatim/)
 - [Open](http://open.mapquestapi.com/geocoding/)
+ [OpenCage](http://geocoder.opencagedata.com/api.html)
+ [HERE](https://developer.here.com/rest-apis/documentation/geocoder)

clients are implemented in ~50 LoC each.

It allows you to switch from one service to another by changing only 1 line. Just like this.

```go
package main

import (
	"fmt"
	"github.com/codingsince1985/geo-golang"
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
}

func try(geocoder geo.Geocoder) {
	location, _ := geocoder.Geocode(addr)
	fmt.Printf("%s location is %v\n", addr, location)
	address, _ := geocoder.ReverseGeocode(lat, lng)
	fmt.Printf("Address of (%f,%f) is %s\n\n", lat, lng, address)
}
```
