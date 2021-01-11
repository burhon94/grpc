package adder

import (
	"context"
	"grpc/pkg/api"
)

type GRPCServer struct {}

func (s *GRPCServer) Add(ctx context.Context, req *api.AddReq) (*api.AddResp, error) {
	return &api.AddResp{Result: req.GetX() + req.GetY()}, nil
}