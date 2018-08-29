package yandex_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/yandex"
	"github.com/stretchr/testify/assert"
)

var token = os.Getenv("YANDEX_API_KEY")

func TestGeocode(t *testing.T) {
	ts := testServer(response1)
	defer ts.Close()

	geocoder := yandex.Geocoder(token, ts.URL+"/")
	location, err := geocoder.Geocode("60 Collins St, Melbourne VIC 3000")
	assert.NoError(t, err)
	assert.Equal(t, geo.Location{Lat: -37.816939, Lng: 144.961515}, *location)
}

func TestReverseGeocode(t *testing.T) {
	ts := testServer(response2)
	defer ts.Close()

	geocoder := yandex.Geocoder(token, ts.URL+"/")
	address, err := geocoder.ReverseGeocode(37.816939, 144.961515)
	assert.NoError(t, err)
	assert.True(t, strings.HasPrefix(address.FormattedAddress, "Victoria, City of Melbourne, Collins Street"))
	assert.True(t, strings.HasPrefix(address.Street, "Collins Street"))
}

func TestReverseGeocodeWithNoResult(t *testing.T) {
	ts := testServer(response3)
	defer ts.Close()

	geocoder := yandex.Geocoder(token, ts.URL+"/")
	addr, err := geocoder.ReverseGeocode(-37.8137683, 164.9718448)
	assert.Nil(t, err)
	assert.Nil(t, addr)
}

func testServer(response string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte(response))
	}))
}

const (
	response1 = `{
   "response":{
      "GeoObjectCollection":{
         "metaDataProperty":{
            "GeocoderResponseMetaData":{
               "request":"60 Collins St, Melbourne VIC 3000",
               "found":"6",
               "results":"10"
            }
         },
         "featureMember":[
            {
               "GeoObject":{
                  "metaDataProperty":{
                     "GeocoderMetaData":{
                        "kind":"street",
                        "text":"Australia, Victoria, City of Melbourne, Collins Street",
                        "precision":"street",
                        "Address":{
                           "country_code":"AU",
                           "formatted":"Victoria, City of Melbourne, Collins Street",
                           "Components":[
                              {
                                 "kind":"country",
                                 "name":"Australia"
                              },
                              {
                                 "kind":"province",
                                 "name":"Victoria"
                              },
                              {
                                 "kind":"locality",
                                 "name":"City of Melbourne"
                              },
                              {
                                 "kind":"street",
                                 "name":"Collins Street"
                              }
                           ]
                        },
                        "AddressDetails":{
                           "Country":{
                              "AddressLine":"Victoria, City of Melbourne, Collins Street",
                              "CountryNameCode":"AU",
                              "CountryName":"Australia",
                              "AdministrativeArea":{
                                 "AdministrativeAreaName":"Victoria",
                                 "Locality":{
                                    "LocalityName":"City of Melbourne",
                                    "Thoroughfare":{
                                       "ThoroughfareName":"Collins Street"
                                    }
                                 }
                              }
                           }
                        }
                     }
                  },
                  "description":"City of Melbourne, Victoria, Australia",
                  "name":"Collins Street",
                  "boundedBy":{
                     "Envelope":{
                        "lowerCorner":"144.948795 -37.820788",
                        "upperCorner":"144.974199 -37.812984"
                     }
                  },
                  "Point":{
                     "pos":"144.961515 -37.816939"
                  }
               }
            },
            {
               "GeoObject":{
                  "metaDataProperty":{
                     "GeocoderMetaData":{
                        "kind":"street",
                        "text":"Australia, Victoria, City of Melbourne, Collins Street",
                        "precision":"street",
                        "Address":{
                           "country_code":"AU",
                           "formatted":"Victoria, City of Melbourne, Collins Street",
                           "Components":[
                              {
                                 "kind":"country",
                                 "name":"Australia"
                              },
                              {
                                 "kind":"province",
                                 "name":"Victoria"
                              },
                              {
                                 "kind":"locality",
                                 "name":"City of Melbourne"
                              },
                              {
                                 "kind":"street",
                                 "name":"Collins Street"
                              }
                           ]
                        },
                        "AddressDetails":{
                           "Country":{
                              "AddressLine":"Victoria, City of Melbourne, Collins Street",
                              "CountryNameCode":"AU",
                              "CountryName":"Australia",
                              "AdministrativeArea":{
                                 "AdministrativeAreaName":"Victoria",
                                 "Locality":{
                                    "LocalityName":"City of Melbourne",
                                    "Thoroughfare":{
                                       "ThoroughfareName":"Collins Street"
                                    }
                                 }
                              }
                           }
                        }
                     }
                  },
                  "description":"City of Melbourne, Victoria, Australia",
                  "name":"Collins Street",
                  "boundedBy":{
                     "Envelope":{
                        "lowerCorner":"144.886595 -37.858987",
                        "upperCorner":"144.889093 -37.858787"
                     }
                  },
                  "Point":{
                     "pos":"144.887844 -37.858887"
                  }
               }
            },
            {
               "GeoObject":{
                  "metaDataProperty":{
                     "GeocoderMetaData":{
                        "kind":"street",
                        "text":"Australia, Victoria, City of Melbourne, Collins Street",
                        "precision":"street",
                        "Address":{
                           "country_code":"AU",
                           "formatted":"Victoria, City of Melbourne, Collins Street",
                           "Components":[
                              {
                                 "kind":"country",
                                 "name":"Australia"
                              },
                              {
                                 "kind":"province",
                                 "name":"Victoria"
                              },
                              {
                                 "kind":"locality",
                                 "name":"City of Melbourne"
                              },
                              {
                                 "kind":"street",
                                 "name":"Collins Street"
                              }
                           ]
                        },
                        "AddressDetails":{
                           "Country":{
                              "AddressLine":"Victoria, City of Melbourne, Collins Street",
                              "CountryNameCode":"AU",
                              "CountryName":"Australia",
                              "AdministrativeArea":{
                                 "AdministrativeAreaName":"Victoria",
                                 "Locality":{
                                    "LocalityName":"City of Melbourne",
                                    "Thoroughfare":{
                                       "ThoroughfareName":"Collins Street"
                                    }
                                 }
                              }
                           }
                        }
                     }
                  },
                  "description":"City of Melbourne, Victoria, Australia",
                  "name":"Collins Street",
                  "boundedBy":{
                     "Envelope":{
                        "lowerCorner":"144.892991 -37.806784",
                        "upperCorner":"144.893198 -37.805686"
                     }
                  },
                  "Point":{
                     "pos":"144.893099 -37.806235"
                  }
               }
            },
            {
               "GeoObject":{
                  "metaDataProperty":{
                     "GeocoderMetaData":{
                        "kind":"street",
                        "text":"Australia, Victoria, City of Melbourne, Collins Street",
                        "precision":"street",
                        "Address":{
                           "country_code":"AU",
                           "formatted":"Victoria, City of Melbourne, Collins Street",
                           "Components":[
                              {
                                 "kind":"country",
                                 "name":"Australia"
                              },
                              {
                                 "kind":"province",
                                 "name":"Victoria"
                              },
                              {
                                 "kind":"locality",
                                 "name":"City of Melbourne"
                              },
                              {
                                 "kind":"street",
                                 "name":"Collins Street"
                              }
                           ]
                        },
                        "AddressDetails":{
                           "Country":{
                              "AddressLine":"Victoria, City of Melbourne, Collins Street",
                              "CountryNameCode":"AU",
                              "CountryName":"Australia",
                              "AdministrativeArea":{
                                 "AdministrativeAreaName":"Victoria",
                                 "Locality":{
                                    "LocalityName":"City of Melbourne",
                                    "Thoroughfare":{
                                       "ThoroughfareName":"Collins Street"
                                    }
                                 }
                              }
                           }
                        }
                     }
                  },
                  "description":"City of Melbourne, Victoria, Australia",
                  "name":"Collins Street",
                  "boundedBy":{
                     "Envelope":{
                        "lowerCorner":"145.069295 -37.98182",
                        "upperCorner":"145.070597 -37.981486"
                     }
                  },
                  "Point":{
                     "pos":"145.069941 -37.981806"
                  }
               }
            },
            {
               "GeoObject":{
                  "metaDataProperty":{
                     "GeocoderMetaData":{
                        "kind":"street",
                        "text":"Australia, Victoria, City of Melbourne, Little Collins Street",
                        "precision":"street",
                        "Address":{
                           "country_code":"AU",
                           "formatted":"Victoria, City of Melbourne, Little Collins Street",
                           "Components":[
                              {
                                 "kind":"country",
                                 "name":"Australia"
                              },
                              {
                                 "kind":"province",
                                 "name":"Victoria"
                              },
                              {
                                 "kind":"locality",
                                 "name":"City of Melbourne"
                              },
                              {
                                 "kind":"street",
                                 "name":"Little Collins Street"
                              }
                           ]
                        },
                        "AddressDetails":{
                           "Country":{
                              "AddressLine":"Victoria, City of Melbourne, Little Collins Street",
                              "CountryNameCode":"AU",
                              "CountryName":"Australia",
                              "AdministrativeArea":{
                                 "AdministrativeAreaName":"Victoria",
                                 "Locality":{
                                    "LocalityName":"City of Melbourne",
                                    "Thoroughfare":{
                                       "ThoroughfareName":"Little Collins Street"
                                    }
                                 }
                              }
                           }
                        }
                     }
                  },
                  "description":"City of Melbourne, Victoria, Australia",
                  "name":"Little Collins Street",
                  "boundedBy":{
                     "Envelope":{
                        "lowerCorner":"144.953699 -37.818087",
                        "upperCorner":"144.973597 -37.812286"
                     }
                  },
                  "Point":{
                     "pos":"144.963671 -37.815293"
                  }
               }
            },
            {
               "GeoObject":{
                  "metaDataProperty":{
                     "GeocoderMetaData":{
                        "kind":"street",
                        "text":"Australia, Victoria, Collins Street",
                        "precision":"street",
                        "Address":{
                           "country_code":"AU",
                           "formatted":"Victoria, Collins Street",
                           "Components":[
                              {
                                 "kind":"country",
                                 "name":"Australia"
                              },
                              {
                                 "kind":"province",
                                 "name":"Victoria"
                              },
                              {
                                 "kind":"street",
                                 "name":"Collins Street"
                              }
                           ]
                        },
                        "AddressDetails":{
                           "Country":{
                              "AddressLine":"Victoria, Collins Street",
                              "CountryNameCode":"AU",
                              "CountryName":"Australia",
                              "AdministrativeArea":{
                                 "AdministrativeAreaName":"Victoria",
                                 "Locality":{
                                    "Thoroughfare":{
                                       "ThoroughfareName":"Collins Street"
                                    }
                                 }
                              }
                           }
                        }
                     }
                  },
                  "description":"Victoria, Australia",
                  "name":"Collins Street",
                  "boundedBy":{
                     "Envelope":{
                        "lowerCorner":"145.147493 -37.673089",
                        "upperCorner":"145.149299 -37.666984"
                     }
                  },
                  "Point":{
                     "pos":"145.148418 -37.670033"
                  }
               }
            }
         ]
      }
   }
}`
	response2 = `{
   "response":{
      "GeoObjectCollection":{
         "metaDataProperty":{
            "GeocoderResponseMetaData":{
               "request":"-37.816939,144.961515",
               "found":"4",
               "results":"1",
               "Point":{
                  "pos":"144.961515 -37.816939"
               }
            }
         },
         "featureMember":[
            {
               "GeoObject":{
                  "metaDataProperty":{
                     "GeocoderMetaData":{
                        "kind":"street",
                        "text":"Australia, Victoria, City of Melbourne, Collins Street",
                        "precision":"street",
                        "Address":{
                           "country_code":"AU",
                           "formatted":"Victoria, City of Melbourne, Collins Street",
                           "Components":[
                              {
                                 "kind":"country",
                                 "name":"Australia"
                              },
                              {
                                 "kind":"province",
                                 "name":"Victoria"
                              },
                              {
                                 "kind":"locality",
                                 "name":"City of Melbourne"
                              },
                              {
                                 "kind":"street",
                                 "name":"Collins Street"
                              }
                           ]
                        },
                        "AddressDetails":{
                           "Country":{
                              "AddressLine":"Victoria, City of Melbourne, Collins Street",
                              "CountryNameCode":"AU",
                              "CountryName":"Australia",
                              "AdministrativeArea":{
                                 "AdministrativeAreaName":"Victoria",
                                 "Locality":{
                                    "LocalityName":"City of Melbourne",
                                    "Thoroughfare":{
                                       "ThoroughfareName":"Collins Street"
                                    }
                                 }
                              }
                           }
                        }
                     }
                  },
                  "description":"City of Melbourne, Victoria, Australia",
                  "name":"Collins Street",
                  "boundedBy":{
                     "Envelope":{
                        "lowerCorner":"144.948795 -37.820788",
                        "upperCorner":"144.974199 -37.812984"
                     }
                  },
                  "Point":{
                     "pos":"144.961515 -37.816939"
                  }
               }
            }
         ]
      }
   }
}`
	response3 = `{
   "response":{
      "GeoObjectCollection":{
         "metaDataProperty":{
            "GeocoderResponseMetaData":{
               "request":"179.9718448,-37.8137683",
               "found":"0",
               "results":"10",
               "boundedBy":{
                  "Envelope":{
                     "lowerCorner":"179.969346 -37.816270",
                     "upperCorner":"179.974341 -37.811267"
                  }
               },
               "Point":{
                  "pos":"179.971845 -37.813768"
               },
               "kind":"house"
            }
         },
         "featureMember":[]
      }
   }
}`
)
