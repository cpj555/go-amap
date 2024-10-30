package types

// GeoReq https://lbs.amap.com/api/webservice/guide/api/georegeo
type GeoReq struct {
	Address string `schema:"address"`
	City    string `schema:"city,omitempty"`
}

type GeoResp struct {
	Count    string `json:"count"`
	Geocodes []struct {
		FormattedAddress string        `json:"formatted_address"`
		Country          string        `json:"country"`
		Province         string        `json:"province"`
		Citycode         string        `json:"citycode"`
		City             string        `json:"city"`
		District         string        `json:"district"`
		Township         []interface{} `json:"township"`
		Neighborhood     struct {
			Name []interface{} `json:"name"`
			Type []interface{} `json:"type"`
		} `json:"neighborhood"`
		Building struct {
			Name []interface{} `json:"name"`
			Type []interface{} `json:"type"`
		} `json:"building"`
		Adcode   string `json:"adcode"`
		Street   string `json:"street"`
		Number   string `json:"number"`
		Location string `json:"location"`
		Level    string `json:"level"`
	} `json:"geocodes"`
}

type ReGeoReq struct {
	Location   string `schema:"location"`
	Poitype    string `schema:"poitype,omitempty"`    //传入内容规则：经度在前，纬度在后，经纬度间以“,”分割，经纬度小数点后不要超过 6 位。
	Radius     int    `schema:"radius,omitempty"`     // 以下内容需要 extensions 参数为 all 时才生效。 逆地理编码在进行坐标解析之后不仅可以返回地址描述，也可以返回经纬度附近符合限定要求的 POI 内容（在 extensions 字段值为 all 时才会返回 POI 内容）。设置 POI 类型参数相当于为上述操作限定要求。参数仅支持传入 POI TYPECODE，可以传入多个 POI TYPECODE，相互之间用“|”分隔。获取 POI TYPECODE 可以参考 POI 分类码表
	Extensions string `schema:"extensions,omitempty"` // extensions 参数默认取值是 base，也就是返回基本地址信息； extensions 参数取值为 all 时会返回基本地址信息、附近 POI 内容、道路信息以及道路交叉口信息。
	Roadlevel  int    `schema:"roadlevel,omitempty"`  //以下内容需要 extensions 参数为 all 时才生效。 可选值：0，1  当 roadlevel=0时，显示所有道路 ； 当 roadlevel=1时，过滤非主干道路，仅输出主干道路数据
	Homeorcorp int    `schema:"homeorcorp,omitempty"` //以下内容需要 extensions 参数为 all 时才生效。 homeorcorp 参数的设置可以影响召回 POI 内容的排序策略，目前提供三个可选参数： 0：不对召回的排序策略进行干扰。 1：综合大数据分析将居家相关的 POI 内容优先返回，即优化返回结果中 pois 字段的poi 顺序。 2：综合大数据分析将公司相关的 POI 内容优先返回，即优化返回结果中 pois 字段的poi 顺序。
}

type ReGeoResp struct {
	Regeocode struct {
		FormattedAddress string `json:"formatted_address"`
		AddressComponent struct {
			Country      string        `json:"country"`
			Province     string        `json:"province"`
			City         []interface{} `json:"city"`
			Citycode     string        `json:"citycode"`
			District     string        `json:"district"`
			Adcode       string        `json:"adcode"`
			Township     string        `json:"township"`
			Towncode     string        `json:"towncode"`
			Neighborhood struct {
				Name string `json:"name"`
				Type string `json:"type"`
			} `json:"neighborhood"`
			Building struct {
				Name string `json:"name"`
				Type string `json:"type"`
			} `json:"building"`
			StreetNumber struct {
				Street    string `json:"street"`
				Number    string `json:"number"`
				Location  string `json:"location"`
				Direction string `json:"direction"`
				Distance  string `json:"distance"`
			} `json:"streetNumber"`
			BusinessAreas []struct {
				Location string `json:"location"`
				Name     string `json:"name"`
				Id       string `json:"id"`
			} `json:"businessAreas"`
		} `json:"addressComponent"`
		Pois []struct {
			Id           string      `json:"id"`
			Name         string      `json:"name"`
			Type         string      `json:"type"`
			Tel          interface{} `json:"tel"`
			Direction    string      `json:"direction"`
			Distance     string      `json:"distance"`
			Location     string      `json:"location"`
			Address      string      `json:"address"`
			Poiweight    string      `json:"poiweight"`
			Businessarea string      `json:"businessarea"`
		} `json:"pois"`
		Roads []struct {
			Id        string `json:"id"`
			Name      string `json:"name"`
			Direction string `json:"direction"`
			Distance  string `json:"distance"`
			Location  string `json:"location"`
		} `json:"roads"`
		Roadinters []struct {
			Direction  string `json:"direction"`
			Distance   string `json:"distance"`
			Location   string `json:"location"`
			FirstId    string `json:"first_id"`
			FirstName  string `json:"first_name"`
			SecondId   string `json:"second_id"`
			SecondName string `json:"second_name"`
		} `json:"roadinters"`
		Aois []struct {
			Id       string `json:"id"`
			Name     string `json:"name"`
			Adcode   string `json:"adcode"`
			Location string `json:"location"`
			Area     string `json:"area"`
			Distance string `json:"distance"`
			Type     string `json:"type"`
		} `json:"aois"`
	} `json:"regeocode"`
}
