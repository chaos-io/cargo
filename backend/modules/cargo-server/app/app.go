package app

import (
	"context"

	"github.com/chaos-io/chaos/pkg/logs"

	// this service api
	pb "github.com/chaos-io/cargo/genproto/cargo/v1"
)

type cargoServer struct {
	pb.UnimplementedCargoServiceServer
}

// NewService returns a naive, stateless implementation of Interface.
func NewService() pb.CargoServiceServer {
	return cargoServer{}
}

// GetCargo implements Interface.
func (s cargoServer) GetCargo(ctx context.Context, in *pb.GetCargoReq) (*pb.GetCargoResp, error) {
	logs.Infow("GetCargo", "request", in)

	resp := &pb.GetCargoResp{
		// Cargo:
	}
	return resp, nil
}

// CreateCargo implements Interface.
func (s cargoServer) CreateCargo(ctx context.Context, in *pb.CreateCargoReq) (*pb.CreateCargoResp, error) {
	logs.Infow("CreateCargo", "request", in)

	resp := &pb.CreateCargoResp{
		// TrackingId:
	}
	return resp, nil
}

// HandingEvent implements Interface.
func (s cargoServer) HandingEvent(ctx context.Context, in *pb.HandingEventReq) (*pb.HandingEventResp, error) {
	logs.Infow("HandingEvent", "request", in)

	resp := &pb.HandingEventResp{
		// Id:
	}
	return resp, nil
}
