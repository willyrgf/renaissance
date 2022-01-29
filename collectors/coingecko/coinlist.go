package coingecko

import (
	"context"

	"github.com/carlmjohnson/requests"
)

type CoinLists []CoinList

type CoinList struct {
	ID        string   `json:"id"`
	Symbol    string   `json:"symbol"`
	Name      string   `json:"name"`
	Platforms struct{} `json:"platforms"`
}

const (
	baseURL = "https://api.coingecko.com/api/v3/"
)

func FetchCoinLists() (*CoinLists, error) {
	const endpoint = "/coins/list"

	var cls CoinLists
	err := requests.
		URL(baseURL+endpoint).
		Param("include_plataform", "true").
		ToJSON(&cls).
		Fetch(context.Background())

	return &cls, err
}
