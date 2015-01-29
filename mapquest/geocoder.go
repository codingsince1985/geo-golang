// Package mapquest has the implementation of using MapRequest geocode and reverse geocode, in ~50 LoC.
package mapquest

import (
	"encoding/json"
	"fmt"
	"github.com/codingsince1985/geo-golang"
	"net/url"
	"strconv"
)

const GEOCODE_BASE_URL = "http://open.mapquestapi.com/nominatim/v1/"

type Endpoint geo.Endpoint
type GeocodeResponse map[string]interface{}

var Geocoder = geo.Geocoder{Endpoint{GEOCODE_BASE_URL}, GeocodeResponse{}}

func (e Endpoint) GeocodeUrl(address string) string {
	return e.BaseUrl + "search.php?format=json&q=" + url.QueryEscape(address)
}

func (e Endpoint) ReverseGeocodeUrl(l geo.Location) string {
	return e.BaseUrl + fmt.Sprintf("reverse.php?format=json&lat=%f&lon=%f", l.Lat, l.Lng)
}

func (r GeocodeResponse) Location(data []byte) geo.Location {
	res := []GeocodeResponse{}
	json.Unmarshal(data, &res)
	return geo.Location{parseFloat(res[0]["lat"]), parseFloat(res[0]["lon"])}
}

func parseFloat(value interface{}) float64 {
	f, _ := strconv.ParseFloat(value.(string), 64)
	return f
}

func (r GeocodeResponse) Address(data []byte) string {
	json.Unmarshal(data, &r)
	if r["error"] != nil {
		return ""
	}
	return r["display_name"].(string)
}
