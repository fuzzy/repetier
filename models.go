package repetier

import "encoding/json"

type ListModelsReturn struct {
	Analysed      int       `json:"analysed"`
	Created       int64     `json:"created"`
	ExtruderUsage []float64 `json:"extruderUsage"`
	FilamentTotal float64   `json:"filamentTotal"`
	Fits          bool      `json:"fits"`
	GcodePatch    string    `json:"gcodePatch"`
	Group         string    `json:"group"`
	ID            int       `json:"id"`
	LastPrintTime int       `json:"lastPrintTime"`
	Layer         int       `json:"layer"`
	Length        int       `json:"length"`
	Lines         int       `json:"lines"`
	Materials     []string  `json:"materials"`
	Name          string    `json:"name"`
	Notes         string    `json:"notes"`
	PrintTime     float64   `json:"printTime"`
	Printed       int       `json:"printed"`
	Radius        float64   `json:"radius"`
	RadiusMove    int       `json:"radiusMove"`
	Slicer        string    `json:"slicer"`
	State         string    `json:"state"`
	Version       int       `json:"version"`
	VolumeTotal   float64   `json:"volumeTotal"`
	VolumeUsage   []float64 `json:"volumeUsage"`
	Volumetric    bool      `json:"volumetric"`
	XMax          float64   `json:"xMax"`
	XMaxMove      float64   `json:"xMaxMove"`
	XMaxView      float64   `json:"xMaxView"`
	XMin          float64   `json:"xMin"`
	XMinMove      float64   `json:"xMinMove"`
	XMinView      float64   `json:"xMinView"`
	YMax          float64   `json:"yMax"`
	YMaxMove      float64   `json:"yMaxMove"`
	YMaxView      float64   `json:"yMaxView"`
	YMin          float64   `json:"yMin"`
	YMinMove      float64   `json:"yMinMove"`
	YMinView      float64   `json:"yMinView"`
	ZMax          float64   `json:"zMax"`
	ZMin          int       `json:"zMin"`
}

func (obj *RestClient) ListModels(g, s string) map[string][]*ListModelsReturn {
	retv := make(map[string][]*ListModelsReturn)
	data := obj.Action("listModels", map[string]interface{}{"group": g}, s)
	json.Unmarshal(data, &retv)
	return retv
}
