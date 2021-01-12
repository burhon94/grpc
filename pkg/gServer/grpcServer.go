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

func (s *GRPCServer) Multiple(ctx context.Context, req *api.Req) (*api.Resp, error) {
	return &api.Resp{Result: req.GetX() * req.GetY()}, nil
}

func (s *GRPCServer) Dived(ctx context.Context, req *api.Req) (*api.Resp, error) {
	return &api.Resp{Result: req.GetX() / req.GetY()}, nil
}

func (s *GRPCServer) PercentAdd(ctx context.Context, req *api.Req) (*api.Resp, error) {
	return &api.Resp{Result: ((req.GetX() * req.GetY()) / 100) + req.GetX()}, nil
}

func (s *GRPCServer) PercentMinus(ctx context.Context, req *api.Req) (*api.Resp, error) {
	return &api.Resp{Result: (req.GetX() * (100 - req.GetY())) / 100}, nil
}
