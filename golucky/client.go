package golucky

type Client struct {
	config Config
}

type Config struct {
	RestIp  string
	AuthKey string
}

func New(config Config) *Client {
	return &Client{config: config}
}
