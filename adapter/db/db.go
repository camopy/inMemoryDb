package db

import inmemory "github.com/camopy/mosaico/adapter/db/inMemory"

type DB interface {
	Get(id uint64) (string, error)
	Set(id uint64, message string)
}

func New() DB {
	return inmemory.New()
}
