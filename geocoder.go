// Package geo is a generic framework to develop geocode/reverse geocode clients
package geo

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ErrTimeout occurs when no response returned within timeoutInSeconds
var ErrTimeout = errors.New("TIMEOUT")
var timeoutInSeconds = time.Second * 8

// ErrNoResult occurs when no result returned
var ErrNoResult = errors.New("NO_RESULT")

// Location is the output of Geocode
type Location struct {
	Lat, Lng float64
}

// EndpointBuilder defines functions that build urls for geocode/reverse geocode
type EndpointBuilder interface {
	GeocodeURL(string) string
	ReverseGeocodeURL(Location) string
}

// ResponseParser defines functions that parse response of geocode/reverse geocode
type ResponseParser interface {
	Location() Location
	Address() string
	ResponseObject() ResponseParser
}

// Geocoder has EndpointBuilder and ResponseParser
type Geocoder struct {
	EndpointBuilder
	ResponseParser
}

// Geocode returns location for address
func (g Geocoder) Geocode(address string) (Location, error) {
	ch := make(chan Location, 1)
	go func() {
		response(g.GeocodeURL(url.QueryEscape(address)), g.ResponseObject())
		ch <- g.Location()
	}()

	select {
	case location := <-ch:
		return location, anyError(location)
	case <-time.After(timeoutInSeconds):
		return Location{}, ErrTimeout
	}
}

// ReverseGeocode returns address for location
func (g Geocoder) ReverseGeocode(lat, lng float64) (string, error) {
	ch := make(chan string, 1)
	go func() {
		response(g.ReverseGeocodeURL(Location{lat, lng}), g.ResponseObject())
		ch <- g.Address()
	}()

	select {
	case address := <-ch:
		return address, anyError(address)
	case <-time.After(timeoutInSeconds):
		return "", ErrTimeout
	}
}

// Response gets response from url
func response(url string, obj ResponseParser) {
	if req, err := http.NewRequest("GET", url, nil); err == nil {
		if resp, err := (&http.Client{}).Do(req); err == nil {
			defer resp.Body.Close()
			if data, err := ioutil.ReadAll(resp.Body); err == nil {
				json.Unmarshal([]byte(strings.Trim(string(data), " []")), obj)
			}
		}
	}
}

func anyError(v interface{}) (err error) {
	switch v := v.(type) {
	case Location:
		if v.Lat == 0 && v.Lng == 0 {
			return ErrNoResult
		}
	case string:
		if v == "" {
			return ErrNoResult
		}
	}
	return
}
