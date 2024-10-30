<p align="center">
  <a href="https://coinmarketcap.com/api/documentation/v1/">
    <img src="examples/amap.svg" width="200" height="200" alt="dodo-open">
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
list, err := client.CryptocurrencyV1.GetMap(context.Background(),&types.GetCryptocurrencyV1MapReq{Start: 1, Limit: 10})
```

### API documentation
https://lbs.amap.com/api/webservice/guide/api/georegeo

- [examples](examples/demo-setup)


* **/v3/geocode/geo** (地理编码/逆地理编码提供结构化地址与经纬度之间的相互转化的能力。)
  ```go
  client.CryptocurrencyV1.GetMap(context.Background(),&types.GetCryptocurrencyMapReq{Start: 1, Limit: 10})
  ```