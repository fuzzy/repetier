package repetier

import (
	"encoding/json"
	"fmt"
)

// Update will query the current state and update the structure
func (obj *Server) Update() {
	args := make(map[string]interface{})
	info := obj.api.request("info", "", args, "")
	tdat := &ServerInfo{}
	json.Unmarshal(info, tdat)
	obj.Name = tdat.Servername
	obj.Software = tdat.Name
	obj.Version = tdat.Version
	for _, v := range tdat.Printers {
		obj.Printers[v.Slug] = newPrinter(obj.api, v.Slug)
	}
}

// Slugs returns a list of printer slugs connected to this server
func (obj *Server) Slugs() []string {
	retv := []string{}
	for k := range obj.Printers {
		retv = append(retv, k)
	}
	return retv
}

// TempPrinterState blah
type TempPrinterState struct {
	Active bool   `json:"active"`
	Job    string `json:"job"`
	Name   string `json:"name"`
	Online int    `json:"online"`
	Slug   string `json:"slug"`
}

// ListPrinters returns a list of printers and their states
func (obj *Server) ListPrinters() map[string]*TempPrinterState {
	temp := []*TempPrinterState{}
	retv := make(map[string]*TempPrinterState)
	_ = json.Unmarshal(obj.api.ListPrinter(), &temp)
	fmt.Printf("%+v\n", temp)
	for _, v := range temp {
		retv[v.Slug] = v
	}
	return retv
}
