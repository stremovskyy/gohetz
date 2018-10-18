// To parse and unparse this JSON data, add this code to your project and do:
//
//    deleteServer, err := UnmarshalDeleteServer(bytes)
//    bytes, err = deleteServer.Marshal()

package models

import "encoding/json"

func UnmarshalDeleteServer(data []byte) (DeleteServer, error) {
	var r DeleteServer
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteServer) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeleteServer struct {
	Action Action `json:"action"`
}
