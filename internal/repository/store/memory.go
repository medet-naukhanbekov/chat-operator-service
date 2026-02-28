package store

import (
	"chat_operator_service/internal/domain/store"
	"errors"
)

type MemoryRepository struct {
	stores map[string]store.Store
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		stores: map[string]store.Store{
			"Алматы Абая 10": {
				Address: "Алматы Абая 10",
				OpenAt:  "10:00",
				CloseAt: "22:00",
			},
			"Астана Назарбаева 5": {
				Address: "Астана Назарбаева 5",
				OpenAt:  "09:00",
				CloseAt: "21: 00",
			},
		},
	}
}

func (r *MemoryRepository) GetByAddress(address string) (*store.Store, error) {
	s, ok := r.stores[address]
	if !ok {
		return nil, errors.New("store not found")
	}
	return &s, nil
}

func (r *MemoryRepository) Add(s store.Store) error {
	r.stores[s.Address] = s
	return nil
}
