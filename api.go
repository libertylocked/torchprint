package torchprint

import (
	"net/http"
	"time"

	"github.com/dghubble/sling"
	"github.com/libertylocked/torchprint/errors"
)

// API is printer API wrapper
type API struct {
	// Base64 encoded "user:pass"
	Credentials string
	UserID      string
}

const (
	baseURL = "https://mobileprint.nyu.edu"
)

// NewAPI returns a new print API
func NewAPI(credentials, userID string) *API {
	return &API{
		Credentials: credentials,
		UserID:      userID,
	}
}

// Client returns a new sling client
func (api *API) Client() *sling.Sling {
	c := sling.New().
		Client(&http.Client{
			Timeout: time.Second * 10,
		}).
		Base(baseURL).
		Set("accept", "application/json").
		// XXX: not very secure, maybe replace it with session?
		Set("X-Authorization", "PHAROS-USER "+api.Credentials).
		// XXX: user ID is not known before logon, so client should call logon first
		// then set the it in API object. Maybe fix this ugly pattern later
		Set("Cookie", "PharosAPI.X-PHAROS-USER-URI=/users/"+api.UserID)
	return c
}

// GetJSON populates a struct with JSON response data
func (api *API) GetJSON(url string, params interface{}, successDat interface{}) (*http.Response, error) {
	resp, err := api.Client().
		Get(url).
		QueryStruct(params).
		ReceiveSuccess(successDat)
	if resp.StatusCode != http.StatusOK {
		return resp, errors.NewHTTPError(resp)
	}
	return resp, err
}
