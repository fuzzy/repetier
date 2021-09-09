package repetier

// PrinterConnection defines the serial, tcp/ip, UNIX socket connection
type PrinterConnection struct {
	Baudrate        int    `json:"baudrate"`
	Device          string `json:"device"`
	InputBufferSize int    `json:"inputBufferSize"`
	PingPong        bool   `json:"pingPong"`
	Protocol        int    `json:"protocol"`
}

// Temperature represents a saved temperature setting and it's label
type Temperature struct {
	Name string `json:"name"`
	Temp int    `json:"temp"`
}

// Printer represents a given printer connected to Server
type Printer struct {
	api        *RestClient
	slug       string
	Connection struct {
		Serial *PrinterConnection `json:"serial"`
	} `json:"connection"`
	Extruders []*PrinterExtruder `json:"extruders"`
	General   struct {
		Active          bool   `json:"active"`
		EepromType      string `json:"eepromType"`
		Fan             bool   `json:"fan"`
		FirmwareName    string `json:"firmwareName"`
		HeatedBed       bool   `json:"heatedBed"`
		Name            string `json:"name"`
		PrinterVariant  string `json:"printerVariant"`
		Sdcard          bool   `json:"sdcard"`
		Slug            string `json:"slug"`
		SoftwarePower   bool   `json:"softwarePower"`
		TempUpdateEvery int    `json:"tempUpdateEvery"`
	} `json:"general"`
	HeatedBed *PrinterHeatedBed `json:"heatedBed"`
	Movement  struct {
		AllEndstops          bool    `json:"allEndstops"`
		MaxXYSpeed           int     `json:"maxXYSpeed"`
		MaxZSpeed            int     `json:"maxZSpeed"`
		Movebuffer           int     `json:"movebuffer"`
		TimeMultiplier       int     `json:"timeMultiplier"`
		XEndstop             bool    `json:"xEndstop"`
		XHome                int     `json:"xHome"`
		XMax                 int     `json:"xMax"`
		XMin                 int     `json:"xMin"`
		XyJerk               int     `json:"xyJerk"`
		XyPrintAcceleration  int     `json:"xyPrintAcceleration"`
		XySpeed              int     `json:"xySpeed"`
		XyTravelAcceleration int     `json:"xyTravelAcceleration"`
		YEndstop             bool    `json:"yEndstop"`
		YHome                int     `json:"yHome"`
		YMax                 int     `json:"yMax"`
		YMin                 int     `json:"yMin"`
		ZEndstop             bool    `json:"zEndstop"`
		ZHome                int     `json:"zHome"`
		ZJerk                float64 `json:"zJerk"`
		ZMax                 int     `json:"zMax"`
		ZMin                 int     `json:"zMin"`
		ZPrintAcceleration   int     `json:"zPrintAcceleration"`
		ZSpeed               int     `json:"zSpeed"`
		ZTravelAcceleration  int     `json:"zTravelAcceleration"`
	} `json:"movement"`
	QuickCommands []interface{} `json:"quickCommands"`
	Shape         struct {
		BasicShape struct {
			Color  string `json:"color"`
			Radius int    `json:"radius"`
			Shape  string `json:"shape"`
			X      int    `json:"x"`
			XMax   int    `json:"xMax"`
			XMin   int    `json:"xMin"`
			Y      int    `json:"y"`
			YMax   int    `json:"yMax"`
			YMin   int    `json:"yMin"`
		} `json:"basicShape"`
		GridColor   string        `json:"gridColor"`
		GridSpacing int           `json:"gridSpacing"`
		Marker      []interface{} `json:"marker"`
	} `json:"shape"`
	State  *PrinterState
	Webcam []*PrinterWebcam `json:"webcam"`
}

// PrinterState defines current state of a Printer
type PrinterState struct {
	ActiveExtruder     int  `json:"activeExtruder"`
	AutostartNextPrint bool `json:"autostartNextPrint"`
	DebugLevel         int  `json:"debugLevel"`
	DoorOpen           bool `json:"doorOpen"`
	Extruder           []struct {
		Error    int     `json:"error"`
		Output   float64 `json:"output"`
		TempRead float64 `json:"tempRead"`
		TempSet  float64 `json:"tempSet"`
	} `json:"extruder"`
	Fans []struct {
		On      bool `json:"on"`
		Voltage int  `json:"voltage"`
	} `json:"fans"`
	FilterFan    bool   `json:"filterFan"`
	Firmware     string `json:"firmware"`
	FirmwareURL  string `json:"firmwareURL"`
	FlowMultiply int    `json:"flowMultiply"`
	HasXHome     bool   `json:"hasXHome"`
	HasYHome     bool   `json:"hasYHome"`
	HasZHome     bool   `json:"hasZHome"`
	HeatedBeds   []struct {
		Error    int     `json:"error"`
		Output   float64 `json:"output"`
		TempRead float64 `json:"tempRead"`
		TempSet  float64 `json:"tempSet"`
	} `json:"heatedBeds"`
	HeatedChambers     []interface{} `json:"heatedChambers"`
	Layer              int           `json:"layer"`
	Lights             int           `json:"lights"`
	Notification       string        `json:"notification"`
	NumExtruder        int           `json:"numExtruder"`
	PowerOn            bool          `json:"powerOn"`
	Rec                bool          `json:"rec"`
	SdcardMounted      bool          `json:"sdcardMounted"`
	ShutdownAfterPrint bool          `json:"shutdownAfterPrint"`
	SpeedMultiply      int           `json:"speedMultiply"`
	Volumetric         bool          `json:"volumetric"`
	X                  float64       `json:"x"`
	Y                  float64       `json:"y"`
	Z                  float64       `json:"z"`
}

// NewPrinter returns a Printer object
func newPrinter(a *RestClient, s string) *Printer {
	retv := &Printer{
		api:       a,
		slug:      s,
		Extruders: []*PrinterExtruder{},
		State:     &PrinterState{},
		Webcam:    []*PrinterWebcam{&PrinterWebcam{}},
	}
	/*
			HeatedBed: &PrinterHeatedBed{},
			State:     &PrinterState{},
			Webcam:    &PrinterWebcam{},
		}
	*/
	retv.HeatedBed = newHeatedBed(a, s)
	retv.Webcam = []*PrinterWebcam{newWebcam(a, s)}
	retv.Update()
	return retv
}
