package db

import inmemory "github.com/camopy/in-memory-db/adapter/db/inMemory"

type DB interface {
	Get(id uint64) (string, error)
	Set(id uint64, message string)
}

func New() DB {
	return inmemory.New()
}
