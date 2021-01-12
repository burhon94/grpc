package gServer

import (
	"context"
	"grpc/pkg/api"
)

type GRPCServer struct {}

func (s *GRPCServer) Add(ctx context.Context, req *api.Req) (*api.Resp, error) {
	return &api.Resp{Result: req.GetX() + req.GetY()}, nil
}

func (s *GRPCServer) Minus(ctx context.Context, req *api.Req) (*api.Resp, error) {
	return &api.Resp{Result: req.GetX() - req.GetY()}, nil
}