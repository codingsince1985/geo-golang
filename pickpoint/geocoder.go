// Package pickpoint is a geo-golang based PickPoint geocode/reverse geocode client
package pickpoint

import (
	"fmt"
	"strings"

	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/osm"
)

type (
	baseURL         string
	geocodeResponse struct {
		DisplayName string `json:"display_name"`
		Lat         string
		Lon         string
		Error       string
		Addr        osm.Address `json:"address"`
	}
)

var key string

// Geocoder constructs PickPoint geocoder
func Geocoder(apiKey string, baseURLs ...string) geo.Geocoder {
	key = apiKey
	return geo.HTTPGeocoder{
		EndpointBuilder:       baseURL(getURL(baseURLs...)),
		ResponseParserFactory: func() geo.ResponseParser { return &geocodeResponse{} },
	}
}

func getURL(baseURLs ...string) string {
	if len(baseURLs) > 0 {
		return baseURLs[0]
	}
	return "https://api.pickpoint.io/v1"
}

func (b baseURL) GeocodeURL(address string) string {
	return string(b) + fmt.Sprintf("/forward?key=%s&limit=1&q=%s", key, address)
}

func (b baseURL) ReverseGeocodeURL(l geo.Location) string {
	return string(b) + fmt.Sprintf("/reverse?key=%s&lat=%f&lon=%f", key, l.Lat, l.Lng)
}

func (r *geocodeResponse) Location() (*geo.Location, error) {
	if r.Error != "" {
		return nil, fmt.Errorf("geocoding error: %s", r.Error)
	}
	if r.Lat == "" && r.Lon == "" {
		return nil, nil
	}

	return &geo.Location{
		Lat: geo.ParseFloat(r.Lat),
		Lng: geo.ParseFloat(r.Lon),
	}, nil
}

func (r *geocodeResponse) Address() (*geo.Address, error) {
	if r.Error != "" {
		return nil, fmt.Errorf("reverse geocoding error: %s", r.Error)
	}

	return &geo.Address{
		FormattedAddress: r.DisplayName,
		HouseNumber:      r.Addr.HouseNumber,
		Street:           r.Addr.Street(),
		Postcode:         r.Addr.Postcode,
		City:             r.Addr.Locality(),
		Suburb:           r.Addr.Suburb,
		State:            r.Addr.State,
		Country:          r.Addr.Country,
		CountryCode:      strings.ToUpper(r.Addr.CountryCode),
	}, nil
}
