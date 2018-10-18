// To parse and unparse this JSON data, add this code to your project and do:
//
//    bytes, err = serverUpdateRequest.Marshal()

package models

import "encoding/json"

func (r *ServerUpdateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ServerUpdateRequest struct {
	Labels map[string]interface{} `json:"labels,omitempty"` // New labels
	Name   *string                `json:"name,omitempty"`   // New name to set
}

func UnmarshalServerUpdateResponse(data []byte) (ServerUpdateResponse, error) {
	var r ServerUpdateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ServerUpdateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ServerUpdateResponse struct {
	Server ServerClass `json:"server"`
}
