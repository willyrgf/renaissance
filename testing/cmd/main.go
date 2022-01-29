package main

import (
	"github.com/willyrgf/renaissance/collector"
	"github.com/willyrgf/renaissance/collector/coingecko"
	"github.com/willyrgf/renaissance/logger"
)

func init() {
	logger.New(logger.WithFlags)
}

func main() {
	collector := collector.New(coingecko.Default())
	coins, err := collector.Coins()
	logger.Infof("main(): collector.Coins(): coins=%+v err=%+v", coins[0:10], err)
}
