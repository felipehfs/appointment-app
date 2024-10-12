package dto

import "time"

type UpdateAddress struct {
	Street   string `json:"street" validate:"required"`
	Number   int    `json:"number" validate:"required"`
	Neighbor string `json:"neighbor" validate:"required"`
	State    string `json:"state" validate:"required"`
	City     string `json:"city" validate:"required"`
}

type UpdateCustomer struct {
	Name     string        `json:"name" validate:"required"`
	SexId    int           `json:"sex_id" validate:"required"`
	Address  CreateAddress `json:"address" validate:"required"`
	Birthday *time.Time    `json:"birthday"`
}
