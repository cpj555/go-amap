package go_amap

import "github.com/cpj555/go-amap/client"

func NewClient(options ...client.OptionHandler) (*client.Client, error) {
	return client.New(options...)
}
