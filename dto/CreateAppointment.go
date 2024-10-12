package dto

import "time"

type CreateAppointmentDto struct {
	CustomerId int        `json:"customer_id" validate:"required"`
	ScheduleOn *time.Time `json:"schedule_on" validate:"required"`
}
