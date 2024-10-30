package client

import (
	"context"
	"github.com/cpj555/go-amap/tools"
	"github.com/cpj555/go-amap/types"
	"github.com/mitchellh/mapstructure"
)

type directionAPI struct {
	cli *Client
}

func newDirectionAPI(client *Client) DirectionAPI {
	return &directionAPI{cli: client}
}

func (d *directionAPI) Walking(ctx context.Context, req *types.WalkingReq) (*types.WalkResp, error) {

	values := tools.ToUrlValues(req)

	resp, err := d.cli.request(ctx).SetQueryParamsFromValues(values).Get(d.cli.getApi(walkingUri))
	if err != nil {
		return nil, err
	}

	var result types.WalkResp

	mapstructure.Decode(d.cli.unmarshalResult(resp), &result)

	return &result, nil
}

func (d *directionAPI) Integrated(ctx context.Context, req *types.IntegratedReq) (*types.IntegratedResp, error) {

	values := tools.ToUrlValues(req)

	resp, err := d.cli.request(ctx).SetQueryParamsFromValues(values).Get(d.cli.getApi(transitUri))
	if err != nil {
		return nil, err
	}

	var result types.IntegratedResp

	mapstructure.Decode(d.cli.unmarshalResult(resp), &result)

	return &result, nil
}

func (d *directionAPI) Driving(ctx context.Context, req *types.DrivingReq) (*types.DrivingResp, error) {

	values := tools.ToUrlValues(req)

	resp, err := d.cli.request(ctx).SetQueryParamsFromValues(values).Get(d.cli.getApi(drivingUri))

	if err != nil {
		return nil, err
	}

	var result types.DrivingResp

	mapstructure.Decode(d.cli.unmarshalResult(resp), &result)

	return &result, nil
}
