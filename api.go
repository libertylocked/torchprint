package torchprint

import (
	"encoding/base64"
	"net/http"
	"time"

	"github.com/dghubble/sling"
	"github.com/libertylocked/torchprint/errors"
)

// API is printer API wrapper
type API struct {
	UserID string // required

	credential string
	token      string // if token is supplied, credential will not be used
}

const (
	// TODO: possibly support other schools using pharos?
	baseURL = "https://mobileprint.nyu.edu"
)

// NewAPI returns a new print API
func NewAPI(userID string) *API {
	return &API{
		UserID: userID,
	}
}

// SetToken set Pharos user token. If token is set, credential will be ignored
func (api *API) SetToken(token string) *API {
	api.token = token
	return api
}

// SetCredential set logon credentials. This is less secure than token, therefore
// it won't be used if token is set
func (api *API) SetCredential(credential string) *API {
	api.credential = credential
	return api
}

// SetUserPass set logon credentials from username and password
func (api *API) SetUserPass(user, pass string) *API {
	cred := base64.StdEncoding.EncodeToString([]byte(user + ":" + pass))
	return api.SetCredential(cred)
}

// Client returns a new sling client
func (api *API) client() *sling.Sling {
	c := sling.New().
		Client(&http.Client{
			Timeout: time.Second * 10,
		}).
		Base(baseURL)
	if len(api.credential) != 0 && len(api.token) == 0 {
		c.Set("X-Authorization", "PHAROS-USER "+api.credential)
	}
	// XXX: user ID is not known before logon, so client should call logon first
	// then set the it in API object. Maybe fix this ugly pattern later
	c.Set("Cookie", "PharosAPI.X-PHAROS-USER-URI=/users/"+api.UserID+
		"; PharosAPI.X-PHAROS-USER-TOKEN="+api.token)
	return c
}

// GetJSON populates a struct with JSON response data
func (api *API) GetJSON(url string, params interface{}, successDat interface{}) (*http.Response, error) {
	resp, err := api.client().
		Get(url).
		QueryStruct(params).
		ReceiveSuccess(successDat)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return resp, errors.NewHTTPError(resp)
	}
	return resp, err
}

// PostJSON posts a JSON and populates a struct with JSON response
func (api *API) PostJSON(url string, params interface{}, postBody interface{}, successDat interface{}) (*http.Response, error) {
	resp, err := api.client().
		Post(url).
		QueryStruct(params).
		BodyJSON(postBody).
		ReceiveSuccess(successDat)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return resp, errors.NewHTTPError(resp)
	}
	return resp, err
}
