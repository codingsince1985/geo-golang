// Package google is a geo-golang based Google Geo Location API
// https://developers.google.com/maps/documentation/geocoding/intro
package google

import (
	"fmt"

	"github.com/codingsince1985/geo-golang"
)

type (
	baseURL         string
	geocodeResponse struct {
		Results []struct {
			FormattedAddress  string                   `json:"formatted_address"`
			AddressComponents []googleAddressComponent `json:"address_components"`
			Geometry          struct {
				Location geo.Location
			}
		}
		Status string `json:"status"`
	}
	googleAddressComponent struct {
		LongName  string   `json:"long_name"`
		ShortName string   `json:"short_name"`
		Types     []string `json:"types"`
	}
)

const (
	statusOK                   = "OK"
	statusNoResults            = "ZERO_RESULTS"
	componentTypeHouseNumber   = "street_number"
	componentTypeStreetName    = "route"
	componentTypeSuburb        = "sublocality"
	componentTypeLocality      = "locality"
	componentTypeStateDistrict = "administrative_area_level_2"
	componentTypeState         = "administrative_area_level_1"
	componentTypeCountry       = "country"
	componentTypePostcode      = "postal_code"
)

// Geocoder constructs Google geocoder
func Geocoder(apiKey string, baseURLs ...string) geo.Geocoder {
	return geo.HTTPGeocoder{
		EndpointBuilder:       baseURL(getURL(apiKey, baseURLs...)),
		ResponseParserFactory: func() geo.ResponseParser { return &geocodeResponse{} },
	}
}

func getURL(apiKey string, baseURLs ...string) string {
	if len(baseURLs) > 0 {
		return baseURLs[0]
	}
	return fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?key=%s&", apiKey)
}

func (b baseURL) GeocodeURL(address string) string { return string(b) + "address=" + address }

func (b baseURL) ReverseGeocodeURL(l geo.Location) string {
	return string(b) + fmt.Sprintf("result_type=street_address&latlng=%f,%f", l.Lat, l.Lng)
}

func (r *geocodeResponse) Location() (*geo.Location, error) {
	if r.Status == statusNoResults {
		return nil, nil
	} else if r.Status != statusOK {
		return nil, fmt.Errorf("geocoding error: %s", r.Status)
	}

	return &r.Results[0].Geometry.Location, nil
}

func (r *geocodeResponse) Address() (*geo.Address, error) {
	if r.Status == statusNoResults {
		return nil, nil
	} else if r.Status != statusOK {
		return nil, fmt.Errorf("reverse geocoding error: %s", r.Status)
	}

	if len(r.Results) == 0 || len(r.Results[0].AddressComponents) == 0 {
		return nil, nil
	}

	addr := parseGoogleResult(r)

	return addr, nil
}

func parseGoogleResult(r *geocodeResponse) *geo.Address {
	addr := &geo.Address{}
	res := r.Results[0]
	addr.FormattedAddress = res.FormattedAddress
OuterLoop:
	for _, comp := range res.AddressComponents {
		for _, typ := range comp.Types {
			switch typ {
			case componentTypeHouseNumber:
				addr.HouseNumber = comp.LongName
				continue OuterLoop
			case componentTypeStreetName:
				addr.Street = comp.LongName
				continue OuterLoop
			case componentTypeSuburb:
				addr.Suburb = comp.LongName
				continue OuterLoop
			case componentTypeLocality:
				addr.City = comp.LongName
				continue OuterLoop
			case componentTypeStateDistrict:
				addr.StateDistrict = comp.LongName
				continue OuterLoop
			case componentTypeState:
				addr.State = comp.LongName
				addr.StateCode = comp.ShortName
				continue OuterLoop
			case componentTypeCountry:
				addr.Country = comp.LongName
				addr.CountryCode = comp.ShortName
				continue OuterLoop
			case componentTypePostcode:
				addr.Postcode = comp.LongName
				continue OuterLoop
			}
		}
	}

	return addr
}
