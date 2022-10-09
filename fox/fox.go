package fox

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"aviary/insights"
	"aviary/twitter"

	"github.com/gin-gonic/gin"
)

type client struct {
	Twitter *twitter.Gateway
	Logger  *insights.Logger
	MediaID int64
	Content []byte
}

func Handler(c *gin.Context) {
	logger := insights.New()
	gateway, err := twitter.New()
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unable to setup")
		logger.Error("twitter.New: error", err)
		return
	}
	// Build client for handler workflow.
	client := &client{
		Twitter: gateway,
		Logger:  logger,
	}
	for _, user := range client.Twitter.Provider.Users {
		err = client.sendTweet(user)
		if err != nil {
			c.JSON(http.StatusBadRequest, "unable to tweet")
			// Not logging this error due to already existing error logging.
			return
		}
	}
	c.JSON(http.StatusOK, "tweet successful")
	return
}

// sendTweet builds opts and sends tweet.
func (c *client) sendTweet(user string) error {
	err := c.getFox()
	if err != nil {
		c.Logger.Error("getFox: error", err)
		return err
	}
	message := fmt.Sprintf("You're good! %s", user)
	err = c.Twitter.SendTweet(message, []int64{c.MediaID})
	if err != nil {
		c.Logger.Error("SendTweet: error", err)
		return err
	}
	return nil
}

// getFox collects fox image, MediaID; stores to pointer.
func (c *client) getFox() error {
	resp, err := http.Get("https://randomfox.ca/floof")
	if err != nil || resp.StatusCode != 200 {
		errMessage := fmt.Sprintf("http.Get error, status code: %v", resp.StatusCode)
		c.Logger.Error(errMessage, err)
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.Logger.Error("io.ReadAll: error", err)
		return err
	}
	resp.Body.Close()

	var foxMap map[string]string
	err = json.Unmarshal(body, &foxMap)
	if err != nil {
		c.Logger.Error("json.Unmarshal: error", err)
		return err
	}

	// get image content
	err = c.getContent(foxMap["image"])
	if err != nil {
		c.Logger.Error("getContent: error", err)
		return err
	}

	// Get twitter media ID
	media, _, err := c.Twitter.Session.Media.Upload(c.Content, "image/jpeg")
	if err != nil {
		c.Logger.Error("Media.Upload: error", err)
		return err
	}
	c.MediaID = media.MediaID
	return nil
}

// getContent called by getFox; stores image content in pointer.
func (c *client) getContent(url string) error {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		errMessage := fmt.Sprintf("http.Get error, status code: %v", resp.StatusCode)
		c.Logger.Error(errMessage, err)
		return err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Logger.Error("ioutil.ReadAll: error", err)
		return err
	}
	c.Content = content
	return nil
}
