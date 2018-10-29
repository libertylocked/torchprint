package torchprint

import (
	"github.com/libertylocked/torchprint/errors"
)

// LogonRequestData is query params of logon request
type LogonRequestData struct {
	KeepMeLoggedIn        string `url:"KeepMeLoggedIn"`
	IncludePrintJobs      string `url:"includeprintjobs"`
	IncludeDeviceActivity string `url:"includedeviceactivity"`
	IncludePrivileges     string `url:"includeprivileges"`
	IncludeCostCenters    string `url:"includecostcenters"`
}

// LogonResponseData is the response of logon request
type LogonResponseData struct {
	Identifier     string   `json:"Identifier"` // this is user ID!
	LogonID        string   `json:"LogonId"`
	DisplayName    string   `json:"DisplayName"`
	EmailAddresses []string `json:"EmailAddresses"`
	Roles          []string `json:"Roles"`
	Active         bool     `json:"Active"`
	Alias          string   `json:"Alias"`
	AccountType    string   `json:"AccountType"`
	Balance        struct {
		Amount string `json:"Amount"`
		Purses []struct {
			Name   string `json:"Name"`
			Amount string `json:"Amount"`
		} `json:"Purses"`
	} `json:"Balance"`
	Location      string `json:"Location"`
	Group         string `json:"Group"`
	Comment       string `json:"Comment"`
	Custom1       string `json:"Custom1"`
	Custom2       string `json:"Custom2"`
	OfflineAmount string `json:"OfflineAmount"`
	OfflineLimit  string `json:"OfflineLimit"`
	ID            int    `json:"Id"`
}

// Logon sends a request to get token. Credential must be set in API object
func (api *API) Logon() (respData *LogonResponseData, token string, err error) {
	if len(api.credential) == 0 {
		return nil, "", errors.NoCredentialError{}
	}
	// TODO: allow customizing the request
	queryParams := LogonRequestData{
		KeepMeLoggedIn:        "yes",
		IncludePrintJobs:      "no",
		IncludeDeviceActivity: "no",
		IncludePrivileges:     "no",
		IncludeCostCenters:    "no",
	}
	resp, err := api.GetJSON("logon", queryParams, &respData)
	if err != nil {
		return nil, "", err
	}
	// get the token from cookie
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "PharosAPI.X-PHAROS-USER-TOKEN" {
			token = cookie.Value
		}
	}
	return respData, token, err
}
