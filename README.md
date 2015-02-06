GeoService in Go
=

Why another one? Short answer, I want one developed in Go's way, not just in golang.

This project is designed to open to any Geocoding service. I've implemented Google, MapQuest and OpenCage clients in ~50 LoC each.

Here is how to use it.

```go
package main

import (
	"fmt"
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/google"
	"github.com/codingsince1985/geo-golang/mapquest"
	"github.com/codingsince1985/geo-golang/opencage"
)

func main() {
	var geocoder geo.Geocoder

	geocoder = google.NewGeocoder()
	location, _ := geocoder.Geocode("Melbourne VIC")
	fmt.Println("Google's Melbourne location is", location)

	geocoder = mapquest.NewGeocoder()
	location, _ = geocoder.Geocode("Melbourne VIC")
	fmt.Println("MapQuest's Melbourne location is", location)

	geocoder = opencage.NewGeocoder("YOUR_KEY")
	location, _ = geocoder.Geocode("Melbourne VIC")
	fmt.Println("OpenCage's Melbourne location is", location)
}
```
