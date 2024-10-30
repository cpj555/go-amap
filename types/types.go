package types

// 高德地图返回的data键值不固定 目前方案是先序列化成map[string]interface{} 再取需要的键
type OpenAPIRsp struct {
	Status   string `json:"status"`
	Info     string `json:"info"`
	Infocode string `json:"infocode"`
	//Data     jsoniter.RawMessage `json:"-"`
}
