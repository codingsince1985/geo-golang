// Package geo has all the types and interfaces needed to implement a geocode service
package geo

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

// GeocodeResponseParser defines functions that parse response of geocode and reverse geocode, and return Location or Address
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
	return g.Location(ResponseData(g.GeocodeUrl(address)))
}

// ReverseGeocode returns address for location
func (g Geocoder) ReverseGeocode(l Location) string {
	return g.Address(ResponseData(g.ReverseGeocodeUrl(l)))
}
