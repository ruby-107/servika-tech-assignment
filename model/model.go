package models

// Pet model
type Pet struct {
	ID      int     `json:"id,omitempty"`
	Name    string  `json:"name,omitempty"`
	Owner   string  `json:"owner,omitempty"`
	Species string  `json:"species,omitempty"`
	Birth   string  `json:"birth,omitempty"`
	Death   string  `json:"death,omitempty"`
	Events  []Event `json:"events,omitempty"`
}

type Event struct {
	PetID  string `json:"id,omitempty"`
	Date   string `json:"date,omitempty"`
	Type   string `json:"type,omitempty"`
	Remark string `json:"remark,omitempty"`
}
