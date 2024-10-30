<p align="center">
  <a href="https://coinmarketcap.com/api/documentation/v1/">
    <img src="examples/amap.svg" width="200" height="200" alt="go-amap-open">
  </a>
</p>

<div align="center">


# go-amap
  <a href="https://github.com/cpj555/go-amap/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/cpj555/go-amap" alt="license">
  </a>

</div>



## 特性

- 高德地图 OpenAPI
- 目前只封装了少量接口


## 起步

```shell
go get -u github.com/cpj555/go-amap
```



```go
apiKey := "your-api-key"

// 下面的第三个参数，设定了 resty 的请求超时为 3 秒：
//
//     client.WithTimeout(time.Second*3)
//
client, err := NewClient(
	client.WithApiKey(apiKey), 
	client.WithTimeout(time.Second*3),
	)

// 获取map列表，可以使用下面的方法
list, err := client.GeocodeAPI.Geo(context.Background(), &types.GeoReq{Address: "浙江省杭州市"})
```

### API documentation
https://lbs.amap.com/api/webservice/guide/api/georegeo

- [examples](examples/demo-setup)


* **/v3/geocode/geo** (地理编码提供结构化地址与经纬度之间的相互转化的能力。)
* **/v3/geocode/regeo** (逆地理编码提供结构化地址与经纬度之间的相互转化的能力。)
  ```go
  client.GeocodeAPI.Geo(context.Background(), &types.GeoReq{Address: "浙江省杭州市"})
  client.GeocodeAPI.ReGeo(context.Background(), &types.ReGeoReq{Location: "116.481488,39.990464", Extensions: "all"})
  ```