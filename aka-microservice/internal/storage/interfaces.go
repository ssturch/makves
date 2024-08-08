package storage

import "microservice/internal/storage/fakestore"

type Store interface {
	GetInfoByIds(ids []string) ([]fakestore.Item, error)

	Close()
}
