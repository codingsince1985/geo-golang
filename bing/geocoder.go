// Package bing is a geo-golang based Microsoft Bing geocode/reverse geocode client
package bing

import (
	"fmt"
	"github.com/codingsince1985/geo-golang"
	"strings"
)

type baseURL string

type geocodeResponse struct {
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

// Geocoder constructs Bing geocoder
func Geocoder(key string) geo.Geocoder {
	return geo.HTTPGeocoder{
		EndpointBuilder:       baseURL("http://dev.virtualearth.net/REST/v1/Locations*key=" + key),
		ResponseParserFactory: func() geo.ResponseParser { return &geocodeResponse{} },
	}
}

func (b baseURL) GeocodeURL(address string) string {
	u := strings.Replace(string(b), "*", "?q="+address+"&", 1)
	return u
}

func (b baseURL) ReverseGeocodeURL(l geo.Location) string {
	return strings.Replace(string(b), "*", fmt.Sprintf("/%f,%f?", l.Lat, l.Lng), 1)
}

func (r *geocodeResponse) Location() (l geo.Location) {
	if len(r.ResourceSets) <= 0 || len(r.ResourceSets[0].Resources) <= 0 {
		return
	}
	c := r.ResourceSets[0].Resources[0].Point.Coordinates
	l = geo.Location{c[0], c[1]}
	return
}

func (r *geocodeResponse) Address() (address string) {
	if len(r.ResourceSets) <= 0 || len(r.ResourceSets[0].Resources) <= 0 {
		return
	}

	address = r.ResourceSets[0].Resources[0].Address.FormattedAddress
	return
}
