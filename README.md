GeoService in Go
=

Why another one? Short answer, I want one developed in Go way, not just in golang.

This project is designed to open to any Geocoding service. All you need to do is following Google, MapQuest, or OpenCage implementation provided to write yours in 50 LoC.

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

	geocoder = google.Geocoder
	location, _ := geocoder.Geocode("Melbourne VIC")
	fmt.Println("Google's Melbourne location is", location)

	geocoder = mapquest.Geocoder
	location, _ = geocoder.Geocode("Melbourne VIC")
	fmt.Println("MapQuest's Melbourne location is", location)

	geocoder = opencage.Geocoder
	location, _ = geocoder.Geocode("Melbourne VIC")
	fmt.Println("OpenCage's Melbourne location is", location)
}
```
