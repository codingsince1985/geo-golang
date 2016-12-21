// Package opencage is a geo-golang based OpenCage geocode/reverse geocode client
package opencage

import (
	"fmt"
	"github.com/codingsince1985/geo-golang"
)

type (
	baseURL         string
	geocodeResponse struct {
		Results []struct {
			Formatted string
			Geometry  geo.Location
		}
	}
)

// Geocoder constructs OpenCage geocoder
func Geocoder(key string, baseURLs ...string) geo.Geocoder {
	return geo.HTTPGeocoder{
		EndpointBuilder:       baseURL(getUrl(key, baseURLs...)),
		ResponseParserFactory: func() geo.ResponseParser { return &geocodeResponse{} },
	}
}

func getUrl(key string, baseURLs ...string) string {
	if len(baseURLs) > 0 {
		return baseURLs[0]
	}
	return "http://api.opencagedata.com/geocode/v1/json?key=" + key + "&q="
}

func (b baseURL) GeocodeURL(address string) string { return string(b) + address }

func (b baseURL) ReverseGeocodeURL(l geo.Location) string {
	return string(b) + fmt.Sprintf("%+f,%+f", l.Lat, l.Lng)
}

func (r *geocodeResponse) Location() geo.Location {
	if len(r.Results) > 0 {
		return r.Results[0].Geometry
	}
	return geo.Location{}
}

func (r *geocodeResponse) Address() string {
	if len(r.Results) > 0 {
		return r.Results[0].Formatted
	}
	return ""
}
