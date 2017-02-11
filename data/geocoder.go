package data

import (
	"github.com/codingsince1985/geo-golang"
)

// AddressToLocation maps address string to location (lat, long)
type AddressToLocation map[geo.Address]geo.Location

// LocationToAddress maps location(lat,lng) to address
type LocationToAddress map[geo.Location]geo.Address

// dataGeocoder represents geo data in memory
type dataGeocoder struct {
	AddressToLocation
	LocationToAddress
}

// Geocoder constructs data geocoder
func Geocoder(addressToLocation AddressToLocation, LocationToAddress LocationToAddress) geo.Geocoder {
	return dataGeocoder{
		AddressToLocation: addressToLocation,
		LocationToAddress: LocationToAddress,
	}
}

// Geocode returns location for address
func (d dataGeocoder) Geocode(address string) (*geo.Location, error) {
	addr := geo.Address{
		FormattedAddress: address,
	}
	if l, ok := d.AddressToLocation[addr]; ok {
		return &l, nil
	}

	return nil, nil
}

// ReverseGeocode returns address for location
func (d dataGeocoder) ReverseGeocode(lat, lng float64) (*geo.Address, error) {
	if address, ok := d.LocationToAddress[geo.Location{Lat: lat, Lng: lng}]; ok {
		return &address, nil
	}
	return nil, nil
}
