package client

import "fmt"

type uri string

const (
	geoUri   uri = "/v3/geocode/geo"
	regeoUri uri = "/v3/geocode/regeo"

	walkingUri uri = "/v3/direction/walking"
	transitUri uri = "/v3/direction/transit/integrated"
	drivingUri uri = "/v3/direction/driving"
)

// getApi build the full api url
func (c *Client) getApi(u uri) string {
	return fmt.Sprintf("%s%s", c.conf.BaseApi, u)
}
