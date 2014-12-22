GeoService in golang
=

Why another one? Short answer, I want one developed in Go way, not just in Go.

This project is designed to open to any Geocoding service. All you need to do is following Google or MapQuest implementations provided to write yours in 50 LoC. In fact, copy then change ~20 out of the 50 lines is more accurate.

Here is how to use it.

```go
package main

import (
	"fmt"
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/google"
	"github.com/codingsince1985/geo-golang/mapquest"
)

func main() {
	var geocoder geo.Geocoder

	geocoder = google.Geocoder
	fmt.Println("Google's Melbourne location is", geocoder.Geocode("Melbourne VIC"))

	geocoder = mapquest.Geocoder
	fmt.Println("MapQuest's Melbourne location is", geocoder.Geocode("Melbourne VIC"))
}
```
