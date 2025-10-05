package entity

import "time"

// TODO: 使用genproto生成的文件全部替代
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
