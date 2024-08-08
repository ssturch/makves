package storage

import (
	"microservice/internal/storage/fakestore"
)

type Storage struct {
	store Store
}

func New() (*Storage, error) {
	st := fakestore.New()
	s := &Storage{
		store: st,
	}
	return s, nil
}

func (s *Storage) Close() {
	s.Close()
}
