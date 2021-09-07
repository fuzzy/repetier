package repetier

// PrinterExtruder structure for the Extruders present on a Printer
type PrinterExtruder struct {
	api               *RestClient
	slug              string
	index             int
	Acceleration      int            `json:"acceleration"`
	CooldownPerSecond float64        `json:"cooldownPerSecond"`
	EJerk             int            `json:"eJerk"`
	ExtrudeSpeed      int            `json:"extrudeSpeed"`
	FilamentDiameter  float64        `json:"filamentDiameter"`
	HeatupPerSecond   float64        `json:"heatupPerSecond"`
	LastTemp          int            `json:"lastTemp"`
	MaxSpeed          int            `json:"maxSpeed"`
	MaxTemp           int            `json:"maxTemp"`
	OffsetX           int            `json:"offsetX"`
	OffsetY           int            `json:"offsetY"`
	RetractSpeed      int            `json:"retractSpeed"`
	Temperatures      []*Temperature `json:"temperatures"`
}

func newExtruder(a *RestClient, s string, i int) *PrinterExtruder {
	retv := &PrinterExtruder{api: a, slug: s, index: i}
	retv.Update()
	return retv
}
