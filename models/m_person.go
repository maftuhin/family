package models

// Member represent an object
type Person struct {
	ID     int    `json:"id"`
	UID    string `json:"uid"`
	Name   string `json:"name"`
	Image  string `json:"image"`
	Gender string `json:"gender,omitempty"`
	Father string `json:"father,omitempty"`
	Mother string `json:"mother,omitempty"`
	Spouse string `json:"sid,omitempty"`

	Birth   string `json:"birth,omitempty"`
	Address string `json:"address,omitempty"`
}
