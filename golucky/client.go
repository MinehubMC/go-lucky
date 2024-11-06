package golucky

import "strings"

type Client struct {
	config Config
}

type Config struct {
	RestIp  string
	AuthKey string
}

func New(config Config) *Client {
	config.RestIp = strings.TrimSuffix(config.RestIp, "/") // Fix for trailing slash
	return &Client{config: config}
}
