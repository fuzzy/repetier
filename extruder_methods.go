package repetier

// Update blah
func (obj *PrinterExtruder) Update() {
	return
}

// SetTemp will set the temperature on a given extruder for a given printer
func (obj *PrinterExtruder) SetTemp(t float64) {
	args := make(map[string]interface{})
	args["temperature"] = t
	args["extruder"] = obj.index
	obj.api.action("setExtruderTemperature", args, obj.slug)
}
