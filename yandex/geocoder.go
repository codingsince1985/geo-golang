// Package yandex is a geo-golang based Yandex Maps Location API
package yandex

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/codingsince1985/geo-golang"
)

type (
	baseURL         string
	geocodeResponse struct {
		Response struct {
			GeoObjectCollection struct {
				MetaDataProperty struct {
					GeocoderResponseMetaData struct {
						Request string `json:"request"`
						Found   string `json:"found"`
						Results string `json:"results"`
					} `json:"GeocoderResponseMetaData"`
				} `json:"metaDataProperty"`
				FeatureMember []*yandexFeatureMember `json:"featureMember"`
			} `json:"GeoObjectCollection"`
		} `json:"response"`
	}

	yandexFeatureMember struct {
		GeoObject struct {
			MetaDataProperty struct {
				GeocoderMetaData struct {
					Kind      string `json:"kind"`
					Text      string `json:"text"`
					Precision string `json:"precision"`
					Address   struct {
						CountryCode string `json:"country_code"`
						PostalCode  string `json:"postal_code"`
						Formatted   string `json:"formatted"`
						Components  []struct {
							Kind string `json:"kind"`
							Name string `json:"name"`
						} `json:"Components"`
					} `json:"Address"`
				} `json:"GeocoderMetaData"`
			} `json:"metaDataProperty"`
			Description string `json:"description"`
			Name        string `json:"name"`
			BoundedBy   struct {
				Envelope struct {
					LowerCorner string `json:"lowerCorner"`
					UpperCorner string `json:"upperCorner"`
				} `json:"Envelope"`
			} `json:"boundedBy"`
			Point struct {
				Pos string `json:"pos"`
			} `json:"Point"`
		} `json:"GeoObject"`
	}
)

const (
	componentTypeHouseNumber   = "house"
	componentTypeStreetName    = "street"
	componentTypeLocality      = "locality"
	componentTypeStateDistrict = "area"
	componentTypeState         = "province"
	componentTypeCountry       = "country"
)

// Geocoder constructs Yandex geocoder
func Geocoder(apiKey string, baseURLs ...string) geo.Geocoder {
	return geo.HTTPGeocoder{
		EndpointBuilder:       baseURL(getURL(apiKey, baseURLs...)),
		ResponseParserFactory: func() geo.ResponseParser { return &geocodeResponse{} },
	}
}

func getURL(apiKey string, baseURLs ...string) string {
	if len(baseURLs) > 0 {
		return baseURLs[0]
	}
	return fmt.Sprintf("https://geocode-maps.yandex.ru/1.x/?results=1&lang=en_US&format=json&apikey=%s&", apiKey)
}

func (b baseURL) GeocodeURL(address string) string {
	return string(b) + "geocode=" + address
}

func (b baseURL) ReverseGeocodeURL(l geo.Location) string {
	return string(b) + fmt.Sprintf("sco=latlong&geocode=%f,%f", l.Lat, l.Lng)
}

func (r *geocodeResponse) Location() (*geo.Location, error) {
	if r.Response.GeoObjectCollection.MetaDataProperty.GeocoderResponseMetaData.Found == "0" {
		return nil, nil
	}
	if len(r.Response.GeoObjectCollection.FeatureMember) == 0 {
		return nil, nil
	}
	featureMember := r.Response.GeoObjectCollection.FeatureMember[0]
	result := &geo.Location{}
	latLng := strings.Split(featureMember.GeoObject.Point.Pos, " ")
	if len(latLng) > 1 {
		// Yandex return geo coord in format "long lat"
		result.Lat, _ = strconv.ParseFloat(latLng[1], 64)
		result.Lng, _ = strconv.ParseFloat(latLng[0], 64)
	}

	return result, nil
}

func (r *geocodeResponse) Address() (*geo.Address, error) {
	if r.Response.GeoObjectCollection.MetaDataProperty.GeocoderResponseMetaData.Found == "0" {
		return nil, nil
	}
	if len(r.Response.GeoObjectCollection.FeatureMember) == 0 {
		return nil, nil
	}

	return parseYandexResult(r.Response.GeoObjectCollection.FeatureMember[0]), nil
}

func parseYandexResult(r *yandexFeatureMember) *geo.Address {
	addr := &geo.Address{}
	res := r.GeoObject.MetaDataProperty.GeocoderMetaData

	for _, comp := range res.Address.Components {
		switch comp.Kind {
		case componentTypeHouseNumber:
			addr.HouseNumber = comp.Name
			continue
		case componentTypeStreetName:
			addr.Street = comp.Name
			continue
		case componentTypeLocality:
			addr.City = comp.Name
			continue
		case componentTypeStateDistrict:
			addr.StateDistrict = comp.Name
			continue
		case componentTypeState:
			addr.State = comp.Name
			continue
		case componentTypeCountry:
			addr.Country = comp.Name
			continue
		}
	}

	addr.Postcode = res.Address.PostalCode
	addr.CountryCode = res.Address.CountryCode
	addr.FormattedAddress = res.Address.Formatted

	return addr
}
