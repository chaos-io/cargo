package entity

import "time"

type Cargo struct {
	TrackingID string
	CustomerID string

	// role
	Shipper   *Customer
	Consignee *Customer

	Spec DeliverySpecification
}

// DeliverySpecification value object
type DeliverySpecification struct {
	ID string

	From *Location
	To   *Location

	// Estimated Time of Arrival(ETA)
	ETA time.Time `json:"eta"`
}
