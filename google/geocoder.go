// Package google is a geo-golang based Google Geo Location API
// https://developers.google.com/maps/documentation/geocoding/intro
package google

import (
	"fmt"
	geo "github.com/codingsince1985/geo-golang"
)

type (
	baseURL         string
	geocodeResponse struct {
		Results []struct {
			FormattedAddress string `json:"formatted_address"`
			Geometry         struct {
				Location geo.Location
			}
		}
	}
)

// Geocoder constructs Google geocoder
func Geocoder(apiKey string, baseURLs ...string) geo.Geocoder {
	return geo.HTTPGeocoder{
		EndpointBuilder:       baseURL(getUrl(apiKey, baseURLs...)),
		ResponseParserFactory: func() geo.ResponseParser { return &geocodeResponse{} },
	}
}

func getUrl(apiKey string, baseURLs ...string) string {
	if len(baseURLs) > 0 {
		return baseURLs[0]
	}
	return fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?key=%s&", apiKey)
}

func (b baseURL) GeocodeURL(address string) string { return string(b) + "address=" + address }

func (b baseURL) ReverseGeocodeURL(l geo.Location) string {
	return string(b) + fmt.Sprintf("latlng=%f,%f", l.Lat, l.Lng)
}

func (r *geocodeResponse) Location() geo.Location {
	if len(r.Results) > 0 {
		return r.Results[0].Geometry.Location
	}
	return geo.Location{}
}

func (r *geocodeResponse) Address() string {
	if len(r.Results) > 0 {
		return r.Results[0].FormattedAddress
	}
	return ""
}
