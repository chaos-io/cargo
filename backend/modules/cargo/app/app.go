// TODO: handlers/handlers.go
package app

import (
	"context"

	"github.com/chaos-io/chaos/logs"

	// this service api
	"github.com/chaos-io/cargo/genproto/cargo"
	pb "github.com/chaos-io/cargo/genproto/cargo/v1"
)

type cargoServer struct {
	pb.UnimplementedCargoServer
}

// NewService returns a naive, stateless implementation of Interface.
func NewService() pb.CargoServer {
	return cargoServer{}
}

// CreateResource implements Interface.
func (s cargoServer) CreateResource(ctx context.Context, in *pb.CreateResourceRequest) (*cargo.Resource, error) {
	logs.Infow("CreateResource", "request", in)

	resp := &cargo.Resource{
		// Id:
		// Name:
	}
	return resp, nil
}

// GetResource implements Interface.
func (s cargoServer) GetResource(ctx context.Context, in *pb.GetResourceRequest) (*cargo.Resource, error) {
	logs.Infow("GetResource", "request", in)

	resp := &cargo.Resource{
		// Id:
		// Name:
	}
	return resp, nil
}

// ListResources implements Interface.
func (s cargoServer) ListResources(ctx context.Context, in *pb.ListResourcesRequest) (*pb.ListResourcesResponse, error) {
	logs.Infow("ListResources", "request", in)

	resp := &pb.ListResourcesResponse{
		// Resources:
	}
	return resp, nil
}

// UpdateResource implements Interface.
func (s cargoServer) UpdateResource(ctx context.Context, in *pb.UpdateResourceRequest) (*cargo.Resource, error) {
	logs.Infow("UpdateResource", "request", in)

	resp := &cargo.Resource{
		// Id:
		// Name:
	}
	return resp, nil
}

// DeleteResource implements Interface.
func (s cargoServer) DeleteResource(ctx context.Context, in *pb.DeleteResourceRequest) (*pb.DeleteResourceResponse, error) {
	logs.Infow("DeleteResource", "request", in)

	resp := &pb.DeleteResourceResponse{
		// Id:
	}
	return resp, nil
}
