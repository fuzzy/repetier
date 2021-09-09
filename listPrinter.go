package repetier

import (
	"encoding/json"
)

// ListPrinterReturn blah
type ListPrinterReturn struct {
	Active          bool    `json:"active"`
	Job             string  `json:"job"`
	Name            string  `json:"name"`
	Online          int     `json:"online"`
	PauseState      int     `json:"pauseState"`
	Paused          bool    `json:"paused"`
	Slug            string  `json:"slug"`
	Analysed        int     `json:"analysed,omitempty"`
	Done            float64 `json:"done,omitempty"`
	Jobid           int     `json:"jobid,omitempty"`
	LinesSend       int     `json:"linesSend,omitempty"`
	OfLayer         int     `json:"ofLayer,omitempty"`
	PrintStart      float64 `json:"printStart,omitempty"`
	PrintTime       float64 `json:"printTime,omitempty"`
	PrintedTimeComp float64 `json:"printedTimeComp,omitempty"`
	Start           int     `json:"start,omitempty"`
	TotalLines      int     `json:"totalLines,omitempty"`
}

// ListPrinter blah
func (obj *RestClient) ListPrinter() []*ListPrinterReturn {
	retv := []*ListPrinterReturn{}
	data := obj.Action("listPrinter", map[string]interface{}{}, "")
	json.Unmarshal(data, &retv)
	return retv
}
