package storage

import (
	"contest/internal/domain/repository"
	"contest/internal/storage/implementations"
	"contest/lib/adapter/db"
	"contest/lib/adapter/db/postgres"
	"context"
	"time"
)

var (
	MaxOpenConnections = 10
)

type Storage struct {
	db               db.DBAdapter
	TestRepository   repository.TestRepository
	TaskRepository   repository.TaskRepository
	LaunchRepository repository.LaunchRepository
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Init(connStr string, timeout time.Duration) error {
	adapter := postgres.NewPostgresAdapter(timeout)
	_, err := adapter.Connect(context.Background(), connStr)
	if err != nil {
		return err
	}

	s.db = adapter
	s.db.GetConnection().SetMaxOpenConns(MaxOpenConnections)

	s.LaunchRepository = implementations.NewLaunchRepository(s.db)
	s.TestRepository = implementations.NewTestRepository(s.db)
	s.TaskRepository = implementations.NewTaskRepository(s.db)

	return nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}
