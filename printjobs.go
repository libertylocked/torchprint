package torchprint

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/libertylocked/torchprint/errors"
)

// FinishingOptions finishing options of a printjob
type FinishingOptions struct {
	Mono            bool   `json:"Mono"`
	Duplex          bool   `json:"Duplex"`
	PagesPerSide    int    `json:"PagesPerSide"`
	Copies          int    `json:"Copies"`
	DefaultPageSize string `json:"DefaultPageSize"` // Letter or something
	PageRange       string `json:"PageRange"`
}

// PrintJobItem is an item in printjobs
type PrintJobItem struct {
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
	FinishingOptions FinishingOptions `json:"FinishingOptions"`
	// SupportedFinishingOptions FinishingOptions `json:"SupportedFinishingOptions"` // its a bool map, dont need for now
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
}

// PrintJobsResponseData is the response of view printjob request
type PrintJobsResponseData struct {
	PreviousPageLink string         `json:"PreviousPageLink"`
	NextPageLink     string         `json:"NextPageLink"`
	Count            int            `json:"Count"`
	Items            []PrintJobItem `json:"Items"`
}

type addPrintJobRequestData struct {
	FinishingOptions struct {
		Mono            bool   `json:"Mono"`
		Duplex          bool   `json:"Duplex"`
		PagesPerSide    string `json:"PagesPerSide"`
		Copies          string `json:"Copies"`
		DefaultPageSize string `json:"DefaultPageSize"`
		PageRange       string `json:"PageRange"`
	} `json:"FinishingOptions"`
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

// AddPrintJob adds a printjob
func (api *API) AddPrintJob(filename string, options FinishingOptions) (*PrintJobItem, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bodyform := &bytes.Buffer{}
	writer := multipart.NewWriter(bodyform)

	metapart, err := writer.CreateFormField("MetaData")
	if err != nil {
		return nil, err
	}
	// XXX: pretty sure it's a bug in pharos web API, that PagesPerSide and Copies
	// is string in the request (server will 500 if they are int)
	// sorry about this nasty hack
	reqMeta := addPrintJobRequestData{}
	reqMeta.FinishingOptions.Copies = strconv.Itoa(options.Copies)
	reqMeta.FinishingOptions.DefaultPageSize = options.DefaultPageSize
	reqMeta.FinishingOptions.Duplex = options.Duplex
	reqMeta.FinishingOptions.Mono = options.Mono
	reqMeta.FinishingOptions.PageRange = options.PageRange
	reqMeta.FinishingOptions.PagesPerSide = strconv.Itoa(options.PagesPerSide)
	err = json.NewEncoder(metapart).Encode(reqMeta)
	if err != nil {
		return nil, err
	}

	filepart, err := writer.CreateFormFile("content", filepath.Base(filename))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(filepart, file)
	if err != nil {
		return nil, err
	}

	// finally send the POST
	writer.Close()
	var respData PrintJobItem
	resp, err := api.client().Body(bodyform).
		Set("Content-Type", writer.FormDataContentType()).
		Post("/PharosAPI/users/" + api.UserID + "/printjobs").
		ReceiveSuccess(&respData)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusCreated {
		return &respData, errors.NewHTTPError(resp)
	}

	return &respData, nil
}
