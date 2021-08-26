package repetier

import "encoding/json"

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
