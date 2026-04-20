// Package ip2geo is a geo-golang based ip2geo.dev IP geolocation client
package ip2geo

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/codingsince1985/geo-golang"
)

type geocoder struct {
	apiKey  string
	baseURL string
}

type apiResponse struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		IP        string `json:"ip"`
		Type      string `json:"type"`
		Continent struct {
			Name    string `json:"name"`
			Code    string `json:"code"`
			Country struct {
				Name        string `json:"name"`
				Code        string `json:"code"`
				PhoneCode   string `json:"phone_code"`
				Capital     string `json:"capital"`
				Subdivision struct {
					Name string `json:"name"`
					Code string `json:"code"`
				} `json:"subdivision"`
				City struct {
					Name           string  `json:"name"`
					Latitude       float64 `json:"latitude"`
					Longitude      float64 `json:"longitude"`
					PostalCode     string  `json:"postal_code"`
					AccuracyRadius int     `json:"accuracy_radius"`
					Timezone       struct {
						Name string `json:"name"`
					} `json:"timezone"`
				} `json:"city"`
			} `json:"country"`
		} `json:"continent"`
		ASN struct {
			Number int    `json:"number"`
			Name   string `json:"name"`
		} `json:"asn"`
		RegisteredCountry struct {
			Name string `json:"name"`
			Code string `json:"code"`
		} `json:"registered_country"`
	} `json:"data"`
}

// Geocoder constructs an ip2geo geocoder
func Geocoder(apiKey string, baseURLs ...string) geo.Geocoder {
	baseURL := "https://api.ip2geo.dev"
	if len(baseURLs) > 0 {
		baseURL = baseURLs[0]
	}
	return &geocoder{apiKey: apiKey, baseURL: baseURL}
}

// Geocode returns location for the given IP address
func (g *geocoder) Geocode(address string) (*geo.Location, error) {
	resp, err := g.fetch(address)
	if err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, errors.New("ip2geo: " + resp.Message)
	}

	city := resp.Data.Continent.Country.City
	if city.Latitude == 0 && city.Longitude == 0 {
		return nil, nil
	}

	return &geo.Location{
		Lat: city.Latitude,
		Lng: city.Longitude,
	}, nil
}

// ReverseGeocode returns address for location.
// ip2geo is an IP geolocation service and does not support reverse geocoding.
func (g *geocoder) ReverseGeocode(lat, lng float64) (*geo.Address, error) {
	return nil, errors.New("ip2geo: reverse geocoding is not supported")
}

func (g *geocoder) fetch(ip string) (*apiResponse, error) {
	reqURL := g.baseURL + "/convert?ip=" + url.QueryEscape(ip)

	ctx, cancel := context.WithTimeout(context.Background(), geo.DefaultTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Api-Key", g.apiKey)
	req.Header.Set("User-Agent", "geo-golang/1.0")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result apiResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
