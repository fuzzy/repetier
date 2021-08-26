package repetier

import (
	"encoding/json"
)

// Update blah
func (obj *Printer) Update() {
	// First let's get our printer config and update ourselves
	args := make(map[string]interface{})
	args["printer"] = obj.slug
	data := obj.api.action("getPrinterConfig", args, obj.slug)
	json.Unmarshal(data, obj)
	// update all the Extruders
	for i := range obj.Extruders {
		obj.Extruders[i].index = i
		obj.Extruders[i].api = obj.api
		obj.Extruders[i].slug = obj.slug
	}
	// get the state list, which despite requiring a slug, returns all printers.
	stateArgs := make(map[string]interface{})
	data = obj.api.action("stateList", stateArgs, obj.slug)
	tmpd := make(map[string]*PrinterState)
	// Unmarshal the data into a map[string]*Printer as a temporary step
	json.Unmarshal(data, &tmpd)
	// Then to update our `obj` we marshal the appropriate chunk
	tmjd, err := json.Marshal(tmpd[obj.slug])
	panicCheck(err)
	// and unmarshal it to our `obj`
	json.Unmarshal(tmjd, obj.State)
}

// AbsMove blah
func (obj *Printer) AbsMove(x, y, z, e float64) {
	obj.Move(x, y, z, e, false)
}

// RelMove blah
func (obj *Printer) RelMove(x, y, z, e float64) {
	obj.Move(x, y, z, e, true)
}

// Move blah
func (obj *Printer) Move(x, y, z, e float64, relative bool) {
	args := make(map[string]interface{})
	args["x"] = x
	args["y"] = y
	args["z"] = z
	args["e"] = e
	args["relative"] = relative
	obj.api.action("move", args, obj.slug)
}
