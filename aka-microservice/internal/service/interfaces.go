package service

import "microservice/internal/storage"

type IStorage interface {
	GetInfoByIds(ids []string) ([]storage.Item, error)
	Close()
}
