package app

import (
	"aviary/config"
	"aviary/insights"
	"aviary/twitter"
)

type Gateway struct {
	Twitter *twitter.Gateway
	Logger  *insights.Logger
	Users   []string
}

// New returns generic client for interacting with Twitter.
func New() (Gateway, error) {
	prov, err := config.New()
	if err != nil {
		return Gateway{}, err
	}
	logger := insights.New(prov)
	twtGateway, err := twitter.New(prov)
	if err != nil {
		logger.Error("twitter.New: error", err)
		return Gateway{}, err
	}
	return Gateway{
		Twitter: twtGateway,
		Logger:  logger,
	}, nil
}
