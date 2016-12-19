// Package locationiq is a geo-golang based LocationIQ geocode/reverse geocode client
package locationiq

import (
	"fmt"
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
	Suburb        string `json:"suburb"`
	City          string `json:"city"`
	County        string `json:"county"`
	Country       string `json:"country"`
	CountryCode   string `json:"country_code"`
	Road          string `json:"road"`
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

<<<<<<< HEAD
func (r *geocodeResponse) Location() geo.Location {
	l := geo.Location{}
	// In case of empty response from LocationIQ or any other error we get zero values
	if r.Lat != "" && r.Lon != "" {
		l.Lat = geo.ParseFloat(r.Lat)
		l.Lng = geo.ParseFloat(r.Lon)
=======
func (r *geocodeResponse) Location() (*geo.Location, error) {
	if r.Lat == "" || r.Lon == "" {
		return nil, fmt.Errorf("empty lat/lon value: %s", r.Error)
>>>>>>> Return pointer instead of value; also return error
	}

	lat, err := parseFloat(r.Lat)
	if err != nil {
		return nil, fmt.Errorf("error converting lat to float: %s", err)
	}
	lon, err := parseFloat(r.Lon)
	if err != nil {
		return nil, fmt.Errorf("error converting lon to float: %s", err)
	}

	return &geo.Location{
		Lat: lat,
		Lng: lon,
	}, nil
}

func (r *geocodeResponse) Address() (*geo.Address, error) {
	if r.Error != "" {
		return nil, fmt.Errorf("error reverse geocoding: %s", r.Error)
	}
	return &geo.Address{
		FormattedAddress: r.DisplayName,
		Street:           r.Addr.Road,
		HouseNumber:      r.Addr.HouseNumber,
		City:             r.Addr.City,
		Postcode:         r.Addr.Postcode,
		Country:          r.Addr.Country,
		CountryCode:      r.Addr.CountryCode,
	}, nil
}

func (r *geocodeResponse) FormattedAddress() string {
	if r.Error != "" {
		return ""
	}
	return r.DisplayName
}
<<<<<<< HEAD
=======

func parseFloat(value string) (float64, error) {
	return strconv.ParseFloat(value, 64)
}
>>>>>>> Return pointer instead of value; also return error
