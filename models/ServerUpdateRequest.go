// To parse and unparse this JSON data, add this code to your project and do:
//
//    updateServerRequest, err := UnmarshalUpdateServerRequest(bytes)
//    bytes, err = updateServerRequest.Marshal()

package models

import "encoding/json"

func UnmarshalUpdateServerRequest(data []byte) (UpdateServerRequest, error) {
	var r UpdateServerRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateServerRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type UpdateServerRequest struct {
	Server Server `json:"server"`
}
