package repetier

import "encoding/json"

// Update blah
func (obj *PrinterExtruder) Update() {
	temp := &Printer{api: obj.api, slug: obj.slug}
	args := make(map[string]interface{})
	args["printer"] = obj.slug
	data := obj.api.Action("getPrinterConfig", args, obj.slug)
	json.Unmarshal(data, temp)
	// slap the bits we need back to json, because it's easier
	// I do recognize that it is WAAY more work for the cpu but
	// atm, I am not that concerned. I'll optimize later.
	tdat, _ := json.Marshal(temp.Extruders[obj.index])
	json.Unmarshal(tdat, obj)
	return
}

// SetTemp will set the temperature on a given extruder for a given printer
func (obj *PrinterExtruder) SetTemp(t float64) {
	args := make(map[string]interface{})
	args["temperature"] = t
	args["extruder"] = obj.index
	obj.api.Action("setExtruderTemperature", args, obj.slug)
}
