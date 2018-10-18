// To parse and unparse this JSON data, add this code to your project and do:
//
//    deleteServer, err := UnmarshalDeleteServer(bytes)
//    bytes, err = deleteServer.Marshal()

package models

import "encoding/json"

func UnmarshalServerDeleteResponse(data []byte) (ServerDeleteResponse, error) {
	var r ServerDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

type ServerDeleteResponse struct {
	Action ActionClass `json:"action"`
}
