package entity

import "time"

type DeliveryHistory struct {
	ID           string
	Type         string
	CargoID      string
	ScheduleID   string
	CompleteTime time.Time
}
