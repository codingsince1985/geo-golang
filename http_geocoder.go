package geo

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var timeout = time.Second * 8

// ErrTimeout occurs when no response returned within timeoutInSeconds
var ErrTimeout = errors.New("TIMEOUT")

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

// ResponseParserFactory creates a new ResponseParser
type ResponseParserFactory func() ResponseParser

// ResponseParser defines functions that parse response of geocode/reverse geocode
type ResponseParser interface {
	Location() Location
	Address() string
}

// HTTPGeocoder has EndpointBuilder and ResponseParser
type HTTPGeocoder struct {
	EndpointBuilder
	ResponseParserFactory
}

// Geocode returns location for address
func (g HTTPGeocoder) Geocode(address string) (Location, error) {
	ch := make(chan Location, 1)
	go func() {
		responseParser := g.ResponseParserFactory()
		response(g.GeocodeURL(url.QueryEscape(address)), responseParser)
		ch <- responseParser.Location()
	}()

	select {
	case location := <-ch:
		return location, anyError(location)
	case <-time.After(timeout):
		return Location{}, ErrTimeout
	}
}

// ReverseGeocode returns address for location
func (g HTTPGeocoder) ReverseGeocode(lat, lng float64) (string, error) {
	ch := make(chan string, 1)
	go func() {
		responseParser := g.ResponseParserFactory()
		response(g.ReverseGeocodeURL(Location{lat, lng}), responseParser)
		ch <- responseParser.Address()
	}()

	select {
	case address := <-ch:
		return address, anyError(address)
	case <-time.After(timeout):
		return "", ErrTimeout
	}
}

// Response gets response from url
func response(url string, obj ResponseParser) {
	if req, err := http.NewRequest("GET", url, nil); err == nil {
		if resp, err := (&http.Client{}).Do(req); err == nil {
			defer resp.Body.Close()
			if data, err := ioutil.ReadAll(resp.Body); err == nil {
				// TODO: don't swallow json unmarshal errors
				// currently it just treats an empty response as a ErrNoResult which
				// is fine for now but we should have some logging or something to indicate
				// failed json unmarshal
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

func ParseFloat(value string) float64 {
	f, _ := strconv.ParseFloat(value, 64)
	return f
}
