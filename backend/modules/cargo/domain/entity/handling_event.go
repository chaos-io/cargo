package entity

import "time"

type HandlingEvent struct {
	cargo *Cargo

	Type           EventType
	CompletingTime time.Time

	movement *CarrierMovement
}

type EventType string

const (
	SentEventType               EventType = "Sent"
	AOGEventType                EventType = "AOG" // Arrival of Goods(AOG)
	DeliveryEventType           EventType = "Delivery"
	CustomsDeclarationEventType EventType = "CustomsDeclaration"
	CustomsClearanceEventType   EventType = "CustomsClearance"
	CustomsInspectionEventType  EventType = "CustomsInspection"
	CustomsReleaseEventType     EventType = "CustomsRelease"
)

type HandlingEventRequest struct {
	CargoID        string
	ScheduleID     string
	EventType      EventType
	CompletingTime time.Time
}
