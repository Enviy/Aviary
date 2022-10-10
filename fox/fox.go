package fox

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"aviary/app"

	"github.com/gin-gonic/gin"
)

type client struct {
	App     app.Gateway
	MediaID int64
	Content []byte
}

// Handler behaves more like a controller,
// it's the entry point for package's business logic.
func Handler(c *gin.Context) {
	app, err := app.New()
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unable to setup")
		return
	}
	// Using client pointer to share data among methods.
	client := &client{
		App: app,
	}
	for _, user := range app.Users {
		err = client.getFox()
		if err != nil {
			app.Logger.Error("getFox: error", err)
			c.JSON(http.StatusUnauthorized, "unable to get fox")
			return
		}
		message := fmt.Sprintf("You're good! %s", user)
		err = client.App.Twitter.SendTweet(message, []int64{client.MediaID})
		if err != nil {
			app.Logger.Error("SendTweet: error", err)
			c.JSON(http.StatusUnauthorized, "unable to tweet")
			return
		}
	}
	c.JSON(http.StatusOK, "tweet successful")
	return
}

// getFox collects fox image, MediaID; stores to pointer.
func (c *client) getFox() error {
	resp, err := http.Get("https://randomfox.ca/floof")
	if err != nil || resp.StatusCode != 200 {
		errMessage := fmt.Sprintf("http.Get error, status code: %v", resp.StatusCode)
		c.App.Logger.Error(errMessage, err)
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.App.Logger.Error("io.ReadAll: error", err)
		return err
	}
	resp.Body.Close()

	var foxMap map[string]string
	err = json.Unmarshal(body, &foxMap)
	if err != nil {
		c.App.Logger.Error("json.Unmarshal: error", err)
		return err
	}

	// get image content
	err = c.getContent(foxMap["image"])
	if err != nil {
		c.App.Logger.Error("getContent: error", err)
		return err
	}

	// Get twitter media ID
	media, _, err := c.App.Twitter.Session.Media.Upload(c.Content, "image/jpeg")
	if err != nil {
		c.App.Logger.Error("Media.Upload: error", err)
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
		c.App.Logger.Error(errMessage, err)
		return err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.App.Logger.Error("ioutil.ReadAll: error", err)
		return err
	}
	c.Content = content
	return nil
}
