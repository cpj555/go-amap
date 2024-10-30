package main

import (
	"context"
	"fmt"
	goamap "github.com/cpj555/go-amap"
	"github.com/cpj555/go-amap/client"
	"github.com/cpj555/go-amap/types"
)

func main() {

	apiKey := ""

	opts := []client.OptionHandler{
		client.WithApiKey(apiKey),
		client.WithDebugMode(true),
	}

	c, err := goamap.NewClient(opts...)
	if err != nil {
		fmt.Printf("创建实例失败：%v\n", err)
		return
	}
	//getGeocode(c)
	//ReGeo(c)

	Walking(c)
}

func getGeocode(c *client.Client) {

	resp, err := c.GeocodeAPI.Geo(context.Background(), &types.GeoReq{Address: "浙江省杭州市"})

	fmt.Println(err)
	fmt.Println(resp)

}

func ReGeo(c *client.Client) {

	resp, err := c.GeocodeAPI.ReGeo(context.Background(), &types.ReGeoReq{Location: "116.481488,39.990464", Extensions: "all"})
	fmt.Println(err)
	fmt.Println(resp)
}

func Walking(c *client.Client) {

	resp, err := c.DirectionAPI.Walking(context.Background(), &types.WalkingReq{Origin: "116.481488,39.990464", Destination: "116.481488,39.990464"})
	fmt.Println(err)
	fmt.Println(resp)
}
