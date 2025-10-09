package repo

import (
	"context"

	"github.com/chaos-io/cargo/backend/modules/cargo-server/domain/entity"
)

type ILocationRepo interface {
	GetLocationByPort(ctx context.Context, port string) (*entity.Location, error)
	GetLocationByCity(ctx context.Context, city string) (*entity.Location, error)
}
