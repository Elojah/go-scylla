package scylla

import (
	"context"
)

// Service embed a scylla client.
type Service struct {
}

// Dial connects scylla client.
func (s *Service) Dial(ctx context.Context, cfg Config) error {
	return nil
}

func (s *Service) Close(ctx context.Context) error {
	return nil
}
