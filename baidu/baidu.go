package baidu

import (
	"fmt"
	"strings"

	"github.com/codingsince1985/geo-golang"
)

var (
	languageList = []string{"el", "gu", "en", "vi", "ca", "it", "iw", "sv", "eu", "ar", "cs", "gl", "id", "es", "en-GB", "ru", "sr", "nl", "pt", "tr", "tl", "lv", "en-AU", "lt", "th", "ro", "fil", "ta", "fr", "bg", "hr", "bn", "de", "hu", "fa", "hi", "pt-BR", "fi", "da", "ja", "te", "pt-PT", "ml", "ko", "kn", "sk", "zh-CN", "pl", "uk", "sl", "mr", "local"}
)

type (
	baseURL string

	//Response payload
	//{'result': {'addressComponent': {'adcode': '310101',
	//                                'city': '上海市',
	//                                'city_level': 2,
	//                                'country': '中国',
	//                                'country_code': 0,
	//                                'country_code_iso': 'CHN',
	//                                'country_code_iso2': 'CN',
	//                                'direction': '东北',
	//                                'distance': '91',
	//                                'district': '黄浦区',
	//                                'province': '上海市',
	//                                'street': '中山南路',
	//                                'street_number': '187',
	//                                'town': '',
	//                                'town_code': ''},
	//           'business': '外滩,陆家嘴,董家渡',
	//           'cityCode': 289,
	//           'formatted_address': '上海市黄浦区中山南路187',
	//           'location': {'lat': 31.22932842411674, 'lng': 121.50989077799083},
	//           'poiRegions': [],
	//           'pois': [],
	//           'roads': [],
	//           'sematic_description': ''},
	//'status': 0}
	geocodeResponse struct {
		Result struct {
			AddressComponent struct {
				AdCode          string `json:"adcode"`
				City            string `json:"city"`
				CityLevel       int    `json:"city_level"`
				Country         string `json:"country"`
				CountryCode     int    `json:"country_code"`
				CountryCodeIso  string `json:"country_code_iso"`
				CountryCodeIso2 string `json:"country_code_iso2"`
				Direction       string `json:"direction"`
				Distance        string `json:"distance"`
				District        string `json:"district"`
				Province        string `json:"province"`
				Street          string `json:"street"`
				StreetNumber    string `json:"street_number"`
				Town            string `json:"town"`
				TownCode        string `json:"town_code"`
			} `json:"addressComponent"`
			Business         string `json:"business"`
			CityCode         int    `json:"cityCode"`
			FormattedAddress string `json:"formatted_address"`
			Location         struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
			PoiRegions         []interface{} `json:"poiRegions"`
			Pois               []interface{} `json:"pois"`
			Roads              []interface{} `json:"roads"`
			SematicDescription string        `json:"sematic_description"`
		} `json:"result"`
		Status int `json:"status"`
	}
)

const (
	statusOK = 0
)

var lang string

// Geocoder constructs Google geocoder
func Geocoder(apiKey string, language string, baseURLs ...string) geo.Geocoder {
	if lang != "" {
		for _, item := range languageList {
			if language == item {
				lang = fmt.Sprintf("&language=%s", language)
			}
		}
	}

	return geo.HTTPGeocoder{
		EndpointBuilder:       baseURL(getURL(apiKey, baseURLs...)),
		ResponseParserFactory: func() geo.ResponseParser { return &geocodeResponse{} },
	}
}

func getURL(apiKey string, baseURLs ...string) string {
	if len(baseURLs) > 0 {
		return baseURLs[0]
	}
	return fmt.Sprintf("https://api.map.baidu.com/*/v3/?ak=%s&", apiKey)
}

// GeocodeURL https://api.map.baidu.com/geocoding/v3/?ak=APPKEY&output=json&address=ADDRESS
func (b baseURL) GeocodeURL(address string) string {
	return strings.Replace(string(b), "*", "geocoding", 1) + fmt.Sprintf("output=json&address=%s", address)
}

// ReverseGeocodeURL https://api.map.baidu.com/reverse_geocoding/v3/?ak=APPKEY&output=json&&coordtype=wgs84ll&location=31.225696563611,121.49884033194
func (b baseURL) ReverseGeocodeURL(l geo.Location) string {
	return strings.Replace(string(b), "*", "reverse_geocoding", 1) + fmt.Sprintf("output=json&coordtype=wgs84ll&location=%f,%f", l.Lat, l.Lng) + lang
}

func (r *geocodeResponse) Location() (*geo.Location, error) {
	var location = &geo.Location{}
	if r.Status == 1 {
		return nil, nil
	} else if r.Status != statusOK {
		return nil, fmt.Errorf("geocoding error: %v", r.Status)
	}
	location.Lat = r.Result.Location.Lat
	location.Lng = r.Result.Location.Lng
	return location, nil

}

func (r *geocodeResponse) Address() (*geo.Address, error) {
	if r.Status == 1 {
		return nil, nil
	} else if r.Status != statusOK {
		return nil, fmt.Errorf("reverse geocoding error: %v", r.Status)
	}

	addr := parseBaiduResult(r)

	return addr, nil
}

func parseBaiduResult(r *geocodeResponse) *geo.Address {
	addr := &geo.Address{}
	res := r.Result
	addr.FormattedAddress = res.FormattedAddress
	addr.HouseNumber = res.AddressComponent.StreetNumber
	addr.Street = res.AddressComponent.Street
	addr.Suburb = res.AddressComponent.District
	addr.City = res.AddressComponent.City
	addr.State = res.AddressComponent.Province
	addr.Country = res.AddressComponent.Country
	addr.CountryCode = res.AddressComponent.CountryCodeIso

	if (*addr == geo.Address{}) {
		return nil
	}

	return addr
}
