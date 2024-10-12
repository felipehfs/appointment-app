package dto

import "time"

type SelectAppointment struct {
	Id         int               `json:"id"`
	Customer   SelectCustomerDTO `json:"customer"`
	ScheduleOn *time.Time        `json:"schedule_on"`
}
