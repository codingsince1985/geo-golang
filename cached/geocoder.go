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
func (c cachedGeocoder) Geocode(address string) (*geo.Location, error) {
	// Check if we've cached this response
	if cachedLoc, found := c.Cache.Get(address); found {
		return cachedLoc.(*geo.Location), nil
	}

	if loc, err := c.Geocoder.Geocode(address); err != nil {
		return loc, err
	} else {
		c.Cache.Set(address, loc, 0)
		return loc, nil
	}
}

// ReverseGeocode returns address for location
func (c cachedGeocoder) ReverseGeocode(lat, lng float64) (*geo.Address, error) {
	// Check if we've cached this response
	locKey := fmt.Sprintf("geo.Location{%f,%f}", lat, lng)
	if cachedAddr, found := c.Cache.Get(locKey); found {
		return cachedAddr.(*geo.Address), nil
	}

	if addr, err := c.Geocoder.ReverseGeocode(lat, lng); err != nil {
		return nil, err
	} else {
		c.Cache.Set(locKey, addr, 0)
		return addr, nil
	}
}
