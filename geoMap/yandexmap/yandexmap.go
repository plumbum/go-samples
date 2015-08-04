package yandexmap

import (
	"net/url"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

// https://tech.yandex.ru/maps/doc/jsapi/2.1/dg/concepts/load-docpage/

type Point struct {
	lat float64
	lng float64
}

type GeoData struct {
	pos Point
	lowerCorner Point
	upperCorner Point
	country string
	countryCode string
	city string
	address string
	precision string
}

func Geocode(address string, api string) (*[]GeoData, error) {
	var err error

	uri := getUri(address, api)

	client := &http.Client{}
	resp, err := client.Get(uri)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, errors.New("ZERO lenght response")
	}

	var ya yandexGeocodeResponse
	err = json.Unmarshal(data, &ya)
	if err != nil {
		return nil, err
	}

	cnt, err := strconv.Atoi(ya.Response.GeoObjectCollection.MetaDataProperty.GeocoderResponseMetaData.Found)
	if err != nil {
		return nil, err
	}

	geoData := make([]GeoData, cnt)

	for idx, obj := range ya.Response.GeoObjectCollection.FeatureMember {
		var gd GeoData

		_, err = fmt.Sscan(obj.GeoObject.Point.Pos, &gd.pos.lng, &gd.pos.lat)
		if err != nil {
			return nil, err
		}
		_, err = fmt.Sscan(obj.GeoObject.BoundedBy.Envelope.LowerCorner, &gd.lowerCorner.lng, &gd.lowerCorner.lat)
		if err != nil {
			return nil, err
		}
		_, err = fmt.Sscan(obj.GeoObject.BoundedBy.Envelope.UpperCorner, &gd.upperCorner.lng, &gd.upperCorner.lat)
		if err != nil {
			return nil, err
		}
		gd.address = obj.GeoObject.MetaDataProperty.GeocoderMetaData.AddressDetails.Country.AddressLine
		gd.city = obj.GeoObject.MetaDataProperty.GeocoderMetaData.AddressDetails.Country.AdministrativeArea.SubAdministrativeArea.Locality.LocalityName
		gd.country = obj.GeoObject.MetaDataProperty.GeocoderMetaData.AddressDetails.Country.CountryName
		gd.countryCode = obj.GeoObject.MetaDataProperty.GeocoderMetaData.AddressDetails.Country.CountryNameCode
		gd.precision = obj.GeoObject.MetaDataProperty.GeocoderMetaData.Precision
		geoData[idx] = gd
	}

	return &geoData, nil
}

type yandexGeocodeResponse struct {

	Response struct {
		GeoObjectCollection struct {
			MetaDataProperty struct {
				GeocoderResponseMetaData struct {
					Request string `json:"request"`
					Found   string `json:"found"`
					Results string `json:"results"`
				}
			} `json:"metaDataProperty"`
			FeatureMember []struct {
				GeoObject struct {
					MetaDataProperty struct {
						GeocoderMetaData struct {
							Kind           string `json:"kind"`
							Text           string `json:"text"`
							Precision      string `json:"precision"`
							AddressDetails struct {
								Country struct {
									AddressLine        string
									CountryNameCode    string
									CountryName        string
									AdministrativeArea struct {
										AdministrativeAreaName string
										SubAdministrativeArea struct {
											SubAdministrativeAreaName string
											Locality struct {
												LocalityName string
												Thoroughfare struct {
													ThoroughfareName string
													Premise struct {
														PremiseName string
														PremiseNumber string
													}
												}
												Premise struct {
													PremiseName string
													PremiseNumber string
												}
											}
										}
									}
								}
							}
						}
					} `json:"metaDataProperty"`
					Description      string `json:"description"`
					Name             string `json:"name"`
					BoundedBy struct {
						Envelope struct {
							LowerCorner string `json:"lowerCorner"`
							UpperCorner string `json:"upperCorner"`
						}
				    } `json:"boundedBy"`
					Point struct {
						Pos string `json:"pos"`
					}
				}
			} `json:"featureMember"`
		}
	} `json:"response"`
}

func getUri(addr string, api string) string {
	// raw := "http://geocode-maps.yandex.ru/1.x/?format=json&geocode=url_addr&key="+api
	q := &url.Values{}
	q.Set("format", "json")
	q.Set("geocode", addr)
	if api != "" {
		q.Set("key", api)
	}

	u := url.URL{
		Scheme: "http",
		Host: "geocode-maps.yandex.ru",
		Path: "/1.x/",
		RawQuery: q.Encode(),
	}

	return u.String()
}



/* Deprecated

type TYandexMap struct {
	Response TResponse `json:"response"`
}

type TResponse struct {
	GeoObjectCollection TGeoObjectCollection
}

type TGeoObjectCollection struct {
	MetaDataProperty TmetaDataProperty1 `json:"metaDataProperty"`
	FeatureMember []TfeatureMember `json:"featureMember"`
}

type TmetaDataProperty1 struct {
	GeocoderResponseMetaData TGeocoderResponseMetaData
}

type TGeocoderResponseMetaData struct {
	Request string `json:"request"`
	Found string `json:"found"`
	Results string `json:"results"`
}

type TfeatureMember struct {
	GeoObject TGeoObject
}

type TGeoObject struct {
	MetaDataProperty TmetaDataProperty2 `json:"metaDataProperty"`
	Description string `json:"description"`
	Name string `json:"name"`
	BoundedBy TboundedBy `json:"boundedBy"`
	Point TPoint
}

type TPoint struct {
	Pos string `json:"pos"`
}

type TboundedBy struct {
	Envelope TEnvelope
}

type TEnvelope struct {
	LowerCorner string `json:"lowerCorner"`
	UpperCorner string `json:"upperCorner"`
}

type TmetaDataProperty2 struct {
	GeocoderMetaData TGeocoderMetaData
}

type TGeocoderMetaData struct {
	Kind string `json:"kind"`
	Text string `json:"text"`
	Precision string `json:"precision"`
	AddressDetails TAddressDetails
}

type TAddressDetails struct {
	Country TCountry
}

type TCountry struct {
	AddressLine string
	CountryNameCode string
	CountryName string
	AdministrativeArea TAdministrativeArea
}

type TAdministrativeArea struct {
	AdministrativeAreaName string
	SubAdministrativeArea TSubAdministrativeArea
}

type TSubAdministrativeArea struct {
	SubAdministrativeAreaName string
	Locality TLocality
}

type TLocality struct {
	LocalityName string
	Thoroughfare TThoroughfare
}

type TThoroughfare struct {
	ThoroughfareName string
	Premise TPremise
}

type TPremise struct {
	PremiseNumber string
}

 */

