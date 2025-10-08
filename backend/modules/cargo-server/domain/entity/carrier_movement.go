package entity

type CarrierMovement struct {
	ScheduleID string
	From       *Location
	To         *Location
}
