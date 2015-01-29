// Package geo has all the types and interfaces needed to wrap a geocode/reverse geocode service.
// Google and MapRequest implementations are based on it in ~50 lines of code each.
package geo

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

var TimeoutError = errors.New("TIMEOUT")
var NoResultError = errors.New("NO_RESULT")

var timeoutInMillisecond = time.Millisecond * 1000

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
func (g Geocoder) Geocode(address string) (Location, error) {
	ch := make(chan Location)
	go func() {
		ch <- g.Location(responseData(g.GeocodeUrl(address)))
	}()

	select {
	case location := <-ch:
		return location, nil
	case <-time.After(timeoutInMillisecond):
		return Location{}, TimeoutError
	}
}

// ReverseGeocode returns address for location
func (g Geocoder) ReverseGeocode(l Location) (string, error) {
	ch := make(chan string)
	go func() {
		ch <- g.Address(responseData(g.ReverseGeocodeUrl(l)))
	}()

	select {
	case address := <-ch:
		if address == "" {
			return "", NoResultError
		}
		return address, nil
	case <-time.After(timeoutInMillisecond):
		return "", TimeoutError
	}
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
