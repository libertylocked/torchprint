package torchprint

import (
	"net/http"
	"time"

	"github.com/dghubble/sling"
)

// API is printer API wrapper
type API struct {
	Token string
}

const (
	baseURL = "https://mobileprint.nyu.edu/PharosAPI"
)

// Client returns a new sling client
func (api *API) Client() *sling.Sling {
	c := sling.New().
		Client(&http.Client{
			Timeout: time.Second * 10,
		}).
		Base(baseURL).
		Set("accept", "application/json")
	return c
}
