// Package here is a geo-golang based HERE geocode/reverse geocode client
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
							Label string
						}
					}
				}
			}
		}
	}
)

var r = 100

// Geocoder constructs HERE geocoder
func Geocoder(id, code string, radius int, baseURLs ...string) geo.Geocoder {
	if radius > 0 {
		r = radius
	}
	p := "gen=8&app_id=" + id + "&app_code=" + code
	return geo.HTTPGeocoder{
		EndpointBuilder: baseURL{
			getGeocodeUrl(p, baseURLs...),
			getReverseGeocodeUrl(p, baseURLs...)},
		ResponseParserFactory: func() geo.ResponseParser { return &geocodeResponse{} },
	}
}

func getGeocodeUrl(p string, baseURLs ...string) string {
	if len(baseURLs) > 0 {
		return baseURLs[0]
	}
	return "http://geocoder.cit.api.here.com/6.2/geocode.json?" + p
}

func getReverseGeocodeUrl(p string, baseURLs ...string) string {
	if len(baseURLs) > 0 {
		return baseURLs[0]
	}
	return "http://reverse.geocoder.cit.api.here.com/6.2/reversegeocode.json?mode=retrieveAddresses&" + p
}

func (b baseURL) GeocodeURL(address string) string { return b.forGeocode + "&searchtext=" + address }

func (b baseURL) ReverseGeocodeURL(l geo.Location) string {
	return b.forReverseGeocode + fmt.Sprintf("&prox=%f,%f,%d", l.Lat, l.Lng, r)
}

func (r *geocodeResponse) Location() geo.Location {
	if len(r.Response.View) > 0 {
		p := r.Response.View[0].Result[0].Location.DisplayPosition
		return geo.Location{p.Latitude, p.Longitude}
	}
	return geo.Location{}
}

func (r *geocodeResponse) Address() string {
	if len(r.Response.View) > 0 {
		return r.Response.View[0].Result[0].Location.Address.Label
	}
	return ""
}
