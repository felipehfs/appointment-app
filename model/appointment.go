package model

import "time"

type Appointment struct {
	ID         int        `json:"id,omitempty"`
	Customer   Customer   `json:"client"`
	CreatedAt  *time.Time `json:"created_at"`
	ScheduleOn *time.Time `json:"schedule_on"`
}
