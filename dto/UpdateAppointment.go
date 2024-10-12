package dto

import "time"

type UpdateAppointmentDto struct {
	CustomerId int        `json:"customer_id" validate:"required"`
	ScheduleOn *time.Time `json:"schedule_on" validate:"required"`
}
