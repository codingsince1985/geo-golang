// Package opencage is a geo-golang based OpenCage geocode/reverse geocode client
package opencage

import (
	"fmt"
	"strings"

	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/osm"
)

type (
	baseURL string

	geocodeResponse struct {
		Results []struct {
			Formatted  string
			Geometry   geo.Location
			Components osm.Address
		}
		Status struct {
			Code    int
			Message string
		}
	}
)

// Geocoder constructs OpenCage geocoder
func Geocoder(key string, baseURLs ...string) geo.Geocoder {
	return geo.HTTPGeocoder{
		EndpointBuilder:       baseURL(getURL(key, baseURLs...)),
		ResponseParserFactory: func() geo.ResponseParser { return &geocodeResponse{} },
	}
}

func getURL(key string, baseURLs ...string) string {
	if len(baseURLs) > 0 {
		return baseURLs[0]
	}
	return "http://api.opencagedata.com/geocode/v1/json?key=" + key + "&q="
}

func (b baseURL) GeocodeURL(address string) string { return string(b) + address }

func (b baseURL) ReverseGeocodeURL(l geo.Location) string {
	return string(b) + fmt.Sprintf("%+f,%+f", l.Lat, l.Lng)
}

func (r *geocodeResponse) Location() (*geo.Location, error) {
	if r.Status.Code >= 400 {
		return nil, fmt.Errorf("geocoding error: %s", r.Status.Message)
	}
	if len(r.Results) == 0 {
		return nil, nil
	}

	return &geo.Location{
		Lat: r.Results[0].Geometry.Lat,
		Lng: r.Results[0].Geometry.Lng,
	}, nil
}

func (r *geocodeResponse) Address() (*geo.Address, error) {
	if r.Status.Code >= 400 {
		return nil, fmt.Errorf("geocoding error: %s", r.Status.Message)
	}
	if len(r.Results) == 0 {
		return nil, nil
	}

	addr := r.Results[0].Components

	locality := addr.Locality()
	if locality == "" {
		locality = addr.Suburb
	}
	return &geo.Address{
		FormattedAddress: r.Results[0].Formatted,
		HouseNumber:      addr.HouseNumber,
		Street:           addr.Street(),
		Suburb:           addr.Suburb,
		Postcode:         addr.Postcode,
		City:             locality,
		CountryCode:      strings.ToUpper(addr.CountryCode),
		Country:          addr.Country,
		County:           addr.County,
		State:            addr.State,
		StateDistrict:    addr.StateDistrict,
	}, nil
}
