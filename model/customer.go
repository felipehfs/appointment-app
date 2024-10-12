package model

import "time"

type Customer struct {
	Id        int        `json:"id,omitempty"`
	Name      string     `json:"name"`
	Sex       *Sex       `json:"sex"`
	CreatedAt *time.Time `json:"created_at"`
	Address   *Address   `json:"address"`
	BirthDay  *time.Time `json:"birthday,omitempty"`
}
