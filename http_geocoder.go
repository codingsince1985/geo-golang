package geo

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Default timeout for the request execution
const DefaultTimeout = time.Second * 8

// ErrNoResult occurs when no result returned
var ErrNoResult = errors.New("NO_RESULT")
var ErrTimeout = errors.New("TIMEOUT")

// Location is the output of Geocode
type Location struct {
	Lat, Lng float64
}

// Address is returned by ReverseGeocode.
// This is a structured representation of an address, including its flat representation
type Address struct {
	FormattedAddress string
	Street           string
	HouseNumber      string
	Suburb           string
	Postcode         string
	State            string
	StateDistrict    string
	County           string
	Country          string
	CountryCode      string
	City             string
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
	Location() (*Location, error)
	Address() (*Address, error)
}

// HTTPGeocoder has EndpointBuilder and ResponseParser
type HTTPGeocoder struct {
	EndpointBuilder
	ResponseParserFactory
}

// Geocode returns location for address
func (g HTTPGeocoder) Geocode(address string) (*Location, error) {
	responseParser := g.ResponseParserFactory()

	ctx, cancel := context.WithTimeout(context.TODO(), DefaultTimeout)
	defer cancel()

	ch := make(chan struct {
		l *Location
		e error
	}, 1)
	go func() {
		if err := response(ctx, g.GeocodeURL(url.QueryEscape(address)), responseParser); err != nil {
			ch <- struct {
				l *Location
				e error
			}{
				l: nil,
				e: err,
			}
		}

		loc, err := responseParser.Location()
		ch <- struct {
			l *Location
			e error
		}{
			l: loc,
			e: err,
		}
	}()

	select {
	case <-ctx.Done():
		return nil, ErrTimeout
	case res := <-ch:
		return res.l, res.e
	}
}

// ReverseGeocode returns address for location
func (g HTTPGeocoder) ReverseGeocode(lat, lng float64) (*Address, error) {
	responseParser := g.ResponseParserFactory()

	ctx, cancel := context.WithTimeout(context.TODO(), DefaultTimeout)
	defer cancel()

	ch := make(chan struct {
		a *Address
		e error
	}, 1)
	go func() {
		if err := response(ctx, g.ReverseGeocodeURL(Location{lat, lng}), responseParser); err != nil {
			ch <- struct {
				a *Address
				e error
			}{
				a: nil,
				e: err,
			}
		}

		addr, err := responseParser.Address()
		ch <- struct {
			a *Address
			e error
		}{
			a: addr,
			e: err,
		}
	}()

	select {
	case <-ctx.Done():
		return nil, ErrTimeout
	case res := <-ch:
		return res.a, res.e
	}
}

// Response gets response from url
func response(ctx context.Context, url string, obj ResponseParser) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	body := strings.Trim(string(data), " []")
	if body == "" {
		return nil
	}
	if err := json.Unmarshal([]byte(body), obj); err != nil {
		return err
	}

	return nil
}

func ParseFloat(value string) float64 {
	f, _ := strconv.ParseFloat(value, 64)
	return f
}
