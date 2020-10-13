// Package here is a geo-golang based HERE geocode/reverse geocode client for the legacy geocoder API
package here

import (
	"fmt"

	"github.com/codingsince1985/geo-golang"
)

type (
	baseURL         struct{ forGeocode, forReverseGeocode string }
	geocodeResponse struct {
		Response struct {
			View []struct {
				Result []struct {
					Location struct {
						DisplayPosition struct {
							Latitude, Longitude float64
						}
						Address struct {
							Label          string
							Country        string
							State          string
							County         string
							City           string
							District       string
							Street         string
							HouseNumber    string
							PostalCode     string
							AdditionalData []struct {
								Key   string
								Value string
							}
						}
					}
				}
			}
		}
	}
)

// Key*Name represents constants for geocoding more address detail
const (
	KeyCountryName = "CountryName"
	KeyStateName   = "StateName"
	KeyCountyName  = "CountyName"
)

var r = 100

// Geocoder constructs HERE geocoder
func Geocoder(id, code string, radius int, baseURLs ...string) geo.Geocoder {
	if radius > 0 {
		r = radius
	}
	p := "gen=9&app_id=" + id + "&app_code=" + code
	return geo.HTTPGeocoder{
		EndpointBuilder: baseURL{
			getGeocodeURL(p, baseURLs...),
			getReverseGeocodeURL(p, baseURLs...)},
		ResponseParserFactory: func() geo.ResponseParser { return &geocodeResponse{} },
	}
}

func getGeocodeURL(p string, baseURLs ...string) string {
	if len(baseURLs) > 0 {
		return baseURLs[0]
	}
	return "http://geocoder.api.here.com/6.2/geocode.json?" + p
}

func getReverseGeocodeURL(p string, baseURLs ...string) string {
	if len(baseURLs) > 0 {
		return baseURLs[0]
	}
	return "http://reverse.geocoder.api.here.com/6.2/reversegeocode.json?mode=retrieveAddresses&" + p
}

func (b baseURL) GeocodeURL(address string) string { return b.forGeocode + "&searchtext=" + address }

func (b baseURL) ReverseGeocodeURL(l geo.Location) string {
	return b.forReverseGeocode + fmt.Sprintf("&prox=%f,%f,%d", l.Lat, l.Lng, r)
}

func (r *geocodeResponse) Location() (*geo.Location, error) {
	if len(r.Response.View) == 0 {
		return nil, nil
	}
	p := r.Response.View[0].Result[0].Location.DisplayPosition
	return &geo.Location{
		Lat: p.Latitude,
		Lng: p.Longitude,
	}, nil
}

func (r *geocodeResponse) Address() (*geo.Address, error) {
	if len(r.Response.View) == 0 || len(r.Response.View[0].Result) == 0 {
		return nil, nil
	}

	res := r.Response.View[0].Result[0].Location.Address

	addr := &geo.Address{
		FormattedAddress: res.Label,
		Street:           res.Street,
		HouseNumber:      res.HouseNumber,
		City:             res.City,
		Postcode:         res.PostalCode,
		CountryCode:      res.Country,
	}
	for _, v := range res.AdditionalData {
		switch v.Key {
		case KeyCountryName:
			addr.Country = v.Value
		case KeyStateName:
			addr.State = v.Value
		}
	}
	return addr, nil
}
