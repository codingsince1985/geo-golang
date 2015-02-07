GeoService in Go
=

A geocoding service developed in Go's way, idiomatic and elegant, not just in golang.

This product is designed to open to any Geocoding service. Based on it
* [Google Maps](https://developers.google.com/maps/documentation/geocoding/)
* [MapQuest](http://www.mapquestapi.com/geocoding/)
* [OpenCage](http://geocoder.opencagedata.com/api.html)

clients are implemented in ~50 LoC each.

It allows you to switch from one service to another by changing only **one** line. Just like this!

```go
package main

import (
	"fmt"
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/google"
	"github.com/codingsince1985/geo-golang/mapquest"
	"github.com/codingsince1985/geo-golang/opencage"
)

const addr = "Melbourne VIC"
const lat, lng = -37.8167416, 144.964463

func main() {
	// Goole Maps
	try(google.NewGeocoder())

	// MapQuest
	try(mapquest.NewGeocoder())

	// OpenCage Data
	try(opencage.NewGeocoder("OPENCAGE_KEY"))
}

func try(geocoder geo.Geocoder) {
	location, _ := geocoder.Geocode(addr)
	fmt.Printf("%s location is %v\n", addr, location)
	address, _ := geocoder.ReverseGeocode(lat, lng)
	fmt.Printf("Address of (%f,%f) is %s\n\n", lat, lng, address)
}
```
