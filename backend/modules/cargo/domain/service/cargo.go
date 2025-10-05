package service

import (
	"context"

	"github.com/chaos-io/cargo/backend/modules/cargo/domain/entity"
)

type CargoService interface {
	CreateCargo(ctx context.Context, cargo *entity.Cargo) (int64, error)
	GetCargo(ctx context.Context, id int64) (*entity.Cargo, error)
	HandlingEvent(ctx context.Context, event *entity.HandlingEventRequest) error
}
