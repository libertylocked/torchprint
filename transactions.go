package torchprint

import "time"

// Transaction is a tx in pharos (print, credit, etc)
type Transaction struct {
	Identifier      int       `json:"Identifier"`
	User            string    `json:"User"`
	Time            time.Time `json:"Time"`
	Amount          string    `json:"Amount"`
	Device          string    `json:"Device"`
	TransactionType string    `json:"TransactionType"`
	Description     string    `json:"Description"`
	Printer         string    `json:"Printer,omitempty"`
	Pages           int       `json:"Pages,omitempty"`
	Sheets          int       `json:"Sheets,omitempty"`
	Application     string    `json:"Application,omitempty"`
	JobName         string    `json:"JobName,omitempty"`
	Attributes      string    `json:"Attributes,omitempty"`
	Cashier         string    `json:"Cashier,omitempty"`
	Purse           string    `json:"Purse,omitempty"`
}

// TransactionResponseData response to get transactions
type TransactionResponseData struct {
	PreviousPageLink string        `json:"PreviousPageLink"`
	NextPageLink     string        `json:"NextPageLink"`
	Count            int           `json:"Count"`
	Items            []Transaction `json:"Items"`
}

// GetTransactions returns transaction history
func (api *API) GetTransactions(skip, pageSize int) (*TransactionResponseData, error) {
	params := pagedRequest{
		Skip:     skip,
		PageSize: pageSize,
	}

	var respData TransactionResponseData
	_, err := api.GetJSON("users/"+api.UserID+"/transactions", params, &respData)
	if err != nil {
		return nil, err
	}
	return &respData, nil
}
