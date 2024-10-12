package dto

import "time"

type SelectAddressDTO struct {
	Street   string `json:"street" validate:"required"`
	Number   int    `json:"number" validate:"required"`
	Neighbor string `json:"neighbor" validate:"required"`
	State    string `json:"state" validate:"required"`
	City     string `json:"city" validate:"required"`
}

type SelectCustomerDTO struct {
	Id        int              `json:"id"`
	Name      string           `json:"name"`
	CreatedAt *time.Time       `json:"created_at"`
	Sex       string           `json:"sex"`
	Address   SelectAddressDTO `json:"address"`
	BirthDay  *time.Time       `json:"birthday"`
}
