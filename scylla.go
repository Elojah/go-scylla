package scylla

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

// Service embed a scylla client.
type Service struct {
	gocqlx.Session
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

	var consistency gocql.Consistency
	if err := consistency.UnmarshalText([]byte(cfg.Consistency)); err != nil {
		return err
	}

	cluster.Consistency = consistency

	// Wrap session on creation, gocqlx session embeds gocql.Session pointer.
	s.Session, err = gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		return err
	}

	return nil
}

// Close closes scylla client.
func (s *Service) Close(ctx context.Context) error {
	s.Session.Close()

	return nil
}
