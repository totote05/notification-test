package mock

import (
	"errors"
	"notificator/registry"
	"sync"
)

type RegistryMockRepository struct {
	failureMode bool
	registry    []registry.Record
	mu          sync.Mutex
}

func (r *RegistryMockRepository) SetFailureMode(active bool) {
	r.failureMode = active
}

func (r *RegistryMockRepository) Add(record registry.Record) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.registry = append(r.registry, record)

	if r.failureMode {
		return errors.New("fail to add record, failure mode is activated")
	} else {
		return nil
	}
}

func (r *RegistryMockRepository) GetAll() []registry.Record {
	r.mu.Lock()
	defer r.mu.Unlock()

	result := make([]registry.Record, len(r.registry))
	copy(result, r.registry)

	return result
}

func NewRegistryRepository() registry.Repository {
	return &RegistryMockRepository{
		failureMode: false,
		registry:    []registry.Record{},
	}
}

var _ registry.Repository = (*RegistryMockRepository)(nil)
