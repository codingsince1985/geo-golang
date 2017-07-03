// Package tomtom is a geo-golang based TomTom geocode/reverse geocode client
package tomtom

import (
	"fmt"
	"strings"

	geo "github.com/codingsince1985/geo-golang"
)

type (
	baseURL         string
	geocodeResponse struct {
		Summary struct {
			Query string
		}

		Results []struct {
			Position struct {
				Lat float64
				Lon float64
			}
		}

		// Reverse Geocoding response
		Addresses []struct {
			Address struct {
				BuildingNumber              string
				StreetNumber                string
				Street                      string
				StreetName                  string
				StreetNameAndNumber         string
				CountryCode                 string
				CountrySubdivision          string // state code
				CountrySecondarySubdivision string
				CountryTertiarySubdivision  string
				Municipality                string // city
				PostalCode                  string
				Country                     string
				CountryCodeISO3             string
				FreeformAddress             string
				CountrySubdivisionName      string
			}
		}
	}
)

// Geocoder constructs TomTom geocoder
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
	return "https://api.tomtom.com/search/2/*?key=" + key
}

func (b baseURL) GeocodeURL(address string) string {
	params := fmt.Sprintf("geocode/%s.json", address)
	return strings.Replace(string(b), "*", params, 1)
}

func (b baseURL) ReverseGeocodeURL(l geo.Location) string {
	params := fmt.Sprintf("reverseGeocode/%f,%f", l.Lat, l.Lng)
	return strings.Replace(string(b), "*", params, 1)
}

func (r *geocodeResponse) Location() (*geo.Location, error) {
	if len(r.Results) > 0 {
		p := r.Results[0].Position

		return &geo.Location{
			Lat: p.Lat,
			Lng: p.Lon,
		}, nil
	}
	return nil, nil
}

func (r *geocodeResponse) Address() (*geo.Address, error) {
	if len(r.Addresses) > 0 {
		a := r.Addresses[0].Address
		return &geo.Address{
			FormattedAddress: a.FreeformAddress,
			Street:           a.StreetName,
			HouseNumber:      a.StreetNumber,
			City:             a.Municipality,
			Postcode:         a.PostalCode,
			State:            a.CountrySubdivision,
			Country:          a.Country,
			CountryCode:      a.CountryCode,
		}, nil
	}
	return nil, nil
}
