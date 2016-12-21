// Package bing is a geo-golang based Microsoft Bing geocode/reverse geocode client
package bing

import (
	"fmt"
	"github.com/codingsince1985/geo-golang"
	"strings"
)

type (
	baseURL         string
	geocodeResponse struct {
		ResourceSets []struct {
			Resources []struct {
				Point struct {
					Coordinates []float64
				}
				Address struct {
					FormattedAddress string
				}
			}
		}
	}
)

// Geocoder constructs Bing geocoder
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
	return "http://dev.virtualearth.net/REST/v1/Locations*key=" + key
}

func (b baseURL) GeocodeURL(address string) string {
	return strings.Replace(string(b), "*", "?q="+address+"&", 1)
}

func (b baseURL) ReverseGeocodeURL(l geo.Location) string {
	return strings.Replace(string(b), "*", fmt.Sprintf("/%f,%f?", l.Lat, l.Lng), 1)
}

func (r *geocodeResponse) Location() geo.Location {
	if len(r.ResourceSets) == 0 || len(r.ResourceSets[0].Resources) == 0 {
		return geo.Location{}
	}
	c := r.ResourceSets[0].Resources[0].Point.Coordinates
	return geo.Location{c[0], c[1]}
}

func (r *geocodeResponse) Address() string {
	if len(r.ResourceSets) == 0 || len(r.ResourceSets[0].Resources) == 0 {
		return ""
	}
	return r.ResourceSets[0].Resources[0].Address.FormattedAddress
}
