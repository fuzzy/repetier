package repetier

// PrinterWebcam defines a webcam object associated with a specific printer
type PrinterWebcam struct {
	api               *client
	slug              string
	DynamicURL        string `json:"dynamicUrl"`
	Method            int    `json:"method"`
	ReloadInterval    int    `json:"reloadInterval"`
	StaticURL         string `json:"staticUrl"`
	TimelapseInterval int    `json:"timelapseInterval"`
	TimelapseMethod   int    `json:"timelapseMethod"`
}

func newWebcam(a *client, s string) *PrinterWebcam {
	retv := &PrinterWebcam{api: a, slug: s}
	retv.Update()
	return retv
}
