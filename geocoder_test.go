package geo_test

import (
	"fmt"
	"os"

	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/bing"
	"github.com/codingsince1985/geo-golang/chained"
	"github.com/codingsince1985/geo-golang/google"
	"github.com/codingsince1985/geo-golang/here"
	"github.com/codingsince1985/geo-golang/mapbox"
	"github.com/codingsince1985/geo-golang/mapquest/nominatim"
	"github.com/codingsince1985/geo-golang/mapquest/open"
	"github.com/codingsince1985/geo-golang/opencage"
	"github.com/codingsince1985/geo-golang/openstreetmap"
)

const addr = "Melbourne VIC"
const lat, lng = -37.8167416, 144.964463
const RADIUS = 50

func ExampleGeocoder() {
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
	// Output: Google Geocoding API
	// Melbourne VIC location is (-37.816279, 144.964246)
	// Address of (-37.816742,144.964463) is 40 Elizabeth St, Melbourne VIC 3004, Australia
	//
	// Mapquest Nominatim
	// Melbourne VIC location is (-37.814218, 144.963161)
	// Address of (-37.816742,144.964463) is Bankwest, Elizabeth Street, Melbourne, City of Melbourne, Greater Melbourne, Victoria, 3000, Australia
	//
	// Mapquest Open streetmaps
	// Melbourne VIC location is (-37.814218, 144.963161)
	// Address of (-37.816742,144.964463) is Elizabeth Street, Melbourne, Victoria, AU
	//
	// OpenCage Data
	// Melbourne VIC location is (-37.814217, 144.963161)
	// Address of (-37.816742,144.964463) is Bankwest, Elizabeth Street, Melbourne VIC 3000, Australia
	//
	// HERE API
	// Melbourne VIC location is (-37.817530, 144.967150)
	// Address of (-37.816742,144.964463) is 40 Elizabeth St, Melbourne VIC 3000, Australia
	//
	// Bing Geocoding API
	// Melbourne VIC location is (-37.824299, 144.977997)
	// Address of (-37.816742,144.964463) is 46 Elizabeth St, Melbourne, VIC 3000
	//
	// Mapbox API
	// Melbourne VIC location is (-37.814200, 144.963200)
	// Address of (-37.816742,144.964463) is Clegs, 60 Elizabeth St, 3000 Melbourne, Australia
	//
	// OpenStreetMap
	// Melbourne VIC location is (-37.814217, 144.963161)
	// Address of (-37.816742,144.964463) is Bankwest, Elizabeth Street, Melbourne, City of Melbourne, Greater Melbourne, Victoria, 3000, Australia
	//
	// ChainedAPI[OpenStreetmap -> Google]
	// Melbourne VIC location is (-37.814217, 144.963161)
	// Address of (-37.816742,144.964463) is Bankwest, Elizabeth Street, Melbourne, City of Melbourne, Greater Melbourne, Victoria, 3000, Australia
}

func try(geocoder geo.Geocoder) {
	location, _ := geocoder.Geocode(addr)
	fmt.Printf("%s location is (%.6f, %.6f)\n", addr, location.Lat, location.Lng)
	address, _ := geocoder.ReverseGeocode(lat, lng)
	fmt.Printf("Address of (%.6f,%.6f) is %s\n\n", lat, lng, address)
}
