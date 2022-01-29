package collector

import "github.com/willyrgf/renaissance/collector/coingecko"

type Collecter interface {
	//FIXME: its just a example
	FetchCoinLists() (*coingecko.CoinLists, error)
}

type Collector[T Collecter] struct {
	provider T
}

func New[T Collecter](p T) *Collector[Collecter] {
	return &Collector[Collecter]{provider: p}
}

func (c *Collector[Provider]) Coins() ([]string, error) {
	coins := []string{}

	//FIXME: its just a example
	cl, err := c.provider.FetchCoinLists()
	if err != nil {
		return coins, err
	}

	for _, c := range *cl {
		coins = append(coins, c.Symbol)
	}

	return coins, nil
}
