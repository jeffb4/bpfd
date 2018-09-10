package daemon

import (
	"context"
)

type daemon struct{}

// New returns grpc server instance for the daemon.
func New() (grpc.APIServer, error) {
	return &daemon{}, nil
}

func (s *apiServer) CreateRule(ctx context.Context, c *grpc.CreateRuleRequest) (*grpc.CreateRuleResponse, error) {
	return &grpc.CreateRuleResponse{}, nil
}

func (s *apiServer) DeleteRule(ctx context.Context, r *grpc.DeleteRuleRequest) (*grpc.DeleteRuleResponse, error) {
	return &grpc.DeleteRuleResponse{}, nil
}

func (s *apiServer) ListRules(ctx context.Context, r *grpc.ListRulesRequest) (*grpc.ListRulesResponse, error) {
	var rules []*grpc.Rule
	return &grpc.ListRulesResponse{Rules: rules}, nil
}
