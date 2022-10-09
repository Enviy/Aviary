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

func New() (Provider, error) {
	var provider Provider
	err := yaml.Unmarshal(c, &provider)
	if err != nil {
		return provider, err
	}
	// os is used intead of commiting credentials to config file.
	cfg := oauth1.NewConfig(os.Getenv("TWT_CONSUMER_KEY"), os.Getenv("TWT_CONSUMER_SECRET"))
	token := oauth1.NewToken(os.Getenv("TWT_ACCESS_TOKEN"), os.Getenv("TWT_ACCESS_TOKEN_SECRET"))
	// http.Client should auto authorize reqs
	httpClient := cfg.Client(oauth1.NoContext, token)
	provider.Client = httpClient
	return provider, nil
}

type Provider struct {
	Client *http.Client
	Users  []string `yaml:"users"`
}
