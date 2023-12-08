package client

type Task struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Weight      int32  `json:"weight,omitempty"`
}
