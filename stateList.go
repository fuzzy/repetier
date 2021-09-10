package repetier

import (
	"encoding/json"
)

// StateListReturn blah
type StateListReturn struct {
	ActiveExtruder     int  `json:"activeExtruder"`
	AutostartNextPrint bool `json:"autostartNextPrint"`
	DebugLevel         int  `json:"debugLevel"`
	DoorOpen           bool `json:"doorOpen"`
	Extruder           []struct {
		Error    int     `json:"error"`
		Output   int     `json:"output"`
		TempRead float64 `json:"tempRead"`
		TempSet  int     `json:"tempSet"`
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
		Output   int     `json:"output"`
		TempRead float64 `json:"tempRead"`
		TempSet  int     `json:"tempSet"`
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
	X                  int           `json:"x"`
	Y                  int           `json:"y"`
	Z                  int           `json:"z"`
}

// StateList blah
func (obj *RestClient) StateList(s string, ih bool) map[string]*StateListReturn {
	retv := make(map[string]*StateListReturn)
	data := obj.Action("stateList", map[string]interface{}{"includeHistory": ih}, s)
	json.Unmarshal(data, &retv)
	return retv
}
