package repo

import (
	"context"

	"github.com/chaos-io/cargo/backend/modules/cargo-server/domain/entity"
)

type ICustomerRepo interface {
	GetCustomerByID(ctx context.Context, id string) (*entity.Customer, error)
	GetCustomerByName(ctx context.Context, name string) (*entity.Customer, error)
	GetCustomerByTrackingID(ctx context.Context, id string) (*entity.Customer, error)
}
