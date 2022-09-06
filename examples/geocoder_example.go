package main

import (
	"fmt"
	"os"

	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/arcgis"
	"github.com/codingsince1985/geo-golang/baidu"
	"github.com/codingsince1985/geo-golang/bing"
	"github.com/codingsince1985/geo-golang/chained"
	"github.com/codingsince1985/geo-golang/frenchapigouv"
	"github.com/codingsince1985/geo-golang/geocod"
	"github.com/codingsince1985/geo-golang/google"
	"github.com/codingsince1985/geo-golang/here"
	"github.com/codingsince1985/geo-golang/locationiq"
	"github.com/codingsince1985/geo-golang/mapbox"
	"github.com/codingsince1985/geo-golang/mapquest/nominatim"
	"github.com/codingsince1985/geo-golang/mapquest/open"
	"github.com/codingsince1985/geo-golang/mapzen"
	"github.com/codingsince1985/geo-golang/opencage"
	"github.com/codingsince1985/geo-golang/openstreetmap"
	"github.com/codingsince1985/geo-golang/pickpoint"
	"github.com/codingsince1985/geo-golang/tomtom"
	"github.com/codingsince1985/geo-golang/yandex"
)

const (
	addr         = "Melbourne VIC"
	lat, lng     = -37.813611, 144.963056
	radius       = 50
	zoom         = 18
	addrFR       = "Champs de Mars Paris"
	latFR, lngFR = 48.854395, 2.304770
)

func main() {
	ExampleGeocoder()
}

// ExampleGeocoder demonstrates the different geocoding services
func ExampleGeocoder() {
	fmt.Println("Google Geocoding API")
	try(google.Geocoder(os.Getenv("GOOGLE_API_KEY")))

	fmt.Println("Mapquest Nominatim")
	try(nominatim.Geocoder(os.Getenv("MAPQUEST_NOMINATIM_KEY")))

	fmt.Println("Mapquest Open streetmaps")
	try(open.Geocoder(os.Getenv("MAPQUEST_OPEN_KEY")))

	fmt.Println("OpenCage Data")
	try(opencage.Geocoder(os.Getenv("OPENCAGE_API_KEY")))

	fmt.Println("HERE Geocoder API")
	try(here.Geocoder(os.Getenv("HERE_APP_ID"), os.Getenv("HERE_APP_CODE"), radius))

	fmt.Println("Bing Geocoding API")
	try(bing.Geocoder(os.Getenv("BING_API_KEY")))

	fmt.Println("Baidu Geocoding API")
	try(baidu.Geocoder(os.Getenv("BAIDU_API_KEY"), "en", "bd09ll"))

	fmt.Println("Mapbox API")
	try(mapbox.Geocoder(os.Getenv("MAPBOX_API_KEY")))

	fmt.Println("OpenStreetMap")
	try(openstreetmap.Geocoder())

	fmt.Println("PickPoint")
	try(pickpoint.Geocoder(os.Getenv("PICKPOINT_API_KEY")))

	fmt.Println("LocationIQ")
	try(locationiq.Geocoder(os.Getenv("LOCATIONIQ_API_KEY"), zoom))

	fmt.Println("ArcGIS")
	try(arcgis.Geocoder(os.Getenv("ARCGIS_TOKEN")))

	fmt.Println("geocod.io")
	try(geocod.Geocoder(os.Getenv("GEOCOD_API_KEY")))

	fmt.Println("Mapzen")
	try(mapzen.Geocoder(os.Getenv("MAPZEN_API_KEY")))

	fmt.Println("TomTom")
	try(tomtom.Geocoder(os.Getenv("TOMTOM_API_KEY")))

	fmt.Println("Yandex")
	try(yandex.Geocoder(os.Getenv("YANDEX_API_KEY")))

	// To use only with french locations or addresses,
	// or values ​​could be estimated and will be false
	fmt.Println("FrenchAPIGouv")
	tryOnlyFRData(frenchapigouv.Geocoder())

	// Chained geocoder will fallback to subsequent geocoders
	fmt.Println("ChainedAPI[OpenStreetmap -> Google]")
	try(chained.Geocoder(
		openstreetmap.Geocoder(),
		google.Geocoder(os.Getenv("GOOGLE_API_KEY")),
	))
	// Output: Google Geocoding API
	// Melbourne VIC location is (-37.813611, 144.963056)
	// Address of (-37.813611,144.963056) is 197 Elizabeth St, Melbourne VIC 3000, Australia
	// Detailed address: &geo.Address{FormattedAddress:"197 Elizabeth St, Melbourne VIC 3000, Australia",
	// 	Street:"Elizabeth Street", HouseNumber:"197", Suburb:"", Postcode:"3000", State:"Victoria",
	// 	StateDistrict:"Melbourne City", County:"", Country:"Australia", CountryCode:"AU", City:"Melbourne"}
	//
	// Mapquest Nominatim
	// Melbourne VIC location is (-37.814218, 144.963161)
	// Address of (-37.813611,144.963056) is Melbourne's GPO, Postal Lane, Melbourne, City of Melbourne, Greater Melbourne, Victoria, 3000, Australia
	// Detailed address: &geo.Address{FormattedAddress:"Melbourne's GPO, Postal Lane, Melbourne, City of Melbourne,
	// 	Greater Melbourne, Victoria, 3000, Australia", Street:"Postal Lane", HouseNumber:"", Suburb:"Melbourne",
	// 	Postcode:"3000", State:"Victoria", StateDistrict:"", County:"City of Melbourne", Country:"Australia", CountryCode:"AU", City:"Melbourne"}
	//
	// Mapquest Open streetmaps
	// Melbourne VIC location is (-37.814218, 144.963161)
	// Address of (-37.813611,144.963056) is Elizabeth Street, Melbourne, Victoria, AU
	// Detailed address: &geo.Address{FormattedAddress:"Elizabeth Street, 3000, Melbourne, Victoria, AU",
	// 	Street:"Elizabeth Street", HouseNumber:"", Suburb:"", Postcode:"3000", State:"Victoria", StateDistrict:"",
	// 	County:"", Country:"", CountryCode:"AU", City:"Melbourne"}
	//
	// OpenCage Data
	// Melbourne VIC location is (-37.814217, 144.963161)
	// Address of (-37.813611,144.963056) is Melbourne's GPO, Postal Lane, Melbourne VIC 3000, Australia
	// Detailed address: &geo.Address{FormattedAddress:"Melbourne's GPO, Postal Lane, Melbourne VIC 3000, Australia",
	// 	Street:"Postal Lane", HouseNumber:"", Suburb:"Melbourne (3000)", Postcode:"3000", State:"Victoria",
	// 	StateDistrict:"", County:"City of Melbourne", Country:"Australia", CountryCode:"AU", City:"Melbourne"}
	//
	// HERE Geocoder API
	// Melbourne VIC location is (-37.817530, 144.967150)
	// Address of (-37.813611,144.963056) is 197 Elizabeth St, Melbourne VIC 3000, Australia
	// Detailed address: &geo.Address{FormattedAddress:"197 Elizabeth St, Melbourne VIC 3000, Australia", Street:"Elizabeth St",
	// 	HouseNumber:"197", Suburb:"", Postcode:"3000", State:"Victoria", StateDistrict:"", County:"", Country:"Australia",
	// 	CountryCode:"AUS", City:"Melbourne"}
	//
	// Bing Geocoding API
	// Melbourne VIC location is (-37.824299, 144.977997)
	// Address of (-37.813611,144.963056) is Elizabeth St, Melbourne, VIC 3000
	// Detailed address: &geo.Address{FormattedAddress:"Elizabeth St, Melbourne, VIC 3000", Street:"Elizabeth St",
	// 	HouseNumber:"", Suburb:"", Postcode:"3000", State:"", StateDistrict:"", County:"", Country:"Australia", CountryCode:"", City:"Melbourne"}
	//
	// Baidu Geocoding API
	// Melbourne VIC location is (31.227015, 121.456967)
	// Address of (-37.813611,144.963056) is 341 Little Bourke Street, Melbourne, Victoria, Australia
	// Detailed address: &geo.Address{FormattedAddress:"341 Little Bourke Street, Melbourne, Victoria, Australia", Street:"Little Bourke Street", HouseNumber:"341", Suburb:"", Postcode:"", State:"Victoria", StateCode:"", StateDistrict:"", County:"", Country:"Australia", CountryCode:"AUS", City:"Melbourne"}
	//
	// Mapbox API
	// Melbourne VIC location is (-37.814200, 144.963200)
	// Address of (-37.813611,144.963056) is Elwood Park Playground, Melbourne, Victoria 3000, Australia
	// Detailed address: &geo.Address{FormattedAddress:"Elwood Park Playground, Melbourne, Victoria 3000, Australia",
	// 	Street:"Elwood Park Playground", HouseNumber:"", Suburb:"", Postcode:"3000", State:"Victoria", StateDistrict:"",
	// 	County:"", Country:"Australia", CountryCode:"AU", City:"Melbourne"}
	//
	// OpenStreetMap
	// Melbourne VIC location is (-37.814217, 144.963161)
	// Address of (-37.813611,144.963056) is Melbourne's GPO, Postal Lane, Chinatown, Melbourne, City of Melbourne, Greater Melbourne, Victoria, 3000, Australia
	// Detailed address: &geo.Address{FormattedAddress:"Melbourne's GPO, Postal Lane, Chinatown, Melbourne, City of Melbourne, Greater Melbourne,
	// 	Victoria, 3000, Australia", Street:"Postal Lane", HouseNumber:"", Suburb:"Melbourne", Postcode:"3000", State:"Victoria",
	// 	StateDistrict:"", County:"", Country:"Australia", CountryCode:"AU", City:"Melbourne"}
	//
	// PickPoint
	// Melbourne VIC location is (-37.814217, 144.963161)
	// Address of (-37.813611,144.963056) is Melbourne's GPO, Postal Lane, Chinatown, Melbourne, City of Melbourne, Greater Melbourne, Victoria, 3000, Australia
	// Detailed address: &geo.Address{FormattedAddress:"Melbourne's GPO, Postal Lane, Chinatown, Melbourne, City of Melbourne, Greater Melbourne,
	// 	Victoria, 3000, Australia", Street:"Postal Lane", HouseNumber:"", Suburb:"Melbourne", Postcode:"3000", State:"Victoria",
	// 	StateDistrict:"", County:"", Country:"Australia", CountryCode:"AU", City:"Melbourne"}
	//
	// LocationIQ
	// Melbourne VIC location is (-37.814217, 144.963161)
	// Address of (-37.813611,144.963056) is Melbourne's GPO, Postal Lane, Chinatown, Melbourne, City of Melbourne, Greater Melbourne, Victoria, 3000, Australia
	// Detailed address: &geo.Address{FormattedAddress:"Melbourne's GPO, Postal Lane, Chinatown, Melbourne, City of Melbourne, Greater Melbourne,
	// 	Victoria, 3000, Australia", Street:"Postal Lane", HouseNumber:"", Suburb:"Melbourne", Postcode:"3000", State:"Victoria",
	// 	StateDistrict:"", County:"", Country:"Australia", CountryCode:"AU", City:"Melbourne"}
	//
	// ArcGIS
	// Melbourne VIC location is (-37.817530, 144.967150)
	// Address of (-37.813611,144.963056) is Melbourne's Gpo
	// Detailed address: &geo.Address{FormattedAddress:"Melbourne's Gpo", Street:"350 Bourke Street Mall", HouseNumber:"350", Suburb:"", Postcode:"3000", State:"Victoria", StateDistrict:"", County:"", Country:"", CountryCode:"AUS", City:""}
	//
	// geocod.io
	// Melbourne VIC location is (28.079357, -80.623618)
	// got <nil> address
	//
	// Mapzen
	// Melbourne VIC location is (45.551136, 11.533929)
	// Address of (-37.813611,144.963056) is Stop 3: Bourke Street Mall, Bourke Street, Melbourne, Australia
	// Detailed address: &geo.Address{FormattedAddress:"Stop 3: Bourke Street Mall, Bourke Street, Melbourne, Australia", Street:"", HouseNumber:"", Suburb:"", Postcode:"", State:"Victoria", StateDistrict:"", County:"", Country:"Australia", CountryCode:"AUS", City:""}
	//
	// TomTom
	// Melbourne VIC location is (-37.815340, 144.963230)
	// Address of (-37.813611,144.963056) is Doyles Road, Elaine, West Central Victoria, Victoria, 3334
	// Detailed address: &geo.Address{FormattedAddress:"Doyles Road, Elaine, West Central Victoria, Victoria, 3334", Street:"Doyles Road", HouseNumber:"", Suburb:"", Postcode:"3334", State:"Victoria", StateDistrict:"", County:"", Country:"Australia", CountryCode:"AU", City:"Elaine"}
	//
	// Yandex
	// Melbourne VIC location is (41.926823, 2.254232)
	// Address of (-37.813611,144.963056) is Victoria, City of Melbourne, Elizabeth Street
	// Detailed address: &geo.Address{FormattedAddress:"Victoria, City of Melbourne, Elizabeth Street", Street:"Elizabeth Street",
	//  HouseNumber:"", Suburb:"", Postcode:"", State:"Victoria", StateDistrict:"", County:"", Country:"Australia", CountryCode:"AU",
	//  City:"City of Melbourne"}
	//
	// FrenchAPIGouv
	//  Champs de Mars Paris location is (2.304770, 48.854395)
	//  Address of (48.854395,2.304770) is 9001, Parc du Champs de Mars, 75007, Paris, Paris, Île-de-France, France
	//  Detailed address: &geo.Address{FormattedAddress:"9001, Parc du Champs de Mars, 75007, Paris, Paris, Île-de-France, France",
	//   Street:"Parc du Champs de Mars", HouseNumber:"9001", Suburb:"", Postcode:"75007", State:" Île-de-France", StateDistrict:"",
	//   County:" Paris", Country:"France", CountryCode:"FRA", City:"Paris"}
	//
	// ChainedAPI[OpenStreetmap -> Google]
	// Melbourne VIC location is (-37.814217, 144.963161)
	// Address of (-37.813611,144.963056) is Melbourne's GPO, Postal Lane, Chinatown, Melbourne, City of Melbourne, Greater Melbourne, Victoria, 3000, Australia
	// Detailed address: &geo.Address{FormattedAddress:"Melbourne's GPO, Postal Lane, Chinatown, Melbourne, City of Melbourne, Greater Melbourne,
	// 	Victoria, 3000, Australia", Street:"Postal Lane", HouseNumber:"", Suburb:"Melbourne", Postcode:"3000", State:"Victoria",
	// 	StateDistrict:"", County:"", Country:"Australia", CountryCode:"AU", City:"Melbourne"}
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
	fmt.Print("\n")
}

func tryOnlyFRData(geocoder geo.Geocoder) {
	location, _ := geocoder.Geocode(addrFR)
	if location != nil {
		fmt.Printf("%s location is (%.6f, %.6f)\n", addrFR, location.Lat, location.Lng)
	} else {
		fmt.Println("got <nil> location")
	}
	address, _ := geocoder.ReverseGeocode(latFR, lngFR)
	if address != nil {
		fmt.Printf("Address of (%.6f,%.6f) is %s\n", latFR, lngFR, address.FormattedAddress)
		fmt.Printf("Detailed address: %#v\n", address)
	} else {
		fmt.Println("got <nil> address")
	}
	fmt.Print("\n")
}
