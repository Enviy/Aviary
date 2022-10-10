package twitter

import (
	"aviary/config"

	"github.com/dghubble/oauth1"
	twt "github.com/drswork/go-twitter/twitter"
)

type Gateway struct {
	Session *twt.Client
}

func New(prov config.Provider) (*Gateway, error) {
	cfg := oauth1.NewConfig(prov.Twitter.Key, prov.Twitter.KeySecret)
	token := oauth1.NewToken(prov.Twitter.Token, prov.Twitter.TokenSecret)
	httpClient := cfg.Client(oauth1.NoContext, token)
	return &Gateway{
		Session: twt.NewClient(httpClient),
	}, nil
}

// SendTweet accepts message and media IDs
func (c *Gateway) SendTweet(message string, mediaIDs []int64) error {
	opts := &twt.StatusUpdateParams{
		Status:   message,
		MediaIds: mediaIDs,
	}
	_, resp, err := c.Session.Statuses.Update(opts.Status, opts)
	if err != nil || resp.StatusCode != 200 {
		return err
	}
	return nil
}
