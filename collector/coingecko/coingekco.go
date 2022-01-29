package coingecko

const (
	baseURL = "https://api.coingecko.com/api/v3/"
)

type Provider struct {
}

func new() *Provider {
	return &Provider{}
}

func Default() *Provider {
	return new()
}
