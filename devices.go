package torchprint

import (
	"strconv"
	"time"
)

// Device a printing device
type Device struct {
	Identifier   string `json:"Identifier"`
	Name         string `json:"Name"`
	Description  string `json:"Description"`
	Make         string `json:"Make"`
	Model        string `json:"Model"`
	Capabilities struct {
		DuplexSupported bool `json:"DuplexSupported"`
		ColorSupported  bool `json:"ColorSupported"`
	} `json:"Capabilities"`
	Server       string    `json:"Server"`
	DeviceGroups []string  `json:"DeviceGroups"`
	Active       bool      `json:"Active"`
	LastModified time.Time `json:"LastModified"`
	Location     string    `json:"Location"`
}

// DevicesResponseData is response of devices request
type DevicesResponseData struct {
	PreviousPageLink string   `json:"PreviousPageLink"`
	NextPageLink     string   `json:"NextPageLink"`
	Count            int      `json:"Count"`
	Items            []Device `json:"Items"`
}

// GetDevices gets printing devices
func (api *API) GetDevices(skip, pageSize int) (*DevicesResponseData, error) {
	params := pagedRequest{
		Skip:     skip,
		PageSize: pageSize,
	}

	var respData DevicesResponseData
	_, err := api.GetJSON("devices", params, &respData)
	if err != nil {
		return nil, err
	}
	return &respData, nil
}

// GetDevice gets info about a printing device
func (api *API) GetDevice(deviceID int) (*Device, error) {
	var respData Device
	_, err := api.GetJSON("devices/"+strconv.Itoa(deviceID), nil, &respData)
	if err != nil {
		return nil, err
	}
	return &respData, nil
}
