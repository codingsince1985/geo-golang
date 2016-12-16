package locationiq

import (
	"fmt"

	"github.com/codingsince1985/geo-golang"
	"strconv"
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

func (r *geocodeResponse) Location() (l geo.Location) {
	if r.Lat != "" && r.Lon != "" {
		l = geo.Location{
			Lat: parseFloat(r.Lat),
			Lng: parseFloat(r.Lon),
		}
	}
	return
}

func (r *geocodeResponse) Address() string {
	if r.Error != "" {
		return ""
	}
	return r.DisplayName
	// TODO fix the interface to return interface{}
	//return r.Address
}

func parseFloat(value string) float64 {
	f, _ := strconv.ParseFloat(value, 64)
	return f
}
