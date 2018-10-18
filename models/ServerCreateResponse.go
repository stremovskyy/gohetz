// To parse and unparse this JSON data, add this code to your project and do:
//
//    createServerResponse, err := UnmarshalCreateServerResponse(bytes)
//    bytes, err = createServerResponse.Marshal()

package models

import "encoding/json"

func UnmarshalCreateServerResponse(data []byte) (CreateServerResponse, error) {
	var r CreateServerResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateServerResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CreateServerResponse struct {
	Action       Action  `json:"action"`
	RootPassword *string `json:"root_password"` // Root password when no SSH keys have been specified
	Server       Server  `json:"server"`
}
