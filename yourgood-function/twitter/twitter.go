package twitter

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"youRgood/config"

	"github.com/drswork/go-twitter/twitter"
	"github.com/gin-gonic/gin"
)

type client struct {
	Provider config.Provider
	Twitter  *twitter.Client
	MediaID  int64
	Content  []byte
}

func Handler(c *gin.Context) {
	client, err := Setup()
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unable to setup")
		log.Fatalf("unable to setup: %v", err)
		return
	}
	for _, user := range client.Provider.Users {
		err = client.sendTweet(user)
		if err != nil {
			c.JSON(http.StatusBadRequest, "unable to tweet")
			log.Fatalf("unable to tweet: %v", err)
			return
		}
	}
	c.JSON(http.StatusOK, "tweet successful")
	return
}

func Setup() (*client, error) {
	provider, err := config.New()
	if err != nil {
		return nil, err
	}
	return &client{
		Provider: provider,
		Twitter:  twitter.NewClient(provider.Client),
	}, nil
}

func (c *client) sendTweet(user string) error {
	err := c.getFox()
	if err != nil {
		return err
	}
	opts := &twitter.StatusUpdateParams{
		Status:   fmt.Sprintf("You're good! @%s", user),
		MediaIds: []int64{c.MediaID},
	}
	_, resp, err := c.Twitter.Statuses.Update(opts.Status, opts)
	if err != nil || resp.StatusCode != 200 {
		return err
	}
	return nil
}

func (c *client) getFox() error {
	resp, err := http.Get("https://randomfox.ca/floof")
	if err != nil || resp.StatusCode != 200 {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	resp.Body.Close()

	var foxMap map[string]string
	err = json.Unmarshal(body, &foxMap)
	if err != nil {
		return err
	}

	// get image content
	err = c.getContent(foxMap["image"])
	if err != nil {
		return err
	}

	// Get twitter media ID
	media, _, err := c.Twitter.Media.Upload(c.Content, "image/jpeg")
	if err != nil {
		return err
	}
	c.MediaID = media.MediaID
	return nil
}

func (c *client) getContent(url string) error {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	c.Content = content
	return nil
}
