package repetier

// Update blah
func (obj *PrinterHeatedBed) Update() {
	return
}

// SetTemp will set the temperature on a given extruder for a given printer
func (obj *PrinterHeatedBed) SetTemp(t float64) {
	args := make(map[string]interface{})
	args["temperature"] = t
	obj.api.action("setBedTemperature", args, obj.slug)
}
