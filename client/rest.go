package client

import (
	"context"
	"github.com/go-resty/resty/v2"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/cpj555/go-amap/errs"
	"github.com/cpj555/go-amap/network"
	"github.com/cpj555/go-amap/tools"
)

// setupResty setup resty client
func (c *Client) setupResty() {
	if c.r == nil {
		c.r = resty.New()
	}

	c.r.
		SetTransport(createTransport(nil, 500)).
		SetDebug(c.conf.IsDebug).
		SetTimeout(c.conf.Timeout).
		OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
			request.SetQueryParam("key", c.conf.ApiKey)
			//todo sign
			return nil
		}).
		OnAfterResponse(
			func(r *resty.Client, resp *resty.Response) error {
				if !network.IsSuccessResponse(resp.StatusCode()) {
					return errs.New(resp.StatusCode(), string(resp.Body()))
				}
				rsp := c.unmarshalResult(resp)

				if rsp["status"].(string) != network.OpenAPIStatusOK {
					code, _ := strconv.Atoi(rsp["infocode"].(string))
					return errs.New(code, rsp["info"].(string))
				}
				return nil
			},
		)
	if c.conf.ProxyUrl != "" {
		c.r.SetProxy(c.conf.ProxyUrl)
	}

	c.r.JSONMarshal = tools.JSON.Marshal
	c.r.JSONUnmarshal = tools.JSON.Unmarshal
}

// request you should create a request object before doing each HTTP request
func (c *Client) request(ctx context.Context) *resty.Request {
	if c.limiter != nil {
		c.limiter.Wait(ctx)
	}
	return c.r.R().
		SetContext(ctx).
		SetResult(map[string]interface{}{})
}

// unmarshalResult get model.OpenAPIRsp result from the response
func (c *Client) unmarshalResult(resp *resty.Response) map[string]interface{} {
	return *resp.Result().(*map[string]interface{})
}

// createTransport customize transport
func createTransport(addr net.Addr, idleConnections int) *http.Transport {
	dialer := &net.Dialer{
		Timeout:   60 * time.Second,
		KeepAlive: 60 * time.Second,
	}
	if addr != nil {
		dialer.LocalAddr = addr
	}
	return &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		ForceAttemptHTTP2:     false,
		MaxIdleConns:          idleConnections,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConnsPerHost:   idleConnections,
		MaxConnsPerHost:       idleConnections,
	}
}
