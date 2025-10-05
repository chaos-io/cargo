package repo

import (
	"context"

	"github.com/chaos-io/cargo/backend/modules/cargo/domain/entity"
)

type ICargoRepo interface {
	GetCargoByTrackingID(ctx context.Context, tid string) (*entity.Cargo, error)
	GetCargoByCustomerID(ctx context.Context, cid string) (*entity.Cargo, error)
}
