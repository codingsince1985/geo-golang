// Package locationiq is a geo-golang based LocationIQ geocode/reverse geocode client
package locationiq

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/codingsince1985/geo-golang"
)

type baseURL string

type geocodeResponse struct {
	DisplayName     string `json:"display_name"`
	Lat, Lon, Error string
	Addr            locationiqAddress `json:"address"`
}

type locationiqAddress struct {
	HouseNumber   string `json:"house_number"`
	Road          string `json:"road"`
	Pedestrian    string `json:"pedestrian"`
	Cycleway      string `json:"cycleway"`
	Highway       string `json:"highway"`
	Path          string `json:"path"`
	Suburb        string `json:"suburb"`
	City          string `json:"city"`
	Town          string `json:"town"`
	Village       string `json:"village"`
	County        string `json:"county"`
	Country       string `json:"country"`
	CountryCode   string `json:"country_code"`
	State         string `json:"state"`
	StateDistrict string `json:"state_district"`
	Postcode      string `json:"postcode"`
}

const (
	defaultURL  = "http://locationiq.org/v1/"
	minZoom     = 0  // Min zoom level for locationiq - country level
	maxZoom     = 18 // Max zoom level for locationiq - house level
	defaultZoom = 18
)

var (
	key  string
	zoom int
)

// Geocoder constructs LocationIQ geocoder
func Geocoder(k string, z int, baseURLs ...string) geo.Geocoder {
	key = k

	var url string
	if len(baseURLs) > 0 {
		url = baseURLs[0]
	} else {
		url = defaultURL
	}

	if z > minZoom && z <= maxZoom {
		zoom = z
	} else {
		zoom = defaultZoom
	}

	return geo.HTTPGeocoder{
		EndpointBuilder:       baseURL(url),
		ResponseParserFactory: func() geo.ResponseParser { return &geocodeResponse{} },
	}
}

func (b baseURL) GeocodeURL(address string) string {
	return string(b) + "search.php?key=" + key + "&format=json&limit=1&q=" + address
}

func (b baseURL) ReverseGeocodeURL(l geo.Location) string {
	return string(b) + "reverse.php?key=" + key + fmt.Sprintf("&format=json&lat=%f&lon=%f&zoom=%d", l.Lat, l.Lng, zoom)
}

func (r *geocodeResponse) Location() (*geo.Location, error) {
	if r.Error != "" {
		return nil, fmt.Errorf("geocoding error: %s", r.Error)
	}
	if r.Lat == "" || r.Lon == "" {
		return nil, fmt.Errorf("empty lat/lon value: %s", r.Error)
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
		Street:           extractStreet(r.Addr),
		HouseNumber:      r.Addr.HouseNumber,
		City:             extractLocality(r.Addr),
		Postcode:         r.Addr.Postcode,
		Suburb:           r.Addr.Suburb,
		State:            r.Addr.State,
		Country:          r.Addr.Country,
		CountryCode:      strings.ToUpper(r.Addr.CountryCode),
	}, nil
}

func (r *geocodeResponse) FormattedAddress() string {
	if r.Error != "" {
		return ""
	}
	return r.DisplayName
}

func extractLocality(a locationiqAddress) string {
	var locality string

	if a.City != "" {
		locality = a.City
	} else if a.Town != "" {
		locality = a.Town
	} else if a.Village != "" {
		locality = a.Village
	}

	return locality
}

func extractStreet(a locationiqAddress) string {
	var street string

	if a.Road != "" {
		street = a.Road
	} else if a.Pedestrian != "" {
		street = a.Pedestrian
	} else if a.Path != "" {
		street = a.Path
	} else if a.Cycleway != "" {
		street = a.Cycleway
	} else if a.Highway != "" {
		street = a.Highway
	}

	return street
}
