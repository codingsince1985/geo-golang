package geo

type Location struct {
	Lat, Lng float64
}

type Endpoint struct {
	BaseUrl string
}

type GeocodeEndpointBuilder interface {
	GeocodeUrl(string) string
	ReverseGeocodeUrl(Location) string
}

type GeocodeResponseParser interface {
	Location([]byte) Location
	Address([]byte) string
}

type Geocoder struct {
	GeocodeEndpointBuilder
	GeocodeResponseParser
}

func (g Geocoder) Geocode(address string) Location {
	return g.Location(ResponseData(g.GeocodeUrl(address)))
}

func (g Geocoder) ReverseGeocode(l Location) string {
	return g.Address(ResponseData(g.ReverseGeocodeUrl(l)))
}
