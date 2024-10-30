package client

import (
	"context"
	"errors"
	"time"

	"github.com/cpj555/go-amap/types"
	"github.com/go-resty/resty/v2"
	"golang.org/x/time/rate"
)

type (

	// Base client basic interface
	Base interface {
		GetConfig() *Config
	}

	GeocodeAPI interface {
		// Geo 地理编码
		Geo(ctx context.Context, req *types.GeoReq) (*types.GeoResp, error)
		// ReGeo 地理解码
		ReGeo(ctx context.Context, req *types.ReGeoReq) (*types.ReGeoResp, error)
	}

	DirectionAPI interface {
		// Walking 步行路径规划
		Walking(ctx context.Context, req *types.WalkingReq) (*types.WalkResp, error)
		// Integrated 公交路径规划
		Integrated(ctx context.Context, req *types.IntegratedReq) (*types.IntegratedResp, error)
		// Driving 驾车路径规划
		Driving(ctx context.Context, req *types.DrivingReq) (*types.DrivingResp, error)
	}
)

type (
	Client struct {
		conf         *Config
		GeocodeAPI   GeocodeAPI
		DirectionAPI DirectionAPI
		r            *resty.Client
		limiter      *rate.Limiter
	}

	// Config amap client configuration
	Config struct {
		BaseApi          string // amap OpenAPI Server Domain
		RestyClient      *resty.Client
		ApiKey           string        // amap ApiKey
		Sign             bool          // amap sign
		IsDebug          bool          // debug mode
		Timeout          time.Duration // resty client request timeout
		ProxyUrl         string
		RequestPerSecond int
	}
)

// New create a new coinmarket instance
func New(options ...OptionHandler) (*Client, error) {
	config := getDefaultConfig()

	// handle custom options
	for _, optionHandler := range options {
		if optionHandler == nil {
			return nil, errors.New("invalid OptionHandler (nil detected)")
		}
		if err := optionHandler(config); err != nil {
			return nil, err
		}
	}

	instance := &Client{conf: config}
	if config.RestyClient != nil {
		instance.r = config.RestyClient
	}
	instance.setupResty()

	initAPI(instance)
	initLimiter(instance, config.RequestPerSecond)

	return instance, nil
}

func initAPI(client *Client) {
	client.GeocodeAPI = newGeocodeAPI(client)
	client.DirectionAPI = newDirectionAPI(client)
}

// 初始化限流器 配置api的等级使用
func initLimiter(client *Client, RequestPerSecond int) {
	if RequestPerSecond > 0 {
		client.limiter = rate.NewLimiter(rate.Every(time.Second), RequestPerSecond)
	}
}

// getDefaultConfig Get the default configuration
func getDefaultConfig() *Config {
	return &Config{
		BaseApi:          "https://restapi.amap.com",
		IsDebug:          false,
		Timeout:          time.Second * 5,
		RequestPerSecond: 1,
	}
}

// GetConfig get instance configuration
func (c *Client) GetConfig() *Config {
	return c.conf
}
