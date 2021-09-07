package repetier

// PrinterHeatedBed defines a heated bed on a given Printer
type PrinterHeatedBed struct {
	api               *RestClient
	slug              string
	CooldownPerSecond int            `json:"cooldownPerSecond"`
	HeatupPerSecond   int            `json:"heatupPerSecond"`
	Installed         bool           `json:"installed"`
	LastTemp          float64        `json:"lastTemp"`
	MaxTemp           float64        `json:"maxTemp"`
	Temperatures      []*Temperature `json:"temperatures"`
}

// NewHeatedBed blah
func newHeatedBed(api *RestClient, slug string) *PrinterHeatedBed {
	retv := &PrinterHeatedBed{api: api, slug: slug}
	retv.Update()
	return retv
}
