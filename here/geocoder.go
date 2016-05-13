// Package here is a geo-golang based HERE geocode/reverse geocode client
package here

import (
	"fmt"
	"github.com/codingsince1985/geo-golang"
)

type baseURL struct {
	forGeocode, forReverseGeocode string
}

type geocodeResponse struct {
	Response struct {
		View []struct {
			Result []struct {
				Location struct {
					DisplayPosition struct {
						Latitude, Longitude float64
					}
					Address struct {
						Label string
					}
				}
			}
		}
	}
}

var r = 100

// Geocoder constructs HERE geocoder
func Geocoder(id, code string, radius int, baseURLs ...string) geo.Geocoder {
	if radius > 0 {
		r = radius
	}
	var url1, url2 string
	if len(baseURLs) > 0 {
		url1 = baseURLs[0]
		url2 = baseURLs[0]
	} else {
		p := "gen=8&app_id=" + id + "&app_code=" + code
		url1 = "http://geocoder.cit.api.here.com/6.2/geocode.json?" + p
		url2 = "http://reverse.geocoder.cit.api.here.com/6.2/reversegeocode.json?mode=retrieveAddresses&" + p
	}
	return geo.HTTPGeocoder{
		EndpointBuilder:       baseURL{url1, url2},
		ResponseParserFactory: func() geo.ResponseParser { return &geocodeResponse{} },
	}
}

func (b baseURL) GeocodeURL(address string) string {
	return b.forGeocode + "&searchtext=" + address
}

func (b baseURL) ReverseGeocodeURL(l geo.Location) string {
	return b.forReverseGeocode + fmt.Sprintf("&prox=%f,%f,%d", l.Lat, l.Lng, r)
}

func (r *geocodeResponse) Location() (l geo.Location) {
	if len(r.Response.View) > 0 {
		p := r.Response.View[0].Result[0].Location.DisplayPosition
		l = geo.Location{p.Latitude, p.Longitude}
	}
	return
}

func (r *geocodeResponse) Address() (address string) {
	if len(r.Response.View) > 0 {
		address = r.Response.View[0].Result[0].Location.Address.Label
	}
	return
}
