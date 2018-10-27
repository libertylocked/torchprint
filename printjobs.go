package torchprint

import "time"

// PrintJobsResponseData is the response of view printjob request
type PrintJobsResponseData struct {
	PreviousPageLink string `json:"PreviousPageLink"`
	NextPageLink     string `json:"NextPageLink"`
	Count            int    `json:"Count"`
	Items            []struct {
		ProtectedBy       string    `json:"ProtectedBy"`
		Owner             string    `json:"Owner"`
		Name              string    `json:"Name"`
		ApplicationName   string    `json:"ApplicationName"`
		JobFormat         string    `json:"JobFormat"`
		SubmissionTimeUtc time.Time `json:"SubmissionTimeUtc"`
		Stats             struct {
			SizeInKb        int `json:"SizeInKb"`
			TotalPages      int `json:"TotalPages"`
			TotalBWPages    int `json:"TotalBWPages"`
			TotalColorPages int `json:"TotalColorPages"`
			TotalSheets     int `json:"TotalSheets"`
		} `json:"Stats"`
		FinishingOptions struct {
			Duplex          bool   `json:"Duplex"`
			Mono            bool   `json:"Mono"`
			PagesPerSide    int    `json:"PagesPerSide"`
			Copies          int    `json:"Copies"`
			DefaultPageSize string `json:"DefaultPageSize"`
			PageRange       string `json:"PageRange"`
		} `json:"FinishingOptions"`
		SupportedFinishingOptions struct {
			Duplex       bool   `json:"Duplex"`
			Color        bool   `json:"Color"`
			PagesPerSide bool   `json:"PagesPerSide"`
			Copies       bool   `json:"Copies"`
			PageRange    string `json:"PageRange"`
		} `json:"SupportedFinishingOptions"`
		DeviceGroup          string    `json:"DeviceGroup"`
		PrinterName          string    `json:"PrinterName"`
		DocumentType         string    `json:"DocumentType"`
		AllowableActions     string    `json:"AllowableActions"`
		PrintState           string    `json:"PrintState"` // this is pretty important
		LastModified         time.Time `json:"LastModified"`
		Location             string    `json:"Location"`
		RawPageCounterResult string    `json:"RawPageCounterResult"`
		SubmissionTimeDelta  float64   `json:"SubmissionTimeDelta"`
		Cost                 string    `json:"Cost"`
	} `json:"Items"`
}

// GetPrintJobs sends a request to view print jobs
func (api *API) GetPrintJobs() (*PrintJobsResponseData, error) {
	var respData PrintJobsResponseData

	_, err := api.GetJSON("/PharosAPI/users/"+api.UserID+"/printjobs", nil, &respData)
	if err != nil {
		return nil, err
	}
	return &respData, err
}
