// To parse and unparse this JSON data, add this code to your project and do:
//
//    createServerRequest, err := UnmarshalCreateServerRequest(bytes)
//    bytes, err = createServerRequest.Marshal()

package models

import "encoding/json"

func (r *ServerCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ServerCreateRequest struct {
	Image            string                 `json:"image"`                        // ID or name of the image the server is created from
	Labels           map[string]interface{} `json:"labels,omitempty"`             // User-defined labels (key-value pairs)
	Name             string                 `json:"name"`                         // Name of the server to create (must be unique per project and a valid hostname as per RFC; 1123)
	ServerType       string                 `json:"server_type"`                  // ID or name of the server type this server should be created with
	SSHKeys          []string               `json:"ssh_keys,omitempty"`           // SSH key IDs or names which should be injected into the server at creation time
	StartAfterCreate *bool                  `json:"start_after_create,omitempty"` // Start Server right after creation. Defaults to true.
	UserData         *string                `json:"user_data,omitempty"`          // Cloud-Init user data to use during server creation. This field is limited to 32KiB.
	Volumes          []string               `json:"volumes,omitempty"`            // Volume IDs which should be attached to the server at the creation time. Volumes must be; in the same location.
	Location         *string                `json:"location,omitempty"`           // ID or name of location to create server in.
	Datacenter       *string                `json:"datacenter,omitempty"`         // ID or name of datacenter to create server in.
}

func UnmarshalCreateServerResponse(data []byte) (ServerCreateResponse, error) {
	var r ServerCreateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ServerCreateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ServerCreateResponse struct {
	Action       ActionClass `json:"action"`
	RootPassword *string     `json:"root_password"` // Root password when no SSH keys have been specified
	Server       ServerClass `json:"server"`
}
