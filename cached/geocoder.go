package cached

import (
	"fmt"

	"github.com/codingsince1985/geo-golang"
	"github.com/patrickmn/go-cache"
)

type cachedGeocoder struct {
	Geocoder geo.Geocoder
	Cache    *cache.Cache
}

// Geocoder creates a chain of Geocoders to lookup address and fallback on
func Geocoder(geocoder geo.Geocoder, cache *cache.Cache) geo.Geocoder {
	return cachedGeocoder{Geocoder: geocoder, Cache: cache}
}

// Geocode returns location for address
func (c cachedGeocoder) Geocode(address string) (geo.Location, error) {
	// Check if we've cached this response
	cachedLoc, found := c.Cache.Get(address)
	if found {
		return cachedLoc.(geo.Location), nil
	}

	loc, err := c.Geocoder.Geocode(address)
	if err != nil {
		return loc, err
	}
	c.Cache.Set(address, loc, 0)
	return loc, nil
}

// ReverseGeocode returns address for location
func (c cachedGeocoder) ReverseGeocode(lat, lng float64) (string, error) {
	// Check if we've cached this response
	locKey := fmt.Sprintf("geo.Location{%f,%f}", lat, lng)
	cachedAddr, found := c.Cache.Get(locKey)
	if found {
		return cachedAddr.(string), nil
	}

	addr, err := c.Geocoder.ReverseGeocode(lat, lng)
	if err != nil {
		return "", err
	}
	c.Cache.Set(locKey, addr, 0)
	return addr, nil
}
