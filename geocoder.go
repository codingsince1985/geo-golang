// Package geo has all the types and interfaces needed to wrap a geocode/reverse geocode service.
// Google and MapRequest implementations are based on it in ~50 lines of code each.
package geo

import (
	"io/ioutil"
	"net/http"
)

// Location is the output of Geocode and also the input of ReverseGeocode
type Location struct {
	Lat, Lng float64
}

// Endpoint contains BaseUrl, on which geocode and reverse geocode urls is built
type Endpoint struct {
	BaseUrl string
}

// GeocodeEndpointBuilder defines functions that build urls for geocode and reverse geocode
type GeocodeEndpointBuilder interface {
	GeocodeUrl(string) string
	ReverseGeocodeUrl(Location) string
}

// GeocodeResponseParser defines functions that parse response of geocode and reverse geocode
type GeocodeResponseParser interface {
	Location([]byte) Location
	Address([]byte) string
}

// Geocoder has GeocodeEndpointBuilder and GeocodeResponseParser
type Geocoder struct {
	GeocodeEndpointBuilder
	GeocodeResponseParser
}

// Geocode returns Location for address
func (g Geocoder) Geocode(address string) Location {
	return g.Location(responseData(g.GeocodeUrl(address)))
}

// ReverseGeocode returns address for location
func (g Geocoder) ReverseGeocode(l Location) string {
	return g.Address(responseData(g.ReverseGeocodeUrl(l)))
}

// ResponseData gets response from url
func responseData(url string) []byte {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil
	}

	return data
}
