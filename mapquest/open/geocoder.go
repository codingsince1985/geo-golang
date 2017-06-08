// Package open is a geo-golang based MapRequest Open geocode/reverse geocode client
package open

import (
	"fmt"
	"strings"

	"github.com/codingsince1985/geo-golang"
)

type (
	baseURL         string
	geocodeResponse struct {
		Results []struct {
			Locations []struct {
				LatLng struct {
					Lat float64
					Lng float64
				}
				PostalCode string
				Street     string
				AdminArea6 string // neighbourhood
				AdminArea5 string // city
				AdminArea4 string // county
				AdminArea3 string // state
				AdminArea1 string // country (ISO 3166-1 alpha-2 code)
			}
		}
	}
)

// Geocoder constructs MapRequest Open geocoder
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
	return "http://open.mapquestapi.com/geocoding/v1/*?key=" + key + "&location="
}

func (b baseURL) GeocodeURL(address string) string {
	return strings.Replace(string(b), "*", "address", 1) + address
}

func (b baseURL) ReverseGeocodeURL(l geo.Location) string {
	return strings.Replace(string(b), "*", "reverse", 1) + fmt.Sprintf("%f,%f", l.Lat, l.Lng)
}

func (r *geocodeResponse) Location() (*geo.Location, error) {
	if len(r.Results) == 0 || len(r.Results[0].Locations) == 0 {
		return nil, nil
	}

	loc := r.Results[0].Locations[0].LatLng
	return &geo.Location{
		Lat: loc.Lat,
		Lng: loc.Lng,
	}, nil
}

func (r *geocodeResponse) Address() (*geo.Address, error) {
	if len(r.Results) == 0 || len(r.Results[0].Locations) == 0 {
		return nil, nil
	}

	p := r.Results[0].Locations[0]
	if p.Street == "" || p.AdminArea5 == "" {
		return nil, nil
	}

	formattedAddress := p.Street + ", " + p.PostalCode + ", " + p.AdminArea5 + ", " + p.AdminArea3 + ", " + p.AdminArea1
	return &geo.Address{
		FormattedAddress: formattedAddress,
		Street:           p.Street,
		Suburb:           p.AdminArea6,
		Postcode:         p.PostalCode,
		City:             p.AdminArea5,
		County:           p.AdminArea4,
		State:            p.AdminArea3,
		CountryCode:      p.AdminArea1,
	}, nil
}
