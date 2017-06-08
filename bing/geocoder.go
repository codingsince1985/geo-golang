// Package bing is a geo-golang based Microsoft Bing geocode/reverse geocode client
package bing

import (
	"errors"
	"fmt"
	"strings"

	"github.com/codingsince1985/geo-golang"
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
					AddressLine      string
					AdminDistrict    string
					AdminDistrict2   string
					CountryRegion    string
					Locality         string
					PostalCode       string
				}
			}
		}
		ErrorDetails []string
	}
)

// Geocoder constructs Bing geocoder
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
	return "http://dev.virtualearth.net/REST/v1/Locations*key=" + key
}

func (b baseURL) GeocodeURL(address string) string {
	return strings.Replace(string(b), "*", "?q="+address+"&", 1)
}

func (b baseURL) ReverseGeocodeURL(l geo.Location) string {
	return strings.Replace(string(b), "*", fmt.Sprintf("/%f,%f?", l.Lat, l.Lng), 1)
}

func (r *geocodeResponse) Location() (*geo.Location, error) {
	if len(r.ResourceSets) <= 0 || len(r.ResourceSets[0].Resources) <= 0 {
		return nil, nil
	}
	c := r.ResourceSets[0].Resources[0].Point.Coordinates
	return &geo.Location{
		Lat: c[0],
		Lng: c[1],
	}, nil
}

func (r *geocodeResponse) Address() (*geo.Address, error) {
	if len(r.ErrorDetails) > 0 {
		return nil, errors.New(strings.Join(r.ErrorDetails, " "))
	}
	if len(r.ResourceSets) <= 0 || len(r.ResourceSets[0].Resources) <= 0 {
		return nil, nil
	}

	a := r.ResourceSets[0].Resources[0].Address
	return &geo.Address{
		FormattedAddress: a.FormattedAddress,
		Street:           a.AddressLine,
		City:             a.Locality,
		Postcode:         a.PostalCode,
		Country:          a.CountryRegion,
	}, nil
}
