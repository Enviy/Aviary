package twitter

import (
	"youRgood/config"

	twt "github.com/drswork/go-twitter/twitter"
)

type Gateway struct {
	Provider config.Provider
	Session  *twt.Client
}

func New() (*Gateway, error) {
	provider, err := config.New()
	if err != nil {
		return nil, err
	}
	return &Gateway{
		Provider: provider,
		Session:  twt.NewClient(provider.Client),
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
