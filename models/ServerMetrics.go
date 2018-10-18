// To parse and unparse this JSON data, add this code to your project and do:
//
//    serverMetrics, err := UnmarshalServerMetrics(bytes)
//    bytes, err = serverMetrics.Marshal()

package models

import "encoding/json"

func UnmarshalServerMetrics(data []byte) (ServerMetrics, error) {
	var r ServerMetrics
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ServerMetrics) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ServerMetrics struct {
	Metrics Metrics `json:"metrics"`
}

type Metrics struct {
	End        string  `json:"end"`         // End of period of metrics reported (in ISO-8601 format)
	Start      string  `json:"start"`       // Start of period of metrics reported (in ISO-8601 format)
	Step       float64 `json:"step"`        // Resolution of results in seconds.
	TimeSeries Series  `json:"time_series"` // Hash with timeseries information, containing the name of timeseries as key
}

// Hash with timeseries information, containing the name of timeseries as key
type Series struct {
	NameOfTimeseries interface{} `json:"name_of_timeseries"`
}
