package torchprint

type pagedRequest struct {
	Skip     int `url:"Skip,omitempty"`
	PageSize int `url:"PageSize,omitempty"`
}
