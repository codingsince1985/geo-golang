package amap

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/codingsince1985/geo-golang"
)

type (
	baseURL string
	geocodeResponse struct {
		XMLName  xml.Name `xml:"response"`
		Status   int      `xml:"status"`
		Info     string   `xml:"info"`
		Infocode int      `xml:"infocode"`
		Count    int      `xml:"count"`
		Geocodes []struct {
			FormattedAddress string `xml:"formatted_address"`
			Country          string `xml:"country"`
			Province         string `xml:"province"`
			Citycode         string `xml:"citycode"`
			City             string `xml:"city"`
			District         string `xml:"district"`
			Adcode           string `xml:"adcode"`
			Street           string `xml:"street"`
			Number           string `xml:"number"`
			Location         string `xml:"location"`
			Level            string `xml:"level"`
		} `xml:"geocodes>geocode"`
		Regeocode struct {
			FormattedAddress string `xml:"formatted_address"`
			AddressComponent struct {
				Country      string `xml:"country"`
				Township     string `xml:"township"`
				District     string `xml:"district"`
				Adcode       string `xml:"adcode"`
				Province     string `xml:"province"`
				Citycode     string `xml:"citycode"`
				StreetNumber struct {
					Number    string `xml:"number"`
					Location  string `xml:"location"`
					Direction string `xml:"direction"`
					Distance  string `xml:"distance"`
					Street    string `xml:"street"`
				} `xml:"streetNumber"`
			} `xml:"addressComponent"`
		} `xml:"regeocode"`
	}
)

const (
	statusOK = 1
)

var r = 1000

// Geocoder constructs AMAP geocoder
func Geocoder(key string, radius int, baseURLs ...string) geo.Geocoder {
	if radius > 0 {
		r = radius
	}
	return geo.HTTPGeocoder{
		EndpointBuilder:       baseURL(getURL(key, baseURLs...)),
		ResponseParserFactory: func() geo.ResponseParser { return &geocodeResponse{} },
		ResponseUnmarshaler:   &geo.XMLUnmarshaler{},
	}
}

func getURL(apiKey string, baseURLs ...string) string {
	if len(baseURLs) > 0 {
		return baseURLs[0]
	}
	return fmt.Sprintf("https://restapi.amap.com/v3/geocode/*?key=%s&", apiKey)
}

// GeocodeURL https://restapi.amap.com/v3/geocode/geo?&output=XML&key=APPKEY&address=ADDRESS
func (b baseURL) GeocodeURL(address string) string {
	return strings.Replace(string(b), "*", "geo", 1) + fmt.Sprintf("output=XML&address=%s", address)
}

// ReverseGeocodeURL https://restapi.amap.com/v3/geocode/regeo?output=XML&key=APPKEY&radius=1000&extensions=all&location=31.225696563611,121.49884033194
func (b baseURL) ReverseGeocodeURL(l geo.Location) string {
	return strings.Replace(string(b), "*", "regeo", 1) + fmt.Sprintf("output=XML&location=%f,%f&radius=%d&extensions=all", l.Lng, l.Lat, r)
}

func (r *geocodeResponse) Location() (*geo.Location, error) {
	var location = &geo.Location{}
	if len(r.Geocodes) == 0 {
		return nil, nil
	}
	if r.Status != statusOK {
		return nil, fmt.Errorf("geocoding error: %v", r.Status)
	}
	fmt.Sscanf(string(r.Geocodes[0].Location), "%f,%f", &location.Lng, &location.Lat)
	return location, nil
}

func (r *geocodeResponse) Address() (*geo.Address, error) {
	if r.Status != statusOK {
		return nil, fmt.Errorf("reverse geocoding error: %v", r.Status)
	}

	addr := parseAmapResult(r)

	return addr, nil
}

func parseAmapResult(r *geocodeResponse) *geo.Address {
	addr := &geo.Address{}
	res := r.Regeocode
	addr.FormattedAddress = string(res.FormattedAddress)
	addr.HouseNumber = string(res.AddressComponent.StreetNumber.Number)
	addr.Street = string(res.AddressComponent.StreetNumber.Street)
	addr.Suburb = string(res.AddressComponent.District)
	addr.State = string(res.AddressComponent.Province)
	addr.Country = string(res.AddressComponent.Country)

	if (*addr == geo.Address{}) {
		return nil
	}
	return addr
}
