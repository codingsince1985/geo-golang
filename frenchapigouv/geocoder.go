// Package frenchapigouv is a geo-golang based French API Gouv geocode/reverse geocode client
package frenchapigouv

import (
	"fmt"
	"strings"

	"github.com/codingsince1985/geo-golang"
)

type (
	baseURL         string
	geocodeResponse struct {
		Type     string
		Version  string
		Features []struct {
			Type     string
			Geometry struct {
				Type        string
				Coordinates []float64
			}
			Properties struct {
				Label       string
				Score       float64
				Housenumber string
				Citycode    string
				Context     string
				Postcode    string
				Name        string
				ID          string
				Y           float64
				Importance  float64
				Type        string
				City        string
				X           float64
				Street      string
			}
		}
		Attribution string
		Licence     string
		Query       string
		Limit       int
	}
	context struct {
		state      string
		county     string
		countyCode string
	}
)

// Geocoder constructs FrenchApiGouv geocoder
func Geocoder() geo.Geocoder { return GeocoderWithURL("https://api-adresse.data.gouv.fr/") }

// GeocoderWithURL constructs French API Gouv geocoder using a custom installation of Nominatim
func GeocoderWithURL(url string) geo.Geocoder {
	return geo.HTTPGeocoder{
		EndpointBuilder:       baseURL(url),
		ResponseParserFactory: func() geo.ResponseParser { return &geocodeResponse{} },
	}
}

func (b baseURL) GeocodeURL(address string) string {
	return string(b) + "search?limit=1&q=" + address
}

func (b baseURL) ReverseGeocodeURL(l geo.Location) string {
	return string(b) + "reverse?" + fmt.Sprintf("lat=%f&lon=%f", l.Lat, l.Lng)
}

func (r *geocodeResponse) Location() (*geo.Location, error) {
	if len(r.Features) == 0 || len(r.Features[0].Geometry.Coordinates) < 2 {
		return nil, nil
	}
	p := r.Features[0].Geometry.Coordinates
	return &geo.Location{
		Lat: p[1],
		Lng: p[0],
	}, nil
}

func (r *geocodeResponse) Address() (*geo.Address, error) {
	if len(r.Features) == 0 || r.Features[0].Properties.Label == "baninfo" {
		return nil, nil
	}
	p := r.Features[0].Properties
	c := r.parseContext()

	if p.Type == "street" || p.Type == "locality" {
		p.Street = p.Name
	}
	return &geo.Address{
		FormattedAddress: strings.Join(strings.Fields(strings.Trim(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s", p.Housenumber, p.Street, p.Postcode, p.City, c.county, c.state, "France"), " ,")), " "),
		HouseNumber:      p.Housenumber,
		Street:           p.Street,
		Postcode:         p.Postcode,
		City:             p.City,
		State:            c.state,
		County:           c.county,
		Country:          "France",
		CountryCode:      "FRA",
	}, nil
}

func (r *geocodeResponse) parseContext() *context {
	var c context
	if len(r.Features) > 0 {
		p := r.Features[0].Properties
		f := strings.Split(p.Context, ",")
		for i := range f {
			switch i {
			case 0:
				c.countyCode = f[i]
			case 1:
				c.county = strings.TrimSpace(f[i])
			case 2:
				c.state = strings.TrimSpace(f[i])
			}
		}
	}
	return &c
}
