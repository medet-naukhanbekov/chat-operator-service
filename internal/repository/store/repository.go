package store

import "chat_operator_service/internal/domain/store"

type Repository interface {
	GetByAddress(address string) (*store.Store, error)
	Add(store store.Store) error
}
