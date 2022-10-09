package config

import (
	_ "embed"
	"net/http"
	"os"

	"github.com/dghubble/oauth1"
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
	// os is used intead of commiting credentials to config file.
	cfg := oauth1.NewConfig(os.Getenv(prov.Twitter.Key), os.Getenv(prov.Twitter.KeySecret))
	token := oauth1.NewToken(os.Getenv(prov.Twitter.Token), os.Getenv(prov.Twitter.TokenSecret))
	// http.Client should auto authorize reqs
	httpClient := cfg.Client(oauth1.NoContext, token)
	prov.Client = httpClient
	return prov, nil
}

// Provider defines the aviary config.
type Provider struct {
	Client  *http.Client
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
