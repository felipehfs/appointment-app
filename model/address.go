package model

type Address struct {
	Id       int    `json:"id,omitempty"`
	Street   string `json:"street"`
	Number   int    `json:"number"`
	Neighbor string `json:"neighbor"`
	State    string `json:"state"`
	City     string `json:"city"`
}
