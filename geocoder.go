// Package geo is a generic framework to develop geocode/reverse geocode clients
package geo

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// TimeoutError occurs when no response returned within timeoutInSeconds
var TimeoutError = errors.New("TIMEOUT")

// NoResultError occurs when no result returned
var NoResultError = errors.New("NO_RESULT")

var timeoutInSeconds = time.Second * 4

// Location is the output of Geocode
type Location struct {
	Lat, Lng float64
}

// Endpoint contains base url, on which geocode/reverse geocode urls are built
type Endpoint string

// GeocodeEndpointBuilder defines functions that build urls for geocode/reverse geocode
type GeocodeEndpointBuilder interface {
	GeocodeUrl(string) string
	ReverseGeocodeUrl(Location) string
}

// GeocodeResponseParser defines functions that parse response of geocode/reverse geocode
type GeocodeResponseParser interface {
	Location([]byte) Location
	Address([]byte) string
}

// Geocoder has GeocodeEndpointBuilder and GeocodeResponseParser
type Geocoder struct {
	GeocodeEndpointBuilder
	GeocodeResponseParser
}

// Geocode returns location for address
func (g Geocoder) Geocode(address string) (Location, error) {
	ch := make(chan Location)
	go func() {
		ch <- g.Location(responseData(g.GeocodeUrl(url.QueryEscape(address))))
	}()

	select {
	case location := <-ch:
		return location, anyError(location)
	case <-time.After(timeoutInSeconds):
		return Location{}, TimeoutError
	}
}

// ReverseGeocode returns address for location
func (g Geocoder) ReverseGeocode(lat, lng float64) (string, error) {
	ch := make(chan string)
	go func() {
		ch <- g.Address(responseData(g.ReverseGeocodeUrl(Location{lat, lng})))
	}()

	select {
	case address := <-ch:
		return address, anyError(address)
	case <-time.After(timeoutInSeconds):
		return "", TimeoutError
	}
}

// ResponseData gets response from url
func responseData(url string) []byte {
	if request, err := http.NewRequest("GET", url, nil); err == nil {
		if response, err := (&http.Client{}).Do(request); err == nil {
			if data, err := ioutil.ReadAll(response.Body); err == nil {
				return data
			}
		}
	}
	return nil
}

func anyError(v interface{}) (err error) {
	switch v := v.(type) {
	case Location:
		if v.Lat == 0 && v.Lng == 0 {
			return NoResultError
		}
	case string:
		if v == "" {
			return NoResultError
		}
	}
	return
}
