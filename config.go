package repetier

import "encoding/json"

// PrinterWebcam blah
type PrinterWebcam struct {
	DynamicURL            string  `json:"dynamicUrl"`
	ForceSnapshotPosition bool    `json:"forceSnapshotPosition"`
	Method                int     `json:"method"`
	Orientation           int     `json:"orientation"`
	Pos                   int     `json:"pos"`
	ReloadInterval        int     `json:"reloadInterval"`
	SnapshotDelay         int     `json:"snapshotDelay"`
	SnapshotX             int     `json:"snapshotX"`
	SnapshotY             int     `json:"snapshotY"`
	StaticURL             string  `json:"staticUrl"`
	TimelapseBitrate      int     `json:"timelapseBitrate"`
	TimelapseFramerate    int     `json:"timelapseFramerate"`
	TimelapseHeight       float64 `json:"timelapseHeight"`
	TimelapseInterval     int     `json:"timelapseInterval"`
	TimelapseLayer        int     `json:"timelapseLayer"`
	TimelapseMethod       int     `json:"timelapseMethod"`
	TimelapseSelected     int     `json:"timelapseSelected"`
}

// PrinterHeatedBed blah
type PrinterHeatedBed struct {
	Alias             string  `json:"alias"`
	CooldownPerSecond float64 `json:"cooldownPerSecond"`
	HeatupPerSecond   float64 `json:"heatupPerSecond"`
	LastTemp          int     `json:"lastTemp"`
	MaxTemp           int     `json:"maxTemp"`
	Offset            int     `json:"offset"`
	Temperatures      []struct {
		Name string `json:"name"`
		Temp int    `json:"temp"`
	} `json:"temperatures"`
}

// PrinterButtonCommand blah
type PrinterButtonCommand struct {
	Command string `json:"command"`
	Name    string `json:"name"`
}

// PrinterExtruder blah
type PrinterExtruder struct {
	Acceleration       int     `json:"acceleration"`
	Alias              string  `json:"alias"`
	ChangeFastDistance int     `json:"changeFastDistance"`
	ChangeSlowDistance int     `json:"changeSlowDistance"`
	CooldownPerSecond  float64 `json:"cooldownPerSecond"`
	EJerk              int     `json:"eJerk"`
	ExtrudeSpeed       int     `json:"extrudeSpeed"`
	FilamentDiameter   float64 `json:"filamentDiameter"`
	HeatupPerSecond    int     `json:"heatupPerSecond"`
	LastTemp           int     `json:"lastTemp"`
	MaxSpeed           int     `json:"maxSpeed"`
	MaxTemp            int     `json:"maxTemp"`
	Num                int     `json:"num"`
	Offset             int     `json:"offset"`
	OffsetX            int     `json:"offsetX"`
	OffsetY            int     `json:"offsetY"`
	RetractSpeed       int     `json:"retractSpeed"`
	SupportTemperature bool    `json:"supportTemperature"`
	TempMaster         int     `json:"tempMaster"`
	Temperatures       []struct {
		Name string `json:"name"`
		Temp int    `json:"temp"`
	} `json:"temperatures"`
	ToolDiameter float64 `json:"toolDiameter"`
	ToolType     int     `json:"toolType"`
}

// GetPrinterConfigReturn blah
type GetPrinterConfigReturn struct {
	ButtonCommands []*PrinterButtonCommand `json:"buttonCommands"`
	Connection     struct {
		ConnectionMethod           int  `json:"connectionMethod"`
		ContinueAfterFastReconnect bool `json:"continueAfterFastReconnect"`
		IP                         struct {
			Address string `json:"address"`
			Port    int    `json:"port"`
		} `json:"ip"`
		LcdTimeMode int    `json:"lcdTimeMode"`
		Password    string `json:"password"`
		Pipe        struct {
			File string `json:"file"`
		} `json:"pipe"`
		PowerOffIdleMinutes    int    `json:"powerOffIdleMinutes"`
		PowerOffMaxTemperature int    `json:"powerOffMaxTemperature"`
		ResetScript            string `json:"resetScript"`
		Serial                 struct {
			Baudrate              int    `json:"baudrate"`
			CommunicationTimeout  int    `json:"communicationTimeout"`
			ConnectionDelay       int    `json:"connectionDelay"`
			Device                string `json:"device"`
			Dtr                   int    `json:"dtr"`
			EmergencySolution     int    `json:"emergencySolution"`
			InputBufferSize       int    `json:"inputBufferSize"`
			Interceptor           bool   `json:"interceptor"`
			MalyanHack            bool   `json:"malyanHack"`
			MaxParallelCommands   int    `json:"maxParallelCommands"`
			PingPong              bool   `json:"pingPong"`
			Rts                   int    `json:"rts"`
			Usbreset              int    `json:"usbreset"`
			VisibleWithoutRunning bool   `json:"visibleWithoutRunning"`
		} `json:"serial"`
	} `json:"connection"`
	Extruders         []*PrinterExtruder `json:"extruders"`
	GcodeReplacements []interface{}      `json:"gcodeReplacements"`
	General           struct {
		Active                bool   `json:"active"`
		DefaultVolumetric     bool   `json:"defaultVolumetric"`
		DoorHandling          int    `json:"doorHandling"`
		EepromType            string `json:"eepromType"`
		FirmwareName          string `json:"firmwareName"`
		G9091OverrideE        bool   `json:"g9091OverrideE"`
		HeatedBed             bool   `json:"heatedBed"`
		LogHistory            bool   `json:"logHistory"`
		Manufacturer          string `json:"manufacturer"`
		Model                 string `json:"model"`
		Name                  string `json:"name"`
		NumFans               int    `json:"numFans"`
		PauseHandling         int    `json:"pauseHandling"`
		PauseSeconds          int    `json:"pauseSeconds"`
		PrinterHomepage       string `json:"printerHomepage"`
		PrinterManual         string `json:"printerManual"`
		PrinterVariant        string `json:"printerVariant"`
		Sdcard                bool   `json:"sdcard"`
		Slug                  string `json:"slug"`
		SoftwareLight         bool   `json:"softwareLight"`
		SoftwarePower         bool   `json:"softwarePower"`
		TempUpdateEvery       int    `json:"tempUpdateEvery"`
		UseModelFromSlug      string `json:"useModelFromSlug"`
		UseOwnModelRepository bool   `json:"useOwnModelRepository"`
	} `json:"general"`
	HeatedBeds     []*PrinterHeatedBed `json:"heatedBeds"`
	HeatedChambers []interface{}       `json:"heatedChambers"`
	Movement       struct {
		G10Distance                int     `json:"G10Distance"`
		G10LongDistance            int     `json:"G10LongDistance"`
		G10Speed                   int     `json:"G10Speed"`
		G10ZLift                   int     `json:"G10ZLift"`
		G11ExtraDistance           int     `json:"G11ExtraDistance"`
		G11ExtraLongDistance       int     `json:"G11ExtraLongDistance"`
		G11Speed                   int     `json:"G11Speed"`
		AllEndstops                bool    `json:"allEndstops"`
		Autolevel                  bool    `json:"autolevel"`
		DefaultAcceleration        int     `json:"defaultAcceleration"`
		DefaultRetractAcceleration int     `json:"defaultRetractAcceleration"`
		DefaultTravelAcceleration  int     `json:"defaultTravelAcceleration"`
		InvertX                    bool    `json:"invertX"`
		InvertY                    bool    `json:"invertY"`
		InvertZ                    bool    `json:"invertZ"`
		MaxXYSpeed                 int     `json:"maxXYSpeed"`
		MaxZSpeed                  int     `json:"maxZSpeed"`
		Movebuffer                 int     `json:"movebuffer"`
		TimeMultiplier             int     `json:"timeMultiplier"`
		XEndstop                   bool    `json:"xEndstop"`
		XHome                      int     `json:"xHome"`
		XMax                       int     `json:"xMax"`
		XMin                       int     `json:"xMin"`
		XyJerk                     int     `json:"xyJerk"`
		XyPrintAcceleration        int     `json:"xyPrintAcceleration"`
		XySpeed                    int     `json:"xySpeed"`
		XyTravelAcceleration       int     `json:"xyTravelAcceleration"`
		YEndstop                   bool    `json:"yEndstop"`
		YHome                      int     `json:"yHome"`
		YMax                       int     `json:"yMax"`
		YMin                       int     `json:"yMin"`
		ZEndstop                   bool    `json:"zEndstop"`
		ZHome                      int     `json:"zHome"`
		ZJerk                      float64 `json:"zJerk"`
		ZMax                       int     `json:"zMax"`
		ZMin                       int     `json:"zMin"`
		ZPrintAcceleration         int     `json:"zPrintAcceleration"`
		ZSpeed                     int     `json:"zSpeed"`
		ZTravelAcceleration        int     `json:"zTravelAcceleration"`
	} `json:"movement"`
	Properties struct {
	} `json:"properties"`
	QuickCommands []interface{} `json:"quickCommands"`
	Recover       struct {
		DelayBeforeReconnect    int    `json:"delayBeforeReconnect"`
		Enabled                 bool   `json:"enabled"`
		ExtraZOnFirmwareDetect  int    `json:"extraZOnFirmwareDetect"`
		FirmwarePowerlossSignal string `json:"firmwarePowerlossSignal"`
		MaxTimeForAutocontinue  int    `json:"maxTimeForAutocontinue"`
		Procedure               string `json:"procedure"`
		ReactivateBedOnConnect  bool   `json:"reactivateBedOnConnect"`
		ReplayExtruderSwitches  bool   `json:"replayExtruderSwitches"`
		RunOnConnect            string `json:"runOnConnect"`
	} `json:"recover"`
	ResponseEvents []interface{} `json:"responseEvents"`
	Shape          struct {
		BasicShape struct {
			Angle  int    `json:"angle"`
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
	Webcams []*PrinterWebcam `json:"webcams"`
}

// GetPrinterConfig blah
func (obj *RestClient) GetPrinterConfig(s string) *GetPrinterConfigReturn {
	retv := &GetPrinterConfigReturn{
		ButtonCommands: []*PrinterButtonCommand{},
		Extruders:      []*PrinterExtruder{},
		Webcams:        []*PrinterWebcam{},
	}
	json.Unmarshal(obj.Action("getPrinterConfig", map[string]interface{}{"printer": s}, s), retv)
	return retv
}

/*
// SetPrinterConfig blah
func (obj *RestClient) SetPrinterConfig(s string, c *GetPrinterConfigReturn) []byte {
	return obj.Action("setPrinterConfig", c, s)
}
*/
