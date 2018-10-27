package torchprint

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
	Privileges    struct {
		UserAdministration struct {
			View           string `json:"View"`
			Update         string `json:"Update"`
			Delete         string `json:"Delete"`
			Create         string `json:"Create"`
			Credit         string `json:"Credit"`
			Debit          string `json:"Debit"`
			SetBalance     string `json:"SetBalance"`
			BlankPassword  string `json:"BlankPassword"`
			StrongPassword string `json:"StrongPassword"`
			Supported      string `json:"Supported"`
		} `json:"UserAdministration"`
		UserGroupAdministration struct {
			View   string `json:"View"`
			Update string `json:"Update"`
			Delete string `json:"Delete"`
			Create string `json:"Create"`
		} `json:"UserGroupAdministration"`
		Printing struct {
			Supported    string `json:"Supported"`
			View         string `json:"View"`
			WebUpload    string `json:"WebUpload"`
			Confirmation string `json:"Confirmation"`
			PayForPrint  struct {
				Costing           string `json:"Costing"`
				CostCenters       string `json:"CostCenters"`
				CreditCardGateway string `json:"CreditCardGateway"`
				AddFunds          string `json:"AddFunds"`
			} `json:"PayForPrint"`
			Administration struct {
				View               string `json:"View"`
				Delete             string `json:"Delete"`
				ChangeChargingUser string `json:"ChangeChargingUser"`
				Supported          string `json:"Supported"`
			} `json:"Administration"`
			ColorPrinting string `json:"ColorPrinting"`
		} `json:"Printing"`
		PPCTheme struct {
			Supported string `json:"Supported"`
			View      string `json:"View"`
			Update    string `json:"Update"`
		} `json:"PPCTheme"`
		Reports struct {
			Supported             string `json:"Supported"`
			CreditCardFundsReport string `json:"CreditCardFundsReport"`
		} `json:"Reports"`
		Activity struct {
			Supported      string `json:"Supported"`
			View           string `json:"View"`
			Administration struct {
				Supported string `json:"Supported"`
				View      string `json:"View"`
			} `json:"Administration"`
		} `json:"Activity"`
		MobilePrintAdministration struct {
			Supported string `json:"Supported"`
			View      string `json:"View"`
		} `json:"MobilePrintAdministration"`
		QuotaPrivileges struct {
			View   string `json:"View"`
			Update string `json:"Update"`
		} `json:"QuotaPrivileges"`
		ManualTransactions struct {
			Supported string `json:"Supported"`
			View      string `json:"View"`
			Create    string `json:"Create"`
			Update    string `json:"Update"`
			Delete    string `json:"Delete"`
		} `json:"ManualTransactions"`
		Costing string `json:"Costing"`
	} `json:"Privileges"`
}

// Logon sends a logon request
func (api *API) Logon() (*LogonResponseData, error) {
	var respData LogonResponseData
	queryParams := LogonRequestData{
		KeepMeLoggedIn:   "yes",
		IncludePrintJobs: "yes",
	}

	_, err := api.GetJSON("/PharosAPI/logon", queryParams, &respData)
	if err != nil {
		return nil, err
	}
	return &respData, err
}
