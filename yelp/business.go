package yelp

import (
	"fmt"
	"net/url"

	"github.com/k4orta/lunchclub-api/models"
	"github.com/k4orta/lunchclub-api/storage/types"
)

var businessEndpoint = "http://api.yelp.com/v2/business/"

// FetchBusiness makes a request to the YelpAPI for
func FetchBusiness(id string) (*models.Location, error) {
	var ret struct {
		ID       string
		Name     string
		Location struct {
			Address    []string
			Coordinate struct {
				Latitude  float64
				Longitude float64
			}
		}
	}
	fmt.Println(businessEndpoint + id)
	err := oauthClient.get(businessEndpoint+id, url.Values{}, &ret)
	if err != nil {
		return nil, err
	}
	m := models.Location{
		Slug:    ret.ID,
		Name:    ret.Name,
		Address: ret.Location.Address[0],
		LatLng: types.FloatList{
			ret.Location.Coordinate.Latitude,
			ret.Location.Coordinate.Longitude,
		},
	}
	return &m, nil
}
