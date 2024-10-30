package client

import (
	"context"

	"github.com/cpj555/go-amap/tools"
	"github.com/cpj555/go-amap/types"
	"github.com/mitchellh/mapstructure"
)

type geocodeAPI struct {
	cli *Client
}

func newGeocodeAPI(client *Client) GeocodeAPI {
	return &geocodeAPI{cli: client}
}

func (c *geocodeAPI) Geo(ctx context.Context, req *types.GeoReq) (*types.GeoResp, error) {
	values := tools.ToUrlValues(req)

	resp, err := c.cli.request(ctx).SetQueryParamsFromValues(values).Get(c.cli.getApi(geoUri))
	if err != nil {
		return nil, err
	}

	var result types.GeoResp

	mapstructure.Decode(c.cli.unmarshalResult(resp), &result)

	return &result, nil
}

func (c *geocodeAPI) ReGeo(ctx context.Context, req *types.ReGeoReq) (*types.ReGeoResp, error) {
	values := tools.ToUrlValues(req)

	resp, err := c.cli.request(ctx).SetQueryParamsFromValues(values).Get(c.cli.getApi(regeoUri))
	if err != nil {
		return nil, err
	}

	var result types.ReGeoResp

	mapstructure.Decode(c.cli.unmarshalResult(resp), &result)

	return &result, nil
}
