package store

import (
	"github.com/hashicorp/raft"
)

type RavelLogStore struct {
}

func NewRavelLogStore() *RavelLogStore {
	return &RavelLogStore{}
}

func (r *RavelLogStore) FirstIndex() (uint64, error) {
	return 0, nil
}

func (r *RavelLogStore) LastIndex() (uint64, error) {
	return 0, nil
}

func (r *RavelLogStore) GetLog(index uint64, log *raft.Log) {}

func (r *RavelLogStore) StoreLog(log *raft.Log) error {
	return r.StoreLogs([]*raft.Log{log})
}

func (r *RavelLogStore) StoreLogs(logs []*raft.Log) error {
	return nil
}

func (r *RavelLogStore) DeleteRange(min, max uint64) error {
	return nil
}
