// To parse and unparse this JSON data, add this code to your project and do:
//
//    actions, err := UnmarshalActions(bytes)
//    bytes, err = actions.Marshal()

package models

import "encoding/json"

func UnmarshalActions(data []byte) (Actions, error) {
	var r Actions
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Actions) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Actions struct {
	Actions []Action `json:"actions"`
}

type Action struct {
	Command   string      `json:"command"`   // Command executed in the action
	Error     *ErrorClass `json:"error"`     // Error message for the action if error occured, otherwise null.
	Finished  *string     `json:"finished"`  // Point in time when the action was finished (in ISO-8601 format). Only set if the action; is finished otherwise null.
	ID        float64     `json:"id"`        // ID of the action
	Progress  float64     `json:"progress"`  // Progress of action in percent
	Resources []Resource  `json:"resources"` // Resources the action relates to
	Started   string      `json:"started"`   // Point in time when the action was started (in ISO-8601 format)
	Status    Status      `json:"status"`    // Status of the action
}

type ActionClass struct {
	Command   string      `json:"command"`   // Command executed in the action
	Error     *ErrorClass `json:"error"`     // Error message for the action if error occured, otherwise null.
	Finished  *string     `json:"finished"`  // Point in time when the action was finished (in ISO-8601 format). Only set if the action; is finished otherwise null.
	ID        float64     `json:"id"`        // ID of the action
	Progress  float64     `json:"progress"`  // Progress of action in percent
	Resources []Resource  `json:"resources"` // Resources the action relates to
	Started   string      `json:"started"`   // Point in time when the action was started (in ISO-8601 format)
	Status    Status      `json:"status"`    // Status of the action
}

type ErrorClass struct {
	Code    string `json:"code"`    // Fixed machine readable code
	Message string `json:"message"` // Humanized error message.
}

type Resource struct {
	ID   float64 `json:"id"`   // ID of resource referenced
	Type string  `json:"type"` // Type of resource referenced
}

// Status of the action
type Status string

const (
	StatusError   Status = "error"
	StatusRunning Status = "running"
	StatusSuccess Status = "success"
)
