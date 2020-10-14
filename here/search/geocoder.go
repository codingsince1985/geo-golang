// Package search is a geo-golang based HERE geocode/reverse geocode client for the Geocoding and Search API
package search

import (
	"fmt"
	"net/url"

	"github.com/codingsince1985/geo-golang"
)

type (
	baseURL         struct{ forGeocode, forReverseGeocode string }
	geocodeResponse struct {
		Items []struct {
			Address struct {
				Label       string
				CountryCode string
				CountryName string
				StateCode   string
				State       string
				County      string
				District    string
				City        string
				Street      string
				PostalCode  string
				HouseNumber string
			}
			Position struct {
				Lat float64
				Lng float64
			}
		}
	}
)

// Geocoder constructs HERE geocoder
func Geocoder(apiKey string, baseURLs ...string) geo.Geocoder {
	p := "apiKey=" + url.QueryEscape(apiKey)
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
	return "https://geocode.search.hereapi.com/v1/geocode?" + p
}

func getReverseGeocodeURL(p string, baseURLs ...string) string {
	if len(baseURLs) > 0 {
		return baseURLs[0]
	}
	return "https://revgeocode.search.hereapi.com/v1/revgeocode?" + p
}

func (b baseURL) GeocodeURL(address string) string { return b.forGeocode + "&limit=1&q=" + address }

func (b baseURL) ReverseGeocodeURL(l geo.Location) string {
	return b.forReverseGeocode + fmt.Sprintf("&limit=1&at=%f,%f", l.Lat, l.Lng)
}

func (r *geocodeResponse) Location() (*geo.Location, error) {
	if len(r.Items) == 0 {
		return nil, nil
	}
	p := r.Items[0].Position
	return &geo.Location{
		Lat: p.Lat,
		Lng: p.Lng,
	}, nil
}

func (r *geocodeResponse) Address() (*geo.Address, error) {
	if len(r.Items) == 0 {
		return nil, nil
	}

	res := r.Items[0].Address

	addr := &geo.Address{
		FormattedAddress: res.Label,
		City:             res.City,
		Street:           res.Street,
		HouseNumber:      res.HouseNumber,
		Postcode:         res.PostalCode,
		State:            res.State,
		County:           res.County,
		Country:          res.CountryName,
		CountryCode:      res.CountryCode,
		Suburb:           res.District,
	}
	return addr, nil
}
