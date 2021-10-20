package scylla

import (
	"context"
	"time"

	"github.com/gocql/gocql"
)

// Service embed a scylla client.
type Service struct {
	*gocql.Session
}

// Dial connects scylla client.
func (s *Service) Dial(ctx context.Context, cfg Config) error {
	var err error

	// Create gocql cluster.
	cluster := gocql.NewCluster(cfg.Hosts...)
	cluster.Keyspace = cfg.Keyspace
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: cfg.Username,
		Password: cfg.Password,
	}

	if cfg.Timeout != 0 {
		cluster.Timeout = time.Duration(cfg.Timeout) * time.Millisecond
	}

	if cfg.TimeoutConnect != 0 {
		cluster.ConnectTimeout = time.Duration(cfg.TimeoutConnect) * time.Millisecond
	}

	var consistency gocql.Consistency
	if err := consistency.UnmarshalText([]byte(cfg.Consistency)); err != nil {
		return err
	}

	cluster.Consistency = consistency

	s.Session, err = cluster.CreateSession()

	return err
}

// Close closes scylla client.
func (s *Service) Close(ctx context.Context) error {
	s.Session.Close()

	return nil
}
