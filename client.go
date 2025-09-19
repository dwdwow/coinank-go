package coinank

type Client struct {
	apiKey string
}

type ClientCfg struct {
	ApiKey string
}

func NewClient(cfg ClientCfg) *Client {
	return &Client{
		apiKey: cfg.ApiKey,
	}
}
