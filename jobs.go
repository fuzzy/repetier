package repetier

import (
	"encoding/json"
)

// ListJobsReturn blah
type ListJobsReturn struct {
	Analysed        int       `json:"analysed"`
	Created         int64     `json:"created"`
	ExtruderUsage   []float64 `json:"extruderUsage"`
	FilamentTotal   float64   `json:"filamentTotal"`
	Fits            bool      `json:"fits"`
	GcodePatch      string    `json:"gcodePatch"`
	Group           string    `json:"group"`
	ID              int       `json:"id"`
	LastPrintTime   int       `json:"lastPrintTime"`
	Layer           int       `json:"layer"`
	Length          int       `json:"length"`
	Lines           int       `json:"lines"`
	Materials       []string  `json:"materials"`
	Name            string    `json:"name"`
	Notes           string    `json:"notes"`
	PrintTime       float64   `json:"printTime"`
	Printed         int       `json:"printed"`
	PrintedTimeComp float64   `json:"printedTimeComp"`
	Radius          float64   `json:"radius"`
	RadiusMove      int       `json:"radiusMove"`
	Slicer          string    `json:"slicer"`
	State           string    `json:"state"`
	Version         int       `json:"version"`
	VolumeTotal     float64   `json:"volumeTotal"`
	VolumeUsage     []float64 `json:"volumeUsage"`
	Volumetric      bool      `json:"volumetric"`
	XMax            float64   `json:"xMax"`
	XMaxMove        float64   `json:"xMaxMove"`
	XMaxView        float64   `json:"xMaxView"`
	XMin            float64   `json:"xMin"`
	XMinMove        float64   `json:"xMinMove"`
	XMinView        float64   `json:"xMinView"`
	YMax            float64   `json:"yMax"`
	YMaxMove        float64   `json:"yMaxMove"`
	YMaxView        float64   `json:"yMaxView"`
	YMin            float64   `json:"yMin"`
	YMinMove        float64   `json:"yMinMove"`
	YMinView        float64   `json:"yMinView"`
	ZMax            float64   `json:"zMax"`
	ZMin            int       `json:"zMin"`
}

// ListJobs blah
func (obj *RestClient) ListJobs(s string) map[string][]*ListJobsReturn {
	retv := make(map[string][]*ListJobsReturn)
	data := obj.Action("listJobs", map[string]interface{}{}, s)
	json.Unmarshal(data, &retv)
	return retv
}

func (obj *RestClient) StartJob(s string, id int) {
	obj.Action("startJob", map[string]interface{}{"id": id}, s)
}
