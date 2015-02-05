// Package mapquest is a geo-golang based MapRequest geocode/reverse geocode client
package mapquest

import (
	"encoding/json"
	"fmt"
	"github.com/codingsince1985/geo-golang"
	"strconv"
)

type Endpoint geo.Endpoint

type GeocodeResponse map[string]interface{}

func NewGeocoder() geo.Geocoder {
	return geo.Geocoder{
		Endpoint("http://open.mapquestapi.com/nominatim/v1/"),
		GeocodeResponse{},
	}
}

func (e Endpoint) GeocodeUrl(address string) string {
	return string(e) + "search.php?format=json&q=" + address
}

func (e Endpoint) ReverseGeocodeUrl(l geo.Location) string {
	return string(e) + fmt.Sprintf("reverse.php?format=json&lat=%f&lon=%f", l.Lat, l.Lng)
}

func (r GeocodeResponse) Location(data []byte) (l geo.Location) {
	res := []GeocodeResponse{}
	if json.Unmarshal(data, &res); len(res) > 0 && res[0]["lat"] != nil && res[0]["lon"] != nil {
		return geo.Location{parseFloat(res[0]["lat"]), parseFloat(res[0]["lon"])}
	}
	return
}

func (r GeocodeResponse) Address(data []byte) (address string) {
	if json.Unmarshal(data, &r); r["error"] == nil {
		address = r["display_name"].(string)
	}
	return
}

func parseFloat(value interface{}) float64 {
	f, _ := strconv.ParseFloat(value.(string), 64)
	return f
}
