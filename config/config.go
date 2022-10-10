package config

import (
	_ "embed"
	"os"

	"gopkg.in/yaml.v2"
)

//go:embed config.yaml
var c []byte

// New manages loading configs.
func New() (Provider, error) {
	var prov Provider
	err := yaml.Unmarshal(c, &prov)
	if err != nil {
		return prov, err
	}
	// Credentials configured in Function App Configuration
	// is used instead of KeyVault.
	prov.Twitter.Key = os.Getenv(prov.Twitter.Key)
	prov.Twitter.KeySecret = os.Getenv(prov.Twitter.KeySecret)
	prov.Twitter.Token = os.Getenv(prov.Twitter.Token)
	prov.Twitter.TokenSecret = os.Getenv(prov.Twitter.TokenSecret)
	return prov, nil
}

// Provider defines the aviary config.
type Provider struct {
	Users   []string `yaml:"users"`
	Twitter Twitter  `yaml:"twitter"`
	Azure   Azure    `yaml:"azure"`
}

// Twitter defines required twitter credentials.
type Twitter struct {
	Key         string `yaml:"key"`          // consumer key
	KeySecret   string `yaml:"key_secret"`   // consumer key secret
	Token       string `yaml:"token"`        // access token
	TokenSecret string `yaml:"token_secret"` // access token secret
}

// Azure defines required azure credentials.
type Azure struct {
	InsightsKey string `yaml:"insights_key"`
}
