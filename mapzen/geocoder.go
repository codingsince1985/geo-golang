// Package mapzen is a geo-golang based Mapzen geocode/reverse client
package mapzen

import (
	"fmt"
	"strings"

	geo "github.com/codingsince1985/geo-golang"
)

type (
	baseURL         string
	geocodeResponse struct {
		Geocoding struct {
			Query struct {
				Text string
			}
		}

		Features []struct {
			Geometry struct {
				Coordinates []float64
			}
			Properties struct {
				Name        string
				HouseNumber string
				Street      string
				PostalCode  string
				Country     string
				CountryCode string `json:"country_a"`
				Region      string
				RegionCode  string `json:"region_a"`
				County      string
				Label       string
			}
		}
	}
)

// Geocoder constructs Mapzen geocoder
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
	return "https://search.mapzen.com/v1/*" + "&api_key=" + key
}

func (b baseURL) GeocodeURL(address string) string {
	params := fmt.Sprintf("search?size=%d&text=%s", 1, address)
	return strings.Replace(string(b), "*", params, 1)
}

func (b baseURL) ReverseGeocodeURL(l geo.Location) string {
	params := fmt.Sprintf("reverse?size=%d&point.lat=%f&point.lon=%f", 1, l.Lat, l.Lng)
	return strings.Replace(string(b), "*", params, 1)
}

func (r *geocodeResponse) Location() (*geo.Location, error) {
	if len(r.Features) == 0 {
		return nil, nil
	}

	pt := r.Features[0].Geometry.Coordinates
	if len(pt) == 0 {
		return nil, nil
	}

	return &geo.Location{Lat: pt[1], Lng: pt[0]}, nil
}

func (r *geocodeResponse) Address() (*geo.Address, error) {
	if len(r.Features) == 0 {
		return nil, nil
	}

	props := r.Features[0].Properties
	addr := &geo.Address{
		FormattedAddress: props.Label,
		Street:           props.Street,
		HouseNumber:      props.HouseNumber,
		Postcode:         props.PostalCode,
		Country:          props.Country,
		CountryCode:      props.CountryCode,
		State:            props.Region,
		StateCode:        props.RegionCode,
	}
	return addr, nil
}
