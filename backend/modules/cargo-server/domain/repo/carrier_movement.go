package repo

import (
	"context"

	"github.com/chaos-io/cargo/backend/modules/cargo/domain/entity"
)

type ICarrierMovementRepo interface {
	GetCarrierMovementByScheduleID(ctx context.Context, id string) (*entity.CarrierMovement, error)
	GetCarrierMovementFromTo(ctx context.Context, from, to *entity.Location) (*entity.CarrierMovement, error)
}
